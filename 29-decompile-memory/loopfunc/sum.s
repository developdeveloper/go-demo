#include "textflag.h"

// func sum(sl []int) int
TEXT ·sum(SB), NOSPLIT, $0-32
  MOVQ $0, SI
  MOVQ sl+0(FP), BX // &sl[0], addr of the first elem
  MOVQ sl+8(FP), CX // len(sl)
  INCQ CX           // CX++, 因为要循环 len 次

start:
  DECQ CX           // CX--
  JZ   done
  ADDQ (BX), SI     // SI += *BX
  ADDQ $8, BX       // 指针移动
  JMP  start

done:
  // 返回地址是 24 是怎么得来的呢?
  // 可以通过 go tool compile -S main.go 得知
  // 在调用 sum 函数时，会传入 reflect.SliceHeader 三个值 Data、Len、Cap
  // 求和只需要 Len，但 Cap 依然会占用空间
  // 就是 16(FP)
  MOVQ SI, ret+24(FP)
  RET
