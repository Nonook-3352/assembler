package rvcore

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
	PLACEHOLDER OptionalType = iota
	IMMEDIATEINT
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

type Instruction struct {
	Format InstrFormat
	Opcode uint32 // 7 bit
	Funct3 uint32 // 3 bit
	Funct7 uint32 // 7 bit
}
