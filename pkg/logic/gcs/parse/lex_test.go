package parse

import (
	"testing"

	"github.com/simimpact/srsim/pkg/logic/gcs/ast"
)

func TestBasicToken(t *testing.T) {
	input := `
	let y = fn(x) {
		return x + 1;
	}
	let x = 5;
	while {
		#comment
		x = y(x);
		if x > 10 {
			break A;
		}
		//comment
		switch x {
		case 1:
			fallthrough;
		case 2:
			fallthrough;
		case 3:
			break A;
		}
	}
	
	for x = 0; x < 5; x = x + 1 {
		let i = y(x);
	}

	-1
	1
	-
	-a
	.123
	`

	expected := []ast.Token{
		// function
		{Typ: ast.KeywordLet, Val: "let"},
		{Typ: ast.ItemIdentifier, Val: "y"},
		{Typ: ast.ItemAssign, Val: "="},
		{Typ: ast.KeywordFn, Val: "fn"},
		{Typ: ast.ItemLeftParen, Val: "("},
		{Typ: ast.ItemIdentifier, Val: "x"},
		// {typ: typeNum, Val: "num"},
		{Typ: ast.ItemRightParen, Val: ")"},
		// {typ: typeNum, Val: "num"}
		{Typ: ast.ItemLeftBrace, Val: "{"},
		{Typ: ast.KeywordReturn, Val: "return"},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.ItemPlus, Val: "+"},
		{Typ: ast.ItemNumber, Val: "1"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		{Typ: ast.ItemRightBrace, Val: "}"},
		// variable
		{Typ: ast.KeywordLet, Val: "let"},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.ItemAssign, Val: "="},
		{Typ: ast.ItemNumber, Val: "5"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		// while loop
		{Typ: ast.KeywordWhile, Val: "while"},
		{Typ: ast.ItemLeftBrace, Val: "{"},
		// comment
		// {typ: itemComment, Val: "comment"},
		// function call
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.ItemAssign, Val: "="},
		{Typ: ast.ItemIdentifier, Val: "y"},
		{Typ: ast.ItemLeftParen, Val: "("},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.ItemRightParen, Val: ")"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		// if statement
		{Typ: ast.KeywordIf, Val: "if"},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.OpGreaterThan, Val: ">"},
		{Typ: ast.ItemNumber, Val: "10"},
		{Typ: ast.ItemLeftBrace, Val: "{"},
		// break
		{Typ: ast.KeywordBreak, Val: "break"},
		{Typ: ast.ItemIdentifier, Val: "A"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		// end if
		{Typ: ast.ItemRightBrace, Val: "}"},
		// comment
		// {typ: itemComment, Val: "comment"},
		// switch
		{Typ: ast.KeywordSwitch, Val: "switch"},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.ItemLeftBrace, Val: "{"},
		// case
		{Typ: ast.KeywordCase, Val: "case"},
		{Typ: ast.ItemNumber, Val: "1"},
		{Typ: ast.ItemColon, Val: ":"},
		{Typ: ast.KeywordFallthrough, Val: "fallthrough"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		// case
		{Typ: ast.KeywordCase, Val: "case"},
		{Typ: ast.ItemNumber, Val: "2"},
		{Typ: ast.ItemColon, Val: ":"},
		{Typ: ast.KeywordFallthrough, Val: "fallthrough"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		// case
		{Typ: ast.KeywordCase, Val: "case"},
		{Typ: ast.ItemNumber, Val: "3"},
		{Typ: ast.ItemColon, Val: ":"},
		{Typ: ast.KeywordBreak, Val: "break"},
		{Typ: ast.ItemIdentifier, Val: "A"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		// end switch
		{Typ: ast.ItemRightBrace, Val: "}"},
		// end while
		{Typ: ast.ItemRightBrace, Val: "}"},
		// for loop
		{Typ: ast.KeywordFor, Val: "for"},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.ItemAssign, Val: "="},
		{Typ: ast.ItemNumber, Val: "0"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.OpLessThan, Val: "<"},
		{Typ: ast.ItemNumber, Val: "5"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.ItemAssign, Val: "="},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.ItemPlus, Val: "+"},
		{Typ: ast.ItemNumber, Val: "1"},
		{Typ: ast.ItemLeftBrace, Val: "{"},
		// body
		{Typ: ast.KeywordLet, Val: "let"},
		{Typ: ast.ItemIdentifier, Val: "i"},
		{Typ: ast.ItemAssign, Val: "="},
		{Typ: ast.ItemIdentifier, Val: "y"},
		{Typ: ast.ItemLeftParen, Val: "("},
		{Typ: ast.ItemIdentifier, Val: "x"},
		{Typ: ast.ItemRightParen, Val: ")"},
		{Typ: ast.ItemTerminateLine, Val: ";"},
		// end for
		{Typ: ast.ItemRightBrace, Val: "}"},
		// misc tests
		{Typ: ast.ItemNumber, Val: "-1"},
		{Typ: ast.ItemNumber, Val: "1"},
		{Typ: ast.ItemMinus, Val: "-"},
		{Typ: ast.ItemMinus, Val: "-"},
		{Typ: ast.ItemIdentifier, Val: "a"},
		{Typ: ast.ItemNumber, Val: ".123"},
	}

	l := lex(input)
	i := 0
	for n := l.nextItem(); n.Typ != ast.ItemEOF; n = l.nextItem() {
		if expected[i].Typ != n.Typ && expected[i].Val != n.Val {
			t.Errorf("expected %v got %v", expected[i], n)
		}
		if i < len(expected)-1 {
			i++
		}
	}
}
