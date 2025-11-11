package rvcore

import "fmt"

type TokenType uint8
type TokenValue string

type OptionalType uint8

type InstrFormat uint8

type (
	OpcodeType uint32
	RdType     uint32
	Rs1Type    uint32
	Rs2Type    uint32
	ImmType    uint32
	LabelType  uint32
)

//go:generate stringer -type=TokenType,OptionalType,InstrFormat
const (
	INSTRUCTION TokenType = iota
	OPERAND
	DIRECTIVE
	LABEL
	COMMA
)

const (
	PLACEHOLDER OptionalType = iota
	IMMEDIATEINT
	IMMEDIATEHEX
	IMMEDIATEBYT
	REGISTER
	UNDEFINED
)

const (
	UNDEFINEDINSTR InstrFormat = iota
	RTYPE                      // Maybe draw the bit fields
	ITYPE                      //
	STYPE                      //
	BTYPE                      //
	UTYPE                      //
	JTYPE                      //
)

const (
	X0 uint32 = iota
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	X29
	X30
	X31
)

type Token struct {
	TokenType    TokenType
	OptionalType OptionalType
	Value        TokenValue
}

type Line struct {
	Value   string
	pos     uint16
	Len     uint16
	FilePos uint
}

type TokenLine struct {
	Tokens  []Token
	FilePos uint
}

type DecodedTokenLine struct {
	Type    InstrFormat
	Instr   string
	Rd      uint32
	Rs1     uint32
	Rs2     uint32
	Imm     uint32
	Label   uint32
	FilePos uint
}

type Instruction struct {
	Format InstrFormat
	Opcode uint32 // 7 bit
	Funct3 uint32 // 3 bit
	Funct7 uint32 // 7 bit
}

type ParseError struct {
	Line    uint
	Token   uint
	Message string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("Line %d, Token %d: %s", e.Line, e.Token, e.Message)
}
