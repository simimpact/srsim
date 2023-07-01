package parse

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/simimpact/srsim/pkg/gcs/ast"
)

const eof = -1

type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	input        string         // the string being scanned
	pos          ast.Pos        // current position in the input
	start        ast.Pos        // start position of this item
	width        ast.Pos        // width of last rune read from input
	items        chan ast.Token // channel of scanned items
	line         int            // 1+number of newlines seen
	startLine    int            // start line of this item
	parenDepth   int
	sqParenDepth int
	braceDepth   int
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = ast.Pos(w)
	l.pos += l.width
	if r == '\n' {
		l.line++
	}
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
	// Correct newline count.
	if l.width == 1 && l.input[l.pos] == '\n' {
		l.line--
	}
}

// emit passes an item back to the client.
func (l *lexer) emit(t ast.TokenType) {
	l.items <- ast.Token{
		Typ:  t,
		Pos:  l.start,
		Val:  l.input[l.start:l.pos],
		Line: l.startLine,
	}
	l.start = l.pos
	l.startLine = l.line
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	// l.line += strings.Count(l.input[l.start:l.pos], "\n")
	l.start = l.pos
	l.startLine = l.line
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- ast.Token{
		Typ:  ast.ItemError,
		Pos:  l.start,
		Val:  fmt.Sprintf(format, args...),
		Line: l.startLine,
	}
	return nil
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextItem() ast.Token {
	return <-l.items
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
// func (l *lexer) drain() {
// 	for range l.items {
// 	}
// }

// lex creates a new scanner for the input string.
func lex(input string) *lexer {
	l := &lexer{
		input:     input,
		items:     make(chan ast.Token),
		line:      1,
		startLine: 1,
	}
	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

// lexText scans until an opening action delimiter, "{{".
func lexText(l *lexer) stateFn {
	// Either number, quoted string, or identifier.
	// Spaces separate arguments; runs of spaces turn into itemSpace.
	// Pipe symbols separate and are emitted.
	// n := l.peek()
	// log.Printf("lexText next is %c\n", n)
	switch r := l.next(); {
	case r == eof:
		l.emit(ast.ItemEOF)
		return nil
	case r == ';':
		l.emit(ast.ItemTerminateLine)
	case r == ':':
		l.emit(ast.ItemColon)
	case isSpace(r):
		l.ignore()
	case r == '#':
		l.ignore()
		return lexComment
	case r == '=':
		n := l.next()
		if n == '=' {
			l.emit(ast.OpEqual)
		} else {
			l.backup()
			l.emit(ast.ItemAssign)
		}
	case r == ',':
		l.emit(ast.ItemComma)
	case r == '*':
		l.emit(ast.ItemAsterisk)
	case r == '+':
		// //check if next item is a number or not; if number lexNumber
		// //otherwise it's a + sign
		// n := l.next()
		// if isNumeric(n) {
		// 	//back up twice
		// 	l.backup()
		// 	l.backup()
		// 	return lexNumber
		// }
		// //otherwise it's a plus sign
		// l.backup()
		l.emit(ast.ItemPlus)
	case r == '/':
		// check if next is another / or not; if / then lexComment
		n := l.next()
		if n == '/' {
			l.ignore()
			return lexComment
		}
		l.backup()
		l.emit(ast.ItemForwardSlash)
	case r == '.':
		n := l.next()
		if isNumeric(n) {
			// backup twice
			l.backup()
			l.backup()
			return lexNumber
		}
		l.backup()
		return l.errorf("unrecognized character in action: %#U", r)
	case ('0' <= r && r <= '9'):
		l.backup()
		return lexNumber
	case r == '-':
		// if next item is a number then lex number
		n := l.next()
		if isNumeric(n) {
			// backup twice
			l.backup()
			l.backup()
			return lexNumber
		}
		// other wise it's a - sign
		l.backup()
		l.emit(ast.ItemMinus)
	case r == '>':
		if n := l.next(); n == '=' {
			l.emit(ast.OpGreaterThanOrEqual)
		} else {
			l.backup()
			l.emit(ast.OpGreaterThan)
		}
	case r == '<':
		switch n := l.next(); n {
		case '=':
			l.emit(ast.OpLessThanOrEqual)
		case '>':
			l.emit(ast.OpNotEqual)
		default:
			l.backup()
			l.emit(ast.OpLessThan)
		}
	case r == '|':
		if n := l.next(); n == '|' {
			l.emit(ast.LogicOr)
		} else {
			return l.errorf("unrecognized character in action: %#U", r)
		}
	case r == '!':
		if n := l.next(); n == '=' {
			l.emit(ast.OpNotEqual)
		} else {
			l.backup()
			l.emit(ast.LogicNot)
		}
	case r == '"':
		return lexQuote
	case r == '&':
		if n := l.next(); n == '&' {
			l.emit(ast.LogicAnd)
		} else {
			return l.errorf("unrecognized character in action: %#U", r)
		}
	case r == '(':
		l.emit(ast.ItemLeftParen)
		l.parenDepth++
	case r == ')':
		l.emit(ast.ItemRightParen)
		l.parenDepth--
		if l.parenDepth < 0 {
			return l.errorf("unexpected right paren %#U", r)
		}
	case r == '[':
		l.emit(ast.ItemLeftSquareParen)
		l.sqParenDepth++
	case r == ']':
		l.emit(ast.ItemRightSquareParen)
		l.sqParenDepth--
		if l.sqParenDepth < 0 {
			return l.errorf("unexpected right sq paren %#U", r)
		}
	case r == '{':
		l.emit(ast.ItemLeftBrace)
		l.braceDepth++
	case r == '}':
		l.emit(ast.ItemRightBrace)
		l.braceDepth--
		if l.braceDepth < 0 {
			return l.errorf("unexpected right brace %#U", r)
		}
	case isAlphaNumeric(r):
		l.backup()
		return lexIdentifier
	default:
		return l.errorf("unrecognized character in action: %#U", r)
	}
	return lexText
}

func lexComment(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case eof, '\n':
			l.backup()
			break Loop
		default:
			// absorb
		}
	}
	// l.emit(itemComment)
	return lexText
}

func lexQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
	l.emit(ast.ItemString)
	return lexText
}

// lexIdentifier scans an alphanumeric.
func lexIdentifier(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// absorb.
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			if !l.atTerminator() {
				return l.errorf("bad character %#U", r)
			}
			switch {
			case ast.Keys[word] > ast.ItemKeyword:
				l.emit(ast.Keys[word])
			case word == "true", word == "false":
				l.emit(ast.ItemBool)
			case word == "null":
				l.emit(ast.ItemNull)
			default:
				l.emit(checkIdentifier(word))
			}
			break Loop
		}
	}
	return lexText
}

func checkIdentifier(word string) ast.TokenType {
	// TODO: need items?
	return ast.ItemIdentifier
}

func lexNumber(l *lexer) stateFn {
	// Optional leading sign.
	l.accept("+-")

	digits := "0123456789"
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}

	l.emit(ast.ItemNumber)

	return lexText
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || r == '-' || unicode.IsLetter(r) || unicode.IsDigit(r) || r == '%'
}

// is Numeric reports whether r is a digit
func isNumeric(r rune) bool {
	return unicode.IsDigit(r)
}

// atTerminator reports whether the input is at valid termination character to
// appear after an identifier. Breaks .X.Y into two pieces. Also catches cases
// like "$x+2" not being acceptable without a space, in case we decide one
// day to implement arithmetic.
func (l *lexer) atTerminator() bool {
	r := l.peek()
	if isSpace(r) {
		return true
	}
	switch r {
	case eof, '.', ',', '|', ':', ')', '(', '+', '=', '>', '<', '&', '!', ';', '[', ']':
		return true
	}
	return false
}
