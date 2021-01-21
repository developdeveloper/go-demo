#include "textflag.h"

TEXT ·add(SB), NOSPLIT, $0-24
  MOVQ b+16(SP), AX
  MOVQ a+8(SP), BX
  ADDQ BX, AX
  MOVQ AX, ret + 24(SP)
  RET
