package rvcore

func (line DecodedTokenLine) EmitAsmLine() uint32 {

	switch line.Type {

	case RTYPE:
		instr := instrMap[line.Instr]
		return EncodeRType(instr.Opcode, line.Rd, instr.Funct3, line.Rs1, line.Rs2, instr.Funct7)
	case ITYPE:
		instr := instrMap[line.Instr]
		return EncodeIType(instr.Opcode, line.Rd, instr.Funct3, line.Rs1, line.Imm)
	case STYPE:
		instr := instrMap[line.Instr]
		return EncodeSType(instr.Opcode, line.Imm, instr.Funct3, line.Rs1, line.Rs2)
	}

	return 0
}
