package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	s "strings"
)

const (
	C_ARITHMETIC = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

type Parser struct {
	file    *os.File
	scanner *bufio.Scanner
	command string
	args    []string
}

func NewParser(file *os.File) *Parser {
	scanner := bufio.NewScanner(file)
	p := Parser{file: file, scanner: scanner}
	return &p
}

// Print contents of input file
func (p *Parser) Print() {
	for p.scanner.Scan() {
		fmt.Println(p.scanner.Text())
	}
}

func (p *Parser) PrintCommand() {
	fmt.Println(p.command)
}

// Move to next command line in input and return true. Return false if no more commands
func (p *Parser) HasMoreCommands() bool {
	for p.scanner.Scan() {
		line := p.scanner.Text()
		line = s.TrimSpace(line)
		if s.HasPrefix(line, "//") || line == "" {
			continue
		}
		return true
	}
	return false
}

// Set Command of parse to current line in input
func (p *Parser) Advance() {
	subs := s.SplitN(p.scanner.Text(), "//", 2)
	p.command = s.TrimSpace(subs[0])
	p.args = s.Split(p.command, " ")
}

// Returns the type of the current command
func (p *Parser) CommandType() int {
	switch p.args[0] {
	case
		"add",
		"sub",
		"neg",
		"eq",
		"gt",
		"lt",
		"and",
		"or",
		"not":
		return C_ARITHMETIC
	case "push":
		return C_PUSH
	case "pop":
		return C_POP
	case "label":
		return C_LABEL
	case "goto":
		return C_GOTO
	case "if-goto":
		return C_IF
	case "function":
		return C_FUNCTION
	case "return":
		return C_RETURN
	case "call":
		return C_CALL
	}
	log.Fatal("Command not recognised")
	return -1
}

// Return the first argument of the current command
func (p *Parser) Arg1() string {
	if p.CommandType() == C_ARITHMETIC {
		return p.args[0]
	}
	return p.args[1]
}

// Return the second argument of the current command
func (p *Parser) Arg2() int64 {
	arg, _ := strconv.ParseInt(p.args[2], 10, 64)
	return arg
}
