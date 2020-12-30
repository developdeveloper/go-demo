package duck

import "testing"

func Test_duck(t *testing.T) {
	d := Duck{"family", "kity", 50, 100}
	d.Walk(100)
	d.Swim(100)
}
