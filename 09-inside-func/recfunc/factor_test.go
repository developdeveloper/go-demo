package recfunc

import (
	"fmt"
	"testing"
)

func Test_factor(t *testing.T) {
	fmt.Println(factor(6))
	fmt.Println(factor2(6))
	fmt.Println(factor2(10))

	//factor2(100000000) // stack overflow
	//tailFactor(100000000-1, 100000000) // stack overflow
}

func Benchmark_factor2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factor2(10)
	}
}

func Test_tailFactor(t *testing.T) {
	fmt.Println(tailFactor(5, 6))
	fmt.Println(tailFactor(9, 10))
}

func Benchmark_tailFactor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tailFactor(9, 10)
	}
}
