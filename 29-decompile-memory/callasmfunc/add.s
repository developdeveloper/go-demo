#include "textflag.h"

TEXT ·add(SB), NOSPLIT, $0-24
  MOVQ b+16(SP), AX     ; MOVQ b+8(FP), AX
  MOVQ a+8(SP), BX      ; MOVQ a+0(FP), BX
  ADDQ BX, AX
  MOVQ AX, ret + 24(SP) ; MOVQ AX, ret + 16(FP)
  RET
