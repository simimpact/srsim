package ast

import (
	"strings"
)

type Node interface {
	String() string
	// Copy does a deep copy of the Node and all its components.
	// To avoid type assertions, some XxxNodes also have specialized
	// CopyXxx methods that return *XxxNode.
	Copy() Node
	Position() Pos // byte position of start of node in full original input string
	// writeTo writes the String output to the builder.
	writeTo(*strings.Builder)
}

type Pos int

func (p Pos) Position() Pos {
	return p
}
