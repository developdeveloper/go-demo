package main

func Sum() int {
        const count = 100
        numbers := make([]int, count)
        for i := range numbers {
                numbers[i] = i + 1
        }

        var sum int
        for _, i := range numbers {
                sum += i
        }
        return sum
}
