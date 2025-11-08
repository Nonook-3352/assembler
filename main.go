package main

import (
	"bufio"
	"fmt"
	"os"
)

type TokenType uint8
type TokenValue string

type OptionalType uint8

const (
	INSTRUCTION TokenType = iota
	OPERAND
	DIRECTIVE
	LABEL
	COMMA
)

const (
	IMMEDIATEINT OptionalType = iota
	IMMEDIATEHEX
	IMMEDIATEBYT
	REGISTER
	UNDEFINED
)

var registerABI []string = []string{
	"zero",           //x0
	"ra",             //x1
	"sp",             //x2
	"gp",             //x3
	"tp",             //x4
	"t0", "t1", "t2", //x5-7
	"so", "fp", //x8
	"s1",       //x9
	"a0", "a1", //x10-11
	"a2", "a3", "a4", "a5", "a6", "a7", //x12-x17
	"s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "s10", "s11", //x18-x27
	"t3", "t4", "t5", "t6", //x28-31
}

type Token struct {
	tokenType    TokenType
	optionalType OptionalType
	value        TokenValue
}

type Line struct {
	value   string
	pos     uint16
	len     uint16
	filePos uint
}

type TokenLine struct {
	tokens  []Token
	filePos uint
}

func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (line *Line) skipWhitespace() {
	for line.pos < line.len && (line.value[line.pos] == ' ' || line.value[line.pos] == '\t') {
		line.pos++
	}
}

func (line *Line) readWord() TokenValue {
	wordStart := line.pos
	for line.pos < line.len && (line.value[line.pos] != ' ' && line.value[line.pos] != ',') {
		line.pos++
	}

	return TokenValue(line.value[wordStart:line.pos])
}

func (line Line) lexeLine() TokenLine {
	tokens := TokenLine{
		tokens:  make([]Token, 0, 16),
		filePos: line.filePos,
	}

	line.skipWhitespace()
	if line.pos > line.len {
		return TokenLine{}
	}

	instr := line.readWord()
	if instr != "" {
		tokens.tokens = append(tokens.tokens, Token{tokenType: INSTRUCTION, value: instr})
	}

	for line.pos < line.len {
		line.skipWhitespace()
		if line.pos < line.len && line.value[line.pos] == ',' {
			tokens.tokens = append(tokens.tokens, Token{tokenType: COMMA, value: ","})
			line.pos++
			continue
		}

		operand := line.readWord()
		if operand != "" {
			tokens.tokens = append(tokens.tokens, Token{tokenType: OPERAND, value: operand})
		}
	}

	return tokens

}

func (tokens TokenLine) refineTokens() TokenLine {
	for index := range tokens.tokens {
		token := &tokens.tokens[index] //Actually modify the token and not just a copy of it.
		switch token.tokenType {
		case COMMA:
			if index == len(tokens.tokens)-1 {
				panic(fmt.Sprintf("Found no operand after a comma (Line: %d Token: %d, After: %+v)", tokens.filePos, index, "End of line"))
			} else if tokens.tokens[index+1].tokenType != OPERAND {
				panic(fmt.Sprintf("Found no operand after a comma (Line: %d Token: %d, After: %+v)", tokens.filePos, index, tokens.tokens[index+1].value))
			}
			//fmt.Printf("%+v (%d) passed\n", token, index)

		case OPERAND:
			if len(token.value) < 2 {
				token.optionalType = UNDEFINED
				continue
			}

			if contains(registerABI, string(token.value)) {
				token.optionalType = REGISTER
				continue
			}

			switch {
			case token.value[:2] == "0x":
				token.optionalType = IMMEDIATEHEX
			case token.value[:2] == "0b":
				token.optionalType = IMMEDIATEBYT
			case token.value[:1] == "x":
				token.optionalType = REGISTER
			default:
				token.optionalType = UNDEFINED
			}

		default:
			//fmt.Printf("%+v (%d) passed\n", token, index)
		}

	}

	return tokens
}

func main() {
	f, err := os.Open("_asm/test.s")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	currentLine := 1
	for scanner.Scan() {
		tokens := Line{value: scanner.Text(), pos: 0, len: uint16(len(scanner.Text())), filePos: uint(currentLine)}.lexeLine()
		tokens = tokens.refineTokens()
		fmt.Printf("%+v\n", tokens)
		currentLine++
	}

	fmt.Println("Lexing was successful")

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
