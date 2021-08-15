package main

import (
	"vmtranslator/codewriter"
	"vmtranslator/parser"

	"log"
	"os"
	s "strings"
)

func main() {
	fname := "./examples/StackTest.vm" // Default input

	// Get input file and open
	if len(os.Args) > 1 {
		fname = os.Args[1]
	}
	infile, err := os.Open(fname)
	check(err)
	defer infile.Close()

	// Create output file and writer
	outfile, err := os.Create(s.TrimSuffix(fname, ".vm") + ".asm")
	check(err)
	defer outfile.Close()

	p := parser.NewParser(infile)
	cw := codewriter.NewCodeWriter(outfile)
	defer cw.Flush()
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.C_ARITHMETIC:
			cw.WriteArithmetic(p.Arg1())
		case parser.C_PUSH, parser.C_POP:
			cw.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
		}
	}
	cw.WriteEnd()

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
