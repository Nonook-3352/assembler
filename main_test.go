package main

import (
	"testing"
)

func TestEncodeRType(t *testing.T) {
	if 0b0000000_00000_00101_000_00111_0110011 != encodeRType(0b0110011, 0b00111, 0b000, 0b00101, 0b00000, 0b0000000) { // ADD x7, x5, x0
		t.Errorf("want %b but got %b", 0b00000000000001110000100000110011, encodeRType(0b0110011, 0b00111, 0b000, 0b00101, 0b00000, 0b0000000))
	}
}
