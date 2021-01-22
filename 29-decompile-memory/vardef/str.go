package vardef

var Str = "hello"

/*
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 76 61 72 64 65 66                                vardef
go.string."hello" SRODATA dupok size=5
	0x0000 68 65 6c 6c 6f                                   hello
"".Str SDATA size=16
	0x0000 00 00 00 00 00 00 00 00 05 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."hello"+0
*/
