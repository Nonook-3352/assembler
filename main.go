package main

import (
	"assembler/internal/rvcore"
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("_asm/test.s")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	currentLine := 1
	for scanner.Scan() {
		tokens := rvcore.Line{Value: scanner.Text(), Len: uint16(len(scanner.Text())), FilePos: uint(currentLine)}.LexeLine()
		tokens = tokens.RefineTokens()
		fmt.Printf("%+v\n", tokens)
		currentLine++
	}

	fmt.Println("Lexing was successful")

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
