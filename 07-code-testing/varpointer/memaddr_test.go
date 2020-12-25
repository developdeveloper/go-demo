package varpointer

import "testing"

func Test_incUpdateScore(t *testing.T) {
	t.Run("should increment score by + number", func(t *testing.T)  {
    score := 10
    expected := 11
    incUpdateScore(&score, 1)

    if score != expected {
      t.Errorf("有问题, 期望 %d, 实际是 %d", expected, score)
    }
  })

  t.Run("should increment score by - number", func(t *testing.T)  {
    score := 10
    expected := 9
    incUpdateScore(&score, -1)

    if score != expected {
      t.Errorf("有问题, 期望 %d, 实际是 %d", expected, score)
    }
  })
}
