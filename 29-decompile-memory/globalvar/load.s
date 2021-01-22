#include "textflag.h"

TEXT ·load(SB), NOSPLIT, $0-8
  MOVQ ·num(SB), AX
  MOVQ AX, ret+0(FP)
  RET
