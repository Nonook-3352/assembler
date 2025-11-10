package rvcore

func EncodeRType(opcode, rd, funct3, rs1, rs2, funct7 uint32) uint32 {
	var inst uint32
	inst |= opcode
	inst |= rd << 7
	inst |= funct3 << 12
	inst |= rs1 << 15
	inst |= rs2 << 20
	inst |= funct7 << 25
	return inst
}

func EncodeIType(opcode, rd, funct3, rs1, imm uint32) uint32 {
	var inst uint32
	inst |= opcode
	inst |= rd << 7
	inst |= funct3 << 12
	inst |= rs1 << 15
	inst |= imm << 20
	return inst
}

func EncodeSType(opcode, imm, funct3, rs1, rs2 uint32) uint32 {
	var inst uint32
	inst |= opcode
	inst |= (imm & 0b1_1111) << 7 //imm[4:0]
	inst |= funct3 << 12
	inst |= rs1 << 15
	inst |= rs2 << 20
	inst |= (imm >> 5 & 0b111_1111) << 25 //imm[11:5]
	return inst
}

func EncodeBType(opcode, imm, funct3, rs1, rs2 uint32) uint32 {
	var inst uint32
	inst |= opcode
	inst |= ((imm & 0b1000_0000_0000) >> 11) << 7 //imm[11]
	inst |= ((imm & 0b1_1110) >> 1) << 8          //imm[4:1]
	inst |= funct3 << 12
	inst |= rs1 << 15
	inst |= rs2 << 20
	inst |= ((imm & 0b111_1110_0000) >> 5) << 25     //imm[10:5]
	inst |= ((imm & 0b1_0000_0000_0000) >> 12) << 31 //imm[12]
	return inst
}

func EncodeUType(opcode, rd, imm uint32) uint32 {
	var inst uint32
	inst |= opcode
	inst |= rd << 7
	inst |= imm << 12

	return inst
}

func EncodeJType(opcode, rd, imm uint32) uint32 {
	var inst uint32
	inst |= opcode
	inst |= rd << 7
	// Tried another way to take slices of the bits
	inst |= ((imm >> 12) & 0b1111_1111) << 12   //imm[19:12]
	inst |= ((imm >> 11) & 0b1) << 20           //imm[11]
	inst |= ((imm >> 1) & 0b11_1111_1111) << 21 //imm[10:1]
	inst |= ((imm >> 20) & 0b1) << 31           //imm[20]
	return inst
}
