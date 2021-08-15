package codewriter

import (
	"bufio"
	"os"
	"strconv"
	s "strings"
	"vmtranslator/parser"
)

const goToNext = `@SP
M=M-1
A=M
`

const incSP = `@SP
M=M+1
`

const arithmeticString = `@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
{{Operation}}
@SP
M=M+1
`

const cmpString = `@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@{{label1}}
D;{{Operation}}
A=0
D=A
@SP
A=M
M=D
@{{label2}}
0;JMP
({{label1}})
A=-1
D=A
@SP
A=M
M=D
({{label2}})
@SP
M=M+1
`

const opStr = "{{Operation}}"

type CodeWriter struct {
	file       *os.File
	writer     *bufio.Writer
	labelCount int
}

func NewCodeWriter(file *os.File) *CodeWriter {
	writer := bufio.NewWriter(file)
	return &CodeWriter{file: file, writer: writer, labelCount: 0}
}

func (cw *CodeWriter) Flush() {
	cw.writer.Flush()
	return
}

func (cw *CodeWriter) WriteEnd() {
	cw.writer.WriteString("(LABEL-END)\n@LABEL-END\n0;JMP\n")
}

func (cw *CodeWriter) WriteArithmetic(command string) {
	cw.writer.WriteString("// " + command + "\n")
	switch command {
	case "add":
		cw.writer.WriteString(s.Replace(arithmeticString, opStr, "M=D+M", 1))
	case "sub":
		cw.writer.WriteString(s.Replace(arithmeticString, opStr, "M=M-D", 1))
	case "neg":
		cw.writer.WriteString(goToNext + "M=-M\n" + incSP)
	case "eq":
		outstr := s.Replace(cmpString, opStr, "JEQ", 1)
		outstr = s.Replace(outstr, "{{label1}}", "LABEL-"+strconv.Itoa(cw.labelCount), -1)
		cw.labelCount += 1
		outstr = s.Replace(outstr, "{{label2}}", "LABEL-"+strconv.Itoa(cw.labelCount), -1)
		cw.labelCount += 1
		cw.writer.WriteString(outstr)
	case "gt":
		outstr := s.Replace(cmpString, opStr, "JLT", 1)
		outstr = s.Replace(outstr, "{{label1}}", "LABEL-"+strconv.Itoa(cw.labelCount), -1)
		cw.labelCount += 1
		outstr = s.Replace(outstr, "{{label2}}", "LABEL-"+strconv.Itoa(cw.labelCount), -1)
		cw.labelCount += 1
		cw.writer.WriteString(outstr)
	case "lt":
		outstr := s.Replace(cmpString, opStr, "JGT", 1)
		outstr = s.Replace(outstr, "{{label1}}", "LABEL-"+strconv.Itoa(cw.labelCount), -1)
		cw.labelCount += 1
		outstr = s.Replace(outstr, "{{label2}}", "LABEL-"+strconv.Itoa(cw.labelCount), -1)
		cw.labelCount += 1
		cw.writer.WriteString(outstr)
	case "and":
		cw.writer.WriteString(s.Replace(arithmeticString, opStr, "M=D&M", 1))
	case "or":
		cw.writer.WriteString(s.Replace(arithmeticString, opStr, "M=D|M", 1))
	case "not":
		cw.writer.WriteString(goToNext + "M=!M\n" + incSP)
	}
	return
}

func (cw *CodeWriter) WritePushPop(cmd int, segment string, index int64) {
	cw.writer.WriteString("// " + strconv.Itoa(cmd) + " " + segment + " " + strconv.FormatInt(index, 10) + "\n")
	switch cmd {
	case parser.C_PUSH:
		cw.writePush(segment, index)
	case parser.C_POP:
		cw.writePop(segment, index)
	}
	return
}

func (cw *CodeWriter) writePush(segment string, index int64) {
	switch segment {
	case "constant":
		cmd := "@" + strconv.FormatInt(index, 10) + "\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"
		cw.writer.WriteString(cmd)
	}
	return
}

func (cw *CodeWriter) writePop(segment string, index int64) {
	return
}
