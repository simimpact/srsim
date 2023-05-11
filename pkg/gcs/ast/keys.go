package ast

import "github.com/simimpact/srsim/pkg/key"

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

var ActionKeys = map[string]key.ActionType{
	"attack": key.ActionAttack,
	"skill":  key.ActionSkill,
	"burst":  key.ActionBurst,
}
