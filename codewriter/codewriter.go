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
	file            *os.File
	writer          *bufio.Writer
	labelCount      int
	staticCount     int
	callCount       int
	currentReadFile string
	currentFunction string
}

func NewCodeWriter(file *os.File) *CodeWriter {
	writer := bufio.NewWriter(file)
	return &CodeWriter{
		file:            file,
		writer:          writer,
		labelCount:      0,
		staticCount:     0,
		callCount:       0,
		currentReadFile: getFilename(file.Name()),
		currentFunction: "main",
	}
}

func (cw *CodeWriter) SetCurrentFile(fname string) {
	cw.currentReadFile = getFilename(fname)
}

func (cw *CodeWriter) Flush() {
	cw.writer.Flush()
	return
}

func (cw *CodeWriter) WriteInit() {

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
	cw.writer.WriteString("// " + commandName(cmd) + " " + segment + " " + strconv.FormatInt(index, 10) + "\n")
	switch cmd {
	case parser.C_PUSH:
		cw.writePush(segment, index)
	case parser.C_POP:
		cw.writePop(segment, index)
	}
	return
}

func commandName(cmd int) string {
	switch cmd {
	case parser.C_PUSH:
		return "push"
	case parser.C_POP:
		return "pop"
	}
	return ""
}

func (cw *CodeWriter) writePush(segment string, index int64) {
	var cmd string
	pushStr := "\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"
	switch segment {
	case "constant":
		cmd = "@" + strconv.FormatInt(index, 10) + "\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"
	case "local":
		cmd = "@LCL\nD=M\n@" + strconv.FormatInt(index, 10) + pushStr
	case "argument":
		cmd = "@ARG\nD=M\n@" + strconv.FormatInt(index, 10) + pushStr
	case "this":
		cmd = "@THIS\nD=M\n@" + strconv.FormatInt(index, 10) + pushStr
	case "that":
		cmd = "@THAT\nD=M\n@" + strconv.FormatInt(index, 10) + pushStr
	case "temp":
		cmd = "@5\nD=A\n@" + strconv.FormatInt(index, 10) + pushStr
	case "static":
		cmd = "@" + cw.currentReadFile + "." + strconv.FormatInt(index, 10) + "\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"
	case "pointer":
		if index == 0 {
			cmd = "@THIS\n"
		} else if index == 1 {
			cmd = "@THAT\n"
		}
		cmd += "D=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"
	}
	cw.writer.WriteString(cmd)
	return
}

func (cw *CodeWriter) writePop(segment string, index int64) {
	var cmd string
	popStr := "\nD=D+A\n@R13\nM=D\n@SP\nM=M-1\nA=M\nD=M\n@R13\nA=M\nM=D\n"
	switch segment {
	case "local":
		cmd = "@LCL\nD=M\n@" + strconv.FormatInt(index, 10) + popStr
	case "argument":
		cmd = "@ARG\nD=M\n@" + strconv.FormatInt(index, 10) + popStr
	case "this":
		cmd = "@THIS\nD=M\n@" + strconv.FormatInt(index, 10) + popStr
	case "that":
		cmd = "@THAT\nD=M\n@" + strconv.FormatInt(index, 10) + popStr
	case "temp":
		cmd = "@5\nD=A\n@" + strconv.FormatInt(index, 10) + popStr
	case "static":
		cmd = "@SP\nM=M-1\nA=M\nD=M\n" + "@" + cw.currentReadFile + "." + strconv.FormatInt(index, 10) + "\nM=D\n"
	case "pointer":
		if index == 0 {
			cmd = "@SP\nM=M-1\nA=M\nD=M\n@THIS\nM=D\n"
		} else if index == 1 {
			cmd = "@SP\nM=M-1\nA=M\nD=M\n@THAT\nM=D\n"
		}
	}
	cw.writer.WriteString(cmd)
	return
}

func (cw *CodeWriter) WriteLabel(arg string) {
	cw.writer.WriteString("// label " + arg + "\n")
	cw.writer.WriteString("(" + cw.genLabel(arg) + ")\n")
}

func (cw *CodeWriter) WriteGoto(arg string) {
	cw.writer.WriteString("// goto " + arg + "\n")
	cw.writer.WriteString("@" + cw.genLabel(arg) + "\n0;JMP\n")
}

func (cw *CodeWriter) WriteIfGoto(arg string) {
	cw.writer.WriteString("// if-goto " + arg + "\n")
	cw.writer.WriteString(goToNext)
	cw.writer.WriteString("D=M\n")
	cw.writer.WriteString("@" + cw.genLabel(arg) + "\nD;JNE\n")
}

func (cw *CodeWriter) WriteCall(functionName string, numArgs int64) {
	cw.writer.WriteString("// call" + functionName + strconv.FormatInt(numArgs, 10) + "\n")
	// Save return address
	cw.writer.WriteString("@" + cw.getReturnAddr() + "\nD=A\n")
	cw.writePushD()
	// Save memory segment locations
	cw.writer.WriteString("@LCL\nD=A\n")
	cw.writePushD()
	cw.writer.WriteString("@ARG\nD=A\n")
	cw.writePushD()
	cw.writer.WriteString("@THIS\nD=A\n")
	cw.writePushD()
	cw.writer.WriteString("@THAT\nD=A\n")
	cw.writePushD()
	//Reposition ARG
	cw.writer.WriteString("@SP\nD=A\n")
	cw.writePushD()
	cw.writePush("constant", numArgs)
	cw.WriteArithmetic("sub")
	cw.writePush("constant", 5)
	cw.WriteArithmetic("sub")
	cw.writePopD()
	cw.writer.WriteString("@ARG\nM=D\n")
	// Reposition LCL
	cw.writer.WriteString("@SP\nD=A\n")
	cw.writer.WriteString("@LCL\nM=D\n")
	// go to function
	cw.writer.WriteString("@" + functionName + "\n0:JMP\n")
	// Write return address label
	cw.writer.WriteString("(" + cw.getReturnAddr() + ")")
	cw.callCount += 1
}

func (cw *CodeWriter) WriteReturn() {

}

func (cw *CodeWriter) WriteFunction(functionName string, numLocals int64) {
	cw.currentFunction = functionName
}

func (cw *CodeWriter) genLabel(raw string) string {
	return cw.currentReadFile + "." + cw.currentFunction + "$" + raw
}

func (cw *CodeWriter) getReturnAddr() {
	return "return-address-" + strconv.FormatInt(cw.callCount, 10)
}

func (cw *CodeWriter) writePushD() {
	cw.CodeWriter.WriteString("@SP\nA=M\nM=D\n@SP\nM=M+1\n")
}

func (cw *CodeWriter) writePopD() {
	cw.writer.WriteString("@SP\nM=M-1\nA=M\nD=M\n")
}

func getFilename(path string) string {
	fnameParts := s.Split(path, "/")
	fnameParts = s.Split(fnameParts[len(fnameParts)-1], ".")
	return fnameParts[0]
}
