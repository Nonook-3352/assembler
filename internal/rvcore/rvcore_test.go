package rvcore

import "testing"

func TestEncodeRType(t *testing.T) {
	if EncodeRType(0b0110011, 0b00111, 0b000, 0b00101, 0b00000, 0b0000000) != 0b0000000_00000_00101_000_00111_0110011 { // ADD x7, x5, x0
		t.Errorf("want %b but got %b", 0b00000000000001110000100000110011, EncodeRType(0b0110011, 0b00111, 0b000, 0b00101, 0b00000, 0b0000000))
	}
}

func TestEncodeIType(t *testing.T) {
	if EncodeIType(0b001_0011, 0b0_0111, 0b000, 0b0_0000, 0b0000_0001_0111) != 0b000000010111_00000_000_00111_0010011 { // ADDI x7, x0, 23
		t.Errorf("want %032b but got %032b", 0b000000010111_00000_000_00111_0010011, EncodeIType(0b001_0011, 0b0_0111, 0b0_0000, 0b000, 0b0000_0001_0111))
	}
}

func TestEncodeSType(t *testing.T) {
	if EncodeSType(0b0100011, 0b000000000001, 0b000, 0b00010, 0b00110) != 0b00000000011000010000000010100011 { //SB x6, 1(x2)
		t.Errorf("want %032b but go %032b", 0b00000000011000010000000010100011, EncodeSType(0b0100011, 0b000000000001, 0b000, 0b00010, 0b00110))
	}
}

func TestEncodeSBType(t *testing.T) {
	result := EncodeSBType(0b1100011, 0b1_1000_0000_0000, 0b000, 0b00010, 0b00110)
	expected := 0b10000000011000010000000011100011
	if result != uint32(expected) {
		t.Errorf("want %032b but got %032b", expected, result)
	}
}

func TestEncodeUType(t *testing.T) {
	result := EncodeUType(0b011_0111, 0b0_0111, 0b1111_1111_1111_1111_1111)
	expected := 0b11111111111111111111001110110111
	if result != uint32(expected) {
		t.Errorf("want %032b but got %032b", expected, result)
	}
}

func TestEncodeUJType(t *testing.T) {
	result := EncodeUJType(0b1101111, 0b00111, 0b1001_0111_1111_1111_0000)
	expected := 0b01111111000110010111_00111_1101111
	if result != uint32(expected) {
		t.Errorf("want %032b but got %032b", expected, result)
	}
}

func TestContains(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	if !Contains(slice, "c") {
		t.Errorf("slice does contain c but got false")
	}

	if Contains(slice, "e") {
		t.Errorf("slice does not contain e but got true")
	}
}

func TestLexeLineAndRefineTokens(t *testing.T) {
	line := Line{Value: "  ADDI a0, x14, 10, 0xF1, 0b11001, 1 ", Len: uint16(len("  ADDI a0, a1, 10, 0xF1, 0b11001, 1 ")), FilePos: 1}
	tokens := line.LexeLine()
	err, tokens := tokens.RefineTokens()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(tokens.Tokens) != 12 {
		t.Errorf("want 12 tokens but got %d", len(tokens.Tokens))
	}
	expectedTokens := []Token{
		{TokenType: INSTRUCTION, Value: "ADDI", OptionalType: UNDEFINED},
		{TokenType: OPERAND, Value: "a0", OptionalType: REGISTER},
		{TokenType: COMMA, Value: ",", OptionalType: UNDEFINED},
		{TokenType: OPERAND, Value: "x14", OptionalType: REGISTER},
		{TokenType: COMMA, Value: ",", OptionalType: UNDEFINED},
		{TokenType: OPERAND, Value: "10", OptionalType: UNDEFINED},
		{TokenType: COMMA, Value: ",", OptionalType: UNDEFINED},
		{TokenType: OPERAND, Value: "0xF1", OptionalType: IMMEDIATEHEX},
		{TokenType: COMMA, Value: ",", OptionalType: UNDEFINED},
		{TokenType: OPERAND, Value: "0b11001", OptionalType: IMMEDIATEBYT},
		{TokenType: COMMA, Value: ",", OptionalType: UNDEFINED},
		{TokenType: OPERAND, Value: "1", OptionalType: UNDEFINED},
	}
	for i, expected := range expectedTokens {
		if tokens.Tokens[i] != expected {
			t.Errorf("at token %d, want %+v but got %+v", i, expected, tokens.Tokens[i])
		}
	}

	line2 := Line{Value: ".arch RISCV32I", Len: uint16(len(".arch RISCV32I")), FilePos: 2}
	tokens2 := line2.LexeLine()
	err, tokens2 = tokens2.RefineTokens()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(tokens2.Tokens) != 2 {
		t.Errorf("want 2 tokens but got %d", len(tokens2.Tokens))
	}
	expectedTokens2 := []Token{
		{TokenType: DIRECTIVE, Value: ".arch", OptionalType: UNDEFINED},
		{TokenType: OPERAND, Value: "RISCV32I", OptionalType: UNDEFINED},
	}
	for i, expected := range expectedTokens2 {
		if tokens2.Tokens[i] != expected {
			t.Errorf("at token %d, want %+v but got %+v", i, expected, tokens2.Tokens[i])
		}
	}

	line3 := Line{Value: "", Len: 0, FilePos: 3}
	tokens3 := line3.LexeLine()
	err, tokens3 = tokens3.RefineTokens()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(tokens3.Tokens) != 0 {
		t.Errorf("want 0 tokens but got %d", len(tokens3.Tokens))
	}

	line4 := Line{Value: "ADDI a0, , x14", Len: uint16(len("ADDI a0, , x14")), FilePos: 4}
	tokens4 := line4.LexeLine()
	err, _ = tokens4.RefineTokens()
	if err == nil {
		t.Errorf("expected error but got nil")
	}
}
