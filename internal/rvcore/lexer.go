package rvcore

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

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (line *Line) skipWhitespace() {
	for line.pos < line.Len && (line.Value[line.pos] == ' ' || line.Value[line.pos] == '\t') {
		line.pos++
	}
}

func (line *Line) readWord() TokenValue {
	wordStart := line.pos
	for line.pos < line.Len && (line.Value[line.pos] != ' ' && line.Value[line.pos] != ',') {
		line.pos++
	}

	return TokenValue(line.Value[wordStart:line.pos])
}

func (line Line) LexeLine() TokenLine {
	tokens := TokenLine{
		Tokens:  make([]Token, 0, 16),
		FilePos: line.FilePos,
	}

	line.skipWhitespace()
	if line.pos > line.Len {
		return TokenLine{}
	}

	if line.Len == 0 {
		return TokenLine{}
	}

	switch line.Value[line.pos] {
	case '.':
		directive := line.readWord()
		tokens.Tokens = append(tokens.Tokens, Token{TokenType: DIRECTIVE, Value: directive})

		for line.pos < line.Len {
			line.skipWhitespace()
			if line.pos < line.Len && line.Value[line.pos] == ',' {
				tokens.Tokens = append(tokens.Tokens, Token{TokenType: COMMA, Value: ","})
				line.pos++
				continue
			}

			operand := line.readWord()
			if operand != "" {
				tokens.Tokens = append(tokens.Tokens, Token{TokenType: OPERAND, Value: operand})
			}
		}

	default:
		instr := line.readWord()
		if instr != "" {
			tokens.Tokens = append(tokens.Tokens, Token{TokenType: INSTRUCTION, Value: instr})
		}

		for line.pos < line.Len {
			line.skipWhitespace()
			if line.pos < line.Len && line.Value[line.pos] == ',' {
				tokens.Tokens = append(tokens.Tokens, Token{TokenType: COMMA, Value: ","})
				line.pos++
				continue
			}

			operand := line.readWord()
			if operand != "" {
				tokens.Tokens = append(tokens.Tokens, Token{TokenType: OPERAND, Value: operand})
			}
		}
	}

	return tokens

}

func (tokens TokenLine) RefineTokens() (error, TokenLine) {
	for index := range tokens.Tokens {
		token := &tokens.Tokens[index] //Actually modify the token and not just a copy of it.
		switch token.TokenType {
		case COMMA:
			if index == len(tokens.Tokens)-1 {
				return ParseError{
					Line:    tokens.FilePos,
					Token:   uint(index),
					Message: "Found no operand after a comma and reached end of line",
				}, TokenLine{}
			} else if tokens.Tokens[index+1].TokenType != OPERAND {
				return ParseError{
					Line:    tokens.FilePos,
					Token:   uint(index),
					Message: "Found no operand after a comma",
				}, TokenLine{}
			}
			token.OptionalType = UNDEFINED
			//fmt.Printf("%+v (%d) passed\n", token, index)

		case OPERAND:
			if len(token.Value) < 2 {
				token.OptionalType = UNDEFINED
				continue
			}

			if Contains(registerABI, string(token.Value)) {
				token.OptionalType = REGISTER
				continue
			}

			switch {
			case token.Value[:2] == "0x":
				token.OptionalType = IMMEDIATEHEX
			case token.Value[:2] == "0b":
				token.OptionalType = IMMEDIATEBYT
			case token.Value[:1] == "x":
				token.OptionalType = REGISTER
			default:
				token.OptionalType = UNDEFINED
			}

		case DIRECTIVE:
			token.OptionalType = UNDEFINED
			if token.Value == ".arch" {
				if len(tokens.Tokens) > index+1 {
					if tokens.Tokens[index+1].Value != "RISCV32I" {
						return ParseError{
							Line:    tokens.FilePos,
							Token:   uint(index + 1),
							Message: "Only .arch RISCV32I is supported",
						}, TokenLine{}
					}
				}
			}
		case INSTRUCTION:
			token.OptionalType = UNDEFINED
		default:
			//fmt.Printf("%+v (%d) passed\n", token, index)
		}

	}

	return nil, tokens
}
