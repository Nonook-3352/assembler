package main

import (
	"assembler/internal/rvcore"
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := os.Args[1]
	verbose := false
	if rvcore.Contains(os.Args, "--verbose") {
		verbose = true
	}
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	currentLine := 1
	for scanner.Scan() {
		tokens := rvcore.Line{Value: scanner.Text(), Len: uint16(len(scanner.Text())), FilePos: uint(currentLine)}.LexeLine()
		err, tokens := tokens.RefineTokens()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if verbose {
			fmt.Printf("%+v\n", tokens)
		}
		currentLine++
	}

	fmt.Println("===Lexing complete===")

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
