package main

import (
	"io/ioutil"
	"vmtranslator/codewriter"
	"vmtranslator/parser"

	"log"
	"os"
	"path/filepath"
	s "strings"
)

func main() {
	arg := "./examples/StackTest.vm" // Default input

	// Get input file and open
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	fname := s.TrimSuffix(arg, ".vm")

	// Create output file and writer
	outfile, err := os.Create(fname + ".asm")
	check(err)
	defer outfile.Close()

	cw := codewriter.NewCodeWriter(outfile)

	if fname == arg {
		// Run on files in directory
		files, err := ioutil.ReadDir(fname)
		check(err)
		for _, file := range files {
			if isVM(file.Name()) {
				parseFile(file.Name(), cw)
			}
		}
	} else if isVM(arg) {
		parseFile(arg, cw)
	} else {
		log.Fatal("Argument was not a vm file or directory.")
	}

	cw.WriteEnd()

}

func parseFile(fname string, cw *codewriter.CodeWriter) {
	infile, err := os.Open(fname)
	check(err)
	defer infile.Close()

	p := parser.NewParser(infile)

	defer cw.Flush()
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.C_ARITHMETIC:
			cw.WriteArithmetic(p.Arg1())
		case parser.C_PUSH, parser.C_POP:
			cw.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
		case parser.C_LABEL:
			cw.WriteLabel(p.Arg1())
		}
	}
}

func isVM(file string) bool {
	return filepath.Ext(file) == ".vm"
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
