package types

type TokenType uint8
type TokenValue string

type OptionalType uint8

type InstrFormat uint8

const (
	INSTRUCTION TokenType = iota
	OPERAND
	DIRECTIVE
	LABEL
	COMMA
)

const (
	IMMEDIATEINT OptionalType = iota
	IMMEDIATEHEX
	IMMEDIATEBYT
	REGISTER
	UNDEFINED
)

const (
	RTYPE  InstrFormat = iota // Maybe draw the bit fields
	ITYPE                     //
	STYPE                     //
	SBTYPE                    //
	UTYPE                     //
	UJTYPE                    //
)
