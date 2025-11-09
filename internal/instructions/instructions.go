package instructions

import . "assembler/internal/types"

type Instruction struct {
	Format InstrFormat
	Opcode uint32 // 7 bit
	Funct3 uint32 // 3 bit
	Funct7 uint32 // 7 bit
}

var instrMap map[string]Instruction = map[string]Instruction{
	"ADD": Instruction{Format: RTYPE, Opcode: 0b011_0011, Funct3: 0b000, Funct7: 0b000_0000},
	"SUB": Instruction{Format: RTYPE, Opcode: 0b011_0011, Funct3: 0b000, Funct7: 0b011_0000},
}
