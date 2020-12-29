"".main STEXT size=59 args=0x0 locals=0x18
	0x0000 00000 (factor2.go:3)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (factor2.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (factor2.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (factor2.go:3)	PCDATA	$0, $-2
	0x000d 00013 (factor2.go:3)	JLS	52
	0x000f 00015 (factor2.go:3)	PCDATA	$0, $-1
	0x000f 00015 (factor2.go:3)	SUBQ	$24, SP
	0x0013 00019 (factor2.go:3)	MOVQ	BP, 16(SP)
	0x0018 00024 (factor2.go:3)	LEAQ	16(SP), BP
	0x001d 00029 (factor2.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (factor2.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (factor2.go:4)	MOVQ	$10, (SP)
	0x0025 00037 (factor2.go:4)	PCDATA	$1, $0
	0x0025 00037 (factor2.go:4)	CALL	"".factor2(SB)
	0x002a 00042 (factor2.go:5)	MOVQ	16(SP), BP
	0x002f 00047 (factor2.go:5)	ADDQ	$24, SP
	0x0033 00051 (factor2.go:5)	RET
	0x0034 00052 (factor2.go:5)	NOP
	0x0034 00052 (factor2.go:3)	PCDATA	$1, $-1
	0x0034 00052 (factor2.go:3)	PCDATA	$0, $-2
	0x0034 00052 (factor2.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0039 00057 (factor2.go:3)	PCDATA	$0, $-1
	0x0039 00057 (factor2.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 25 48  eH..%....H;a.v%H
	0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 c7 04  ...H.l$.H.l$.H..
	0x0020 24 0a 00 00 00 e8 00 00 00 00 48 8b 6c 24 10 48  $.........H.l$.H
	0x0030 83 c4 18 c3 e8 00 00 00 00 eb c5                 ...........
	rel 5+4 t=17 TLS+0
	rel 38+4 t=8 "".factor2+0
	rel 53+4 t=8 runtime.morestack_noctxt+0
"".factor2 STEXT size=125 args=0x10 locals=0x20
	0x0000 00000 (factor2.go:7)	TEXT	"".factor2(SB), ABIInternal, $32-16
	0x0000 00000 (factor2.go:7)	MOVQ	(TLS), CX
	0x0009 00009 (factor2.go:7)	CMPQ	SP, 16(CX)
	0x000d 00013 (factor2.go:7)	PCDATA	$0, $-2
	0x000d 00013 (factor2.go:7)	JLS	118
	0x000f 00015 (factor2.go:7)	PCDATA	$0, $-1
	0x000f 00015 (factor2.go:7)	SUBQ	$32, SP
	0x0013 00019 (factor2.go:7)	MOVQ	BP, 24(SP)
	0x0018 00024 (factor2.go:7)	LEAQ	24(SP), BP
	0x001d 00029 (factor2.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (factor2.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (factor2.go:7)	MOVQ	$0, "".~r1+48(SP)
	0x0026 00038 (factor2.go:8)	CMPQ	"".num+40(SP), $1
	0x002c 00044 (factor2.go:8)	JEQ	48
	0x002e 00046 (factor2.go:8)	JMP	67
	0x0030 00048 (factor2.go:9)	MOVQ	$1, "".~r1+48(SP)
	0x0039 00057 (factor2.go:9)	MOVQ	24(SP), BP
	0x003e 00062 (factor2.go:9)	ADDQ	$32, SP
	0x0042 00066 (factor2.go:9)	RET
	0x0043 00067 (factor2.go:12)	MOVQ	"".num+40(SP), AX
	0x0048 00072 (factor2.go:12)	DECQ	AX
	0x004b 00075 (factor2.go:12)	MOVQ	AX, (SP)
	0x004f 00079 (factor2.go:12)	PCDATA	$1, $0
	0x004f 00079 (factor2.go:12)	CALL	"".factor2(SB)
	0x0054 00084 (factor2.go:12)	MOVQ	8(SP), AX
	0x0059 00089 (factor2.go:12)	MOVQ	AX, ""..autotmp_2+16(SP)
	0x005e 00094 (factor2.go:12)	MOVQ	"".num+40(SP), CX
	0x0063 00099 (factor2.go:12)	IMULQ	AX, CX
	0x0067 00103 (factor2.go:12)	MOVQ	CX, "".~r1+48(SP)
	0x006c 00108 (factor2.go:12)	MOVQ	24(SP), BP
	0x0071 00113 (factor2.go:12)	ADDQ	$32, SP
	0x0075 00117 (factor2.go:12)	RET
	0x0076 00118 (factor2.go:12)	NOP
	0x0076 00118 (factor2.go:7)	PCDATA	$1, $-1
	0x0076 00118 (factor2.go:7)	PCDATA	$0, $-2
	0x0076 00118 (factor2.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x007b 00123 (factor2.go:7)	PCDATA	$0, $-1
	0x007b 00123 (factor2.go:7)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 67 48  eH..%....H;a.vgH
	0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 c7 44  .. H.l$.H.l$.H.D
	0x0020 24 30 00 00 00 00 48 83 7c 24 28 01 74 02 eb 13  $0....H.|$(.t...
	0x0030 48 c7 44 24 30 01 00 00 00 48 8b 6c 24 18 48 83  H.D$0....H.l$.H.
	0x0040 c4 20 c3 48 8b 44 24 28 48 ff c8 48 89 04 24 e8  . .H.D$(H..H..$.
	0x0050 00 00 00 00 48 8b 44 24 08 48 89 44 24 10 48 8b  ....H.D$.H.D$.H.
	0x0060 4c 24 28 48 0f af c8 48 89 4c 24 30 48 8b 6c 24  L$(H...H.L$0H.l$
	0x0070 18 48 83 c4 20 c3 e8 00 00 00 00 eb 83           .H.. ........
	rel 5+4 t=17 TLS+0
	rel 80+4 t=8 "".factor2+0
	rel 119+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
