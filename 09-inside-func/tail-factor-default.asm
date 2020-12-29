"".main STEXT size=71 args=0x0 locals=0x20
	0x0000 00000 (tail-factor.go:3)	TEXT	"".main(SB), ABIInternal, $32-0
	0x0000 00000 (tail-factor.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (tail-factor.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (tail-factor.go:3)	PCDATA	$0, $-2
	0x000d 00013 (tail-factor.go:3)	JLS	64
	0x000f 00015 (tail-factor.go:3)	PCDATA	$0, $-1
	0x000f 00015 (tail-factor.go:3)	SUBQ	$32, SP
	0x0013 00019 (tail-factor.go:3)	MOVQ	BP, 24(SP)
	0x0018 00024 (tail-factor.go:3)	LEAQ	24(SP), BP
	0x001d 00029 (tail-factor.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (tail-factor.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (tail-factor.go:4)	MOVQ	$9, (SP)
	0x0025 00037 (tail-factor.go:4)	MOVQ	$10, 8(SP)
	0x002e 00046 (tail-factor.go:4)	PCDATA	$1, $0
	0x002e 00046 (tail-factor.go:4)	CALL	"".tailFactor(SB)
	0x0033 00051 (tail-factor.go:5)	MOVQ	24(SP), BP
	0x0038 00056 (tail-factor.go:5)	ADDQ	$32, SP
	0x003c 00060 (tail-factor.go:5)	RET
	0x003d 00061 (tail-factor.go:5)	NOP
	0x003d 00061 (tail-factor.go:3)	PCDATA	$1, $-1
	0x003d 00061 (tail-factor.go:3)	PCDATA	$0, $-2
	0x003d 00061 (tail-factor.go:3)	NOP
	0x0040 00064 (tail-factor.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0045 00069 (tail-factor.go:3)	PCDATA	$0, $-1
	0x0045 00069 (tail-factor.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 31 48  eH..%....H;a.v1H
	0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 c7 04  .. H.l$.H.l$.H..
	0x0020 24 09 00 00 00 48 c7 44 24 08 0a 00 00 00 e8 00  $....H.D$.......
	0x0030 00 00 00 48 8b 6c 24 18 48 83 c4 20 c3 0f 1f 00  ...H.l$.H.. ....
	0x0040 e8 00 00 00 00 eb b9                             .......
	rel 5+4 t=17 TLS+0
	rel 47+4 t=8 "".tailFactor+0
	rel 65+4 t=8 runtime.morestack_noctxt+0
"".tailFactor STEXT size=139 args=0x18 locals=0x28
	0x0000 00000 (tail-factor.go:7)	TEXT	"".tailFactor(SB), ABIInternal, $40-24
	0x0000 00000 (tail-factor.go:7)	MOVQ	(TLS), CX
	0x0009 00009 (tail-factor.go:7)	CMPQ	SP, 16(CX)
	0x000d 00013 (tail-factor.go:7)	PCDATA	$0, $-2
	0x000d 00013 (tail-factor.go:7)	JLS	129
	0x000f 00015 (tail-factor.go:7)	PCDATA	$0, $-1
	0x000f 00015 (tail-factor.go:7)	SUBQ	$40, SP
	0x0013 00019 (tail-factor.go:7)	MOVQ	BP, 32(SP)
	0x0018 00024 (tail-factor.go:7)	LEAQ	32(SP), BP
	0x001d 00029 (tail-factor.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (tail-factor.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (tail-factor.go:7)	MOVQ	$0, "".~r2+64(SP)
	0x0026 00038 (tail-factor.go:8)	CMPQ	"".num+48(SP), $1
	0x002c 00044 (tail-factor.go:8)	JEQ	48
	0x002e 00046 (tail-factor.go:8)	JMP	68
	0x0030 00048 (tail-factor.go:9)	MOVQ	"".current+56(SP), AX
	0x0035 00053 (tail-factor.go:9)	MOVQ	AX, "".~r2+64(SP)
	0x003a 00058 (tail-factor.go:9)	MOVQ	32(SP), BP
	0x003f 00063 (tail-factor.go:9)	ADDQ	$40, SP
	0x0043 00067 (tail-factor.go:9)	RET
	0x0044 00068 (tail-factor.go:12)	MOVQ	"".num+48(SP), AX
	0x0049 00073 (tail-factor.go:12)	DECQ	AX
	0x004c 00076 (tail-factor.go:12)	MOVQ	AX, (SP)
	0x0050 00080 (tail-factor.go:12)	MOVQ	"".num+48(SP), AX
	0x0055 00085 (tail-factor.go:12)	MOVQ	"".current+56(SP), CX
	0x005a 00090 (tail-factor.go:12)	IMULQ	CX, AX
	0x005e 00094 (tail-factor.go:12)	MOVQ	AX, 8(SP)
	0x0063 00099 (tail-factor.go:12)	PCDATA	$1, $0
	0x0063 00099 (tail-factor.go:12)	CALL	"".tailFactor(SB)
	0x0068 00104 (tail-factor.go:12)	MOVQ	16(SP), AX
	0x006d 00109 (tail-factor.go:12)	MOVQ	AX, ""..autotmp_3+24(SP)
	0x0072 00114 (tail-factor.go:12)	MOVQ	AX, "".~r2+64(SP)
	0x0077 00119 (tail-factor.go:12)	MOVQ	32(SP), BP
	0x007c 00124 (tail-factor.go:12)	ADDQ	$40, SP
	0x0080 00128 (tail-factor.go:12)	RET
	0x0081 00129 (tail-factor.go:12)	NOP
	0x0081 00129 (tail-factor.go:7)	PCDATA	$1, $-1
	0x0081 00129 (tail-factor.go:7)	PCDATA	$0, $-2
	0x0081 00129 (tail-factor.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0086 00134 (tail-factor.go:7)	PCDATA	$0, $-1
	0x0086 00134 (tail-factor.go:7)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 72 48  eH..%....H;a.vrH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 44  ..(H.l$ H.l$ H.D
	0x0020 24 40 00 00 00 00 48 83 7c 24 30 01 74 02 eb 14  $@....H.|$0.t...
	0x0030 48 8b 44 24 38 48 89 44 24 40 48 8b 6c 24 20 48  H.D$8H.D$@H.l$ H
	0x0040 83 c4 28 c3 48 8b 44 24 30 48 ff c8 48 89 04 24  ..(.H.D$0H..H..$
	0x0050 48 8b 44 24 30 48 8b 4c 24 38 48 0f af c1 48 89  H.D$0H.L$8H...H.
	0x0060 44 24 08 e8 00 00 00 00 48 8b 44 24 10 48 89 44  D$......H.D$.H.D
	0x0070 24 18 48 89 44 24 40 48 8b 6c 24 20 48 83 c4 28  $.H.D$@H.l$ H..(
	0x0080 c3 e8 00 00 00 00 e9 75 ff ff ff                 .......u...
	rel 5+4 t=17 TLS+0
	rel 100+4 t=8 "".tailFactor+0
	rel 130+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
