package ast

var Keys = map[string]TokenType{
	".":           ItemDot,
	"let":         KeywordLet,
	"while":       KeywordWhile,
	"if":          KeywordIf,
	"else":        KeywordElse,
	"fn":          KeywordFn,
	"switch":      KeywordSwitch,
	"case":        KeywordCase,
	"default":     KeywordDefault,
	"break":       KeywordBreak,
	"continue":    KeywordContinue,
	"fallthrough": KeywordFallthrough,
	"return":      KeywordReturn,
	"for":         KeywordFor,
}
