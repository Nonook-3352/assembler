package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var filePath string
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	output, err := os.Create("output.out")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	reader := csv.NewReader(file)
	for record, err := reader.Read(); err != io.EOF; record, err = reader.Read() {
		if record[0] == "op" {
			continue
		}

		//fmt.Printf("%v", record)
		null := "-"
		opcode := record[0]
		funct3 := record[1]
		funct7 := record[2]
		optype := record[3]
		instr := strings.Fields(record[4])[0]

		if funct3 == null {
			funct3 = "0"
		}
		if funct7 == null {
			funct7 = "0"
		}

		line := fmt.Sprintf("\"%s\": {Format: %sTYPE, Opcode: 0b%s, Funct3: 0b%s, Funct7: 0b%s},\n", instr, optype, opcode, funct3, funct7)

		_, err := output.WriteString(line)
		if err != nil {
			panic(err)
		}
	}

}
