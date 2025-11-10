package rvcore

func (line DecodedTokenLine) EmitAsmLine() uint32 {

	switch line.Type {

	case RTYPE:
		instr := instrMap[line.Instr]
		return EncodeRType(instr.Opcode, line.Rd, instr.Funct3, line.Rs1, line.Rs2, instr.Funct7)
	}

	return 0
}
