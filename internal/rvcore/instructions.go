package rvcore

var instrMap map[string]Instruction = map[string]Instruction{
	"ADD": Instruction{Format: RTYPE, Opcode: 0b011_0011, Funct3: 0b000, Funct7: 0b000_0000},
	"SUB": Instruction{Format: RTYPE, Opcode: 0b011_0011, Funct3: 0b000, Funct7: 0b011_0000},
}
