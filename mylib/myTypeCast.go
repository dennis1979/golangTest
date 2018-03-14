package mytest

import (
	"fmt"
)

func TypeCastTest() {
	println("=====================Type Cast Test=====================")
	var sum int = 17
	var count int = 5
	var mean float32
	fmt.Println("sum: %d, count: %d", sum, count)
	mean = float32(sum) / float32(count)
	fmt.Printf("type cast int to float, mean: %f\n", mean)
}
