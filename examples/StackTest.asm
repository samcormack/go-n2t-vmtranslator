// 1 constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@LABEL-0
D;JEQ
A=0
D=A
@SP
A=M
M=D
@LABEL-1
0;JMP
(LABEL-0)
A=-1
D=A
@SP
A=M
M=D
(LABEL-1)
@SP
M=M+1
// 1 constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@LABEL-2
D;JEQ
A=0
D=A
@SP
A=M
M=D
@LABEL-3
0;JMP
(LABEL-2)
A=-1
D=A
@SP
A=M
M=D
(LABEL-3)
@SP
M=M+1
// 1 constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@LABEL-4
D;JEQ
A=0
D=A
@SP
A=M
M=D
@LABEL-5
0;JMP
(LABEL-4)
A=-1
D=A
@SP
A=M
M=D
(LABEL-5)
@SP
M=M+1
// 1 constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@LABEL-6
D;JGT
A=0
D=A
@SP
A=M
M=D
@LABEL-7
0;JMP
(LABEL-6)
A=-1
D=A
@SP
A=M
M=D
(LABEL-7)
@SP
M=M+1
// 1 constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@LABEL-8
D;JGT
A=0
D=A
@SP
A=M
M=D
@LABEL-9
0;JMP
(LABEL-8)
A=-1
D=A
@SP
A=M
M=D
(LABEL-9)
@SP
M=M+1
// 1 constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@LABEL-10
D;JGT
A=0
D=A
@SP
A=M
M=D
@LABEL-11
0;JMP
(LABEL-10)
A=-1
D=A
@SP
A=M
M=D
(LABEL-11)
@SP
M=M+1
// 1 constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@LABEL-12
D;JLT
A=0
D=A
@SP
A=M
M=D
@LABEL-13
0;JMP
(LABEL-12)
A=-1
D=A
@SP
A=M
M=D
(LABEL-13)
@SP
M=M+1
// 1 constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@LABEL-14
D;JLT
A=0
D=A
@SP
A=M
M=D
@LABEL-15
0;JMP
(LABEL-14)
A=-1
D=A
@SP
A=M
M=D
(LABEL-15)
@SP
M=M+1
// 1 constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=D-M
@LABEL-16
D;JLT
A=0
D=A
@SP
A=M
M=D
@LABEL-17
0;JMP
(LABEL-16)
A=-1
D=A
@SP
A=M
M=D
(LABEL-17)
@SP
M=M+1
// 1 constant 57
@57
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 31
@31
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 constant 53
@53
D=A
@SP
A=M
M=D
@SP
M=M+1
// add
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=D+M
@SP
M=M+1
// 1 constant 112
@112
D=A
@SP
A=M
M=D
@SP
M=M+1
// sub
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=M-D
@SP
M=M+1
// neg
@SP
M=M-1
A=M
M=-M
@SP
M=M+1
// and
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=D&M
@SP
M=M+1
// 1 constant 82
@82
D=A
@SP
A=M
M=D
@SP
M=M+1
// or
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=D|M
@SP
M=M+1
// not
@SP
M=M-1
A=M
M=!M
@SP
M=M+1
(LABEL-END)
@LABEL-END
0;JMP