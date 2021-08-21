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
	var outfileName string
	if fname == arg {
		outfileName = filepath.Join(fname, filepath.Base(fname)) + ".asm"
	} else {
		outfileName = fname + ".asm"
	}

	// Create output file and writer
	outfile, err := os.Create(outfileName)
	check(err)
	defer outfile.Close()

	cw := codewriter.NewCodeWriter(outfile)
	defer cw.Flush()

	if fname == arg {
		// Run on files in directory
		files, err := ioutil.ReadDir(fname)
		check(err)
		for _, file := range files {
			if isVM(filepath.Join(fname, file.Name())) {
				parseFile(filepath.Join(fname, file.Name()), cw)
			}
		}
	} else if isVM(arg) {
		parseFile(arg, cw)
	} else {
		log.Fatal("Argument was not a vm file or directory.")
	}

	// cw.WriteEnd()

}

func parseFile(fname string, cw *codewriter.CodeWriter) {
	infile, err := os.Open(fname)
	check(err)
	defer infile.Close()

	cw.SetCurrentFile(fname)
	p := parser.NewParser(infile)

	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.C_ARITHMETIC:
			cw.WriteArithmetic(p.Arg1())
		case parser.C_PUSH, parser.C_POP:
			cw.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
		case parser.C_LABEL:
			cw.WriteLabel(p.Arg1())
		case parser.C_GOTO:
			cw.WriteGoto(p.Arg1())
		case parser.C_IF:
			cw.WriteIfGoto(p.Arg1())
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
