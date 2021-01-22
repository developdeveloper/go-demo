#include "textflag.h"

TEXT ·output(SB), $8-48
  MOVQ 24(SP), CX         // 硬件寄存器 SP
  MOVQ CX, ret3+24(FP)    // 第三个返回值

  MOVQ arg1+16(SP), BX    // 栈空间是 8，+16 后第第一个返回值
  MOVQ BX, ret2+16(FP)    // 第二个返回值

  MOVQ arg1+0(FP), AX
  MOVQ AX, ret1+8(FP)     // 第一个返回值

  RET
