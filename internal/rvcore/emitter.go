package rvcore

func (tokens TokenLine) EmitAsmLine() uint32 {
	if len(tokens.Tokens) == 0 {
		return 0
	}

	switch tokens.Tokens[0].TokenType {
	case INSTRUCTION:
		switch instrMap[string(tokens.Tokens[0].Value)].Format {
		case RTYPE:
			instr := instrMap[string(tokens.Tokens[0].Value)]
			rd := regMap[string(tokens.Tokens[1].Value)]
			rs1 := regMap[string(tokens.Tokens[3].Value)]
			rs2 := regMap[string(tokens.Tokens[5].Value)]
			return EncodeRType(instr.Opcode, rd, instr.Funct3, rs1, rs2, instr.Funct7)
		}

	}

	return 0
}
