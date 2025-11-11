package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Nonook-3352/assembler/internal/rvcore"
)

func main() {
	filePath := os.Args[1]
	verbose := false
	if rvcore.Contains(os.Args, "--verbose") {
		verbose = true
	}
	startTime := time.Now()
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	currentLine := 1
	fmt.Println("Lexing started")

	for scanner.Scan() {
		tokens := rvcore.Line{Value: scanner.Text(), Len: uint16(len(scanner.Text())), FilePos: uint(currentLine)}.LexeLine()
		tokens, err := tokens.RefineTokens()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		decoded, err := tokens.Decode()
		if err != nil {
			panic(err)
		}
		output := decoded.EmitAsmLine()

		if verbose {
			fmt.Printf("Tokens: %+v\n", tokens)
			fmt.Printf("DecodedTokens: %+v\n", tokens)
			fmt.Printf("Output: %032b\n\n", output)
		}
		currentLine++
	}

	fmt.Printf("Lexing completed in %s", time.Since(startTime).String())

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
