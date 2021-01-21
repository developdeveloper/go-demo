package deferfunc

func f() int {
  i := 0

  defer func(val int) {
    val++
  } (i)

  return i
}
