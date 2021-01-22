#include "textflag.h"

TEXT ·add(SB), NOSPLIT, $24-24
  MOVQ a+0(FP), DX
  MOVQ DX, 0(SP)
  MOVQ b+8(FP), CX
  MOVQ CX, 8(SP)
  CALL ·doAdd(SB)
  MOVQ 16(SP), AX     // doadd 函数会返回值在在 16(SP)
  MOVQ AX, ret+16(FP)
  RET
