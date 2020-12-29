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
"".tailFactor STEXT size=114 args=0x18 locals=0x20
	0x0000 00000 (tail-factor.go:7)	TEXT	"".tailFactor(SB), ABIInternal, $32-24
	0x0000 00000 (tail-factor.go:7)	MOVQ	(TLS), CX
	0x0009 00009 (tail-factor.go:7)	CMPQ	SP, 16(CX)
	0x000d 00013 (tail-factor.go:7)	PCDATA	$0, $-2
	0x000d 00013 (tail-factor.go:7)	JLS	107
	0x000f 00015 (tail-factor.go:7)	PCDATA	$0, $-1
	0x000f 00015 (tail-factor.go:7)	SUBQ	$32, SP
	0x0013 00019 (tail-factor.go:7)	MOVQ	BP, 24(SP)
	0x0018 00024 (tail-factor.go:7)	LEAQ	24(SP), BP
	0x001d 00029 (tail-factor.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (tail-factor.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (tail-factor.go:8)	MOVQ	"".num+40(SP), AX
	0x0022 00034 (tail-factor.go:8)	CMPQ	AX, $1
	0x0026 00038 (tail-factor.go:8)	JNE	60
	0x0028 00040 (tail-factor.go:9)	MOVQ	"".current+48(SP), AX
	0x002d 00045 (tail-factor.go:9)	MOVQ	AX, "".~r2+56(SP)
	0x0032 00050 (tail-factor.go:9)	MOVQ	24(SP), BP
	0x0037 00055 (tail-factor.go:9)	ADDQ	$32, SP
	0x003b 00059 (tail-factor.go:9)	RET
	0x003c 00060 (tail-factor.go:12)	LEAQ	-1(AX), CX
	0x0040 00064 (tail-factor.go:12)	MOVQ	CX, (SP)
	0x0044 00068 (tail-factor.go:12)	MOVQ	"".current+48(SP), CX
	0x0049 00073 (tail-factor.go:12)	IMULQ	CX, AX
	0x004d 00077 (tail-factor.go:12)	MOVQ	AX, 8(SP)
	0x0052 00082 (tail-factor.go:12)	PCDATA	$1, $0
	0x0052 00082 (tail-factor.go:12)	CALL	"".tailFactor(SB)
	0x0057 00087 (tail-factor.go:12)	MOVQ	16(SP), AX
	0x005c 00092 (tail-factor.go:12)	MOVQ	AX, "".~r2+56(SP)
	0x0061 00097 (tail-factor.go:12)	MOVQ	24(SP), BP
	0x0066 00102 (tail-factor.go:12)	ADDQ	$32, SP
	0x006a 00106 (tail-factor.go:12)	RET
	0x006b 00107 (tail-factor.go:12)	NOP
	0x006b 00107 (tail-factor.go:7)	PCDATA	$1, $-1
	0x006b 00107 (tail-factor.go:7)	PCDATA	$0, $-2
	0x006b 00107 (tail-factor.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0070 00112 (tail-factor.go:7)	PCDATA	$0, $-1
	0x0070 00112 (tail-factor.go:7)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 5c 48  eH..%....H;a.v\H
	0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 8b 44  .. H.l$.H.l$.H.D
	0x0020 24 28 48 83 f8 01 75 14 48 8b 44 24 30 48 89 44  $(H...u.H.D$0H.D
	0x0030 24 38 48 8b 6c 24 18 48 83 c4 20 c3 48 8d 48 ff  $8H.l$.H.. .H.H.
	0x0040 48 89 0c 24 48 8b 4c 24 30 48 0f af c1 48 89 44  H..$H.L$0H...H.D
	0x0050 24 08 e8 00 00 00 00 48 8b 44 24 10 48 89 44 24  $......H.D$.H.D$
	0x0060 38 48 8b 6c 24 18 48 83 c4 20 c3 e8 00 00 00 00  8H.l$.H.. ......
	0x0070 eb 8e                                            ..
	rel 5+4 t=17 TLS+0
	rel 83+4 t=8 "".tailFactor+0
	rel 108+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
