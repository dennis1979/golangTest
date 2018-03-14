package mytest

import (
	"fmt"
	"math"
)

func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func FuncTest() {
	println("=====================function return=====================")
	a := 1010
	b := 2020
	fmt.Println(a, b)
	maxval := max(a, b)
	println("get max val")
	fmt.Println(maxval)
}

func swap(x, y string) (string, string) {
	return y, x
}

func FuncMultiReturnTest() {
	println("=====================function multireturn=====================")
	a := "Mahesh"
	b := "Kumar"
	fmt.Println(a, b)
	a, b = swap(a, b)
	println("after swap")
	fmt.Println(a, b)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func FuncRecurTest() {
	println("=====================Function Recursive Test=====================")
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}

func FuncArrayParamTest() {
	println("=====================Function Array Param Test=====================")
	/* 数组长度为 5 */
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f\n ", avg)
}

func FuncDelegateTest() {
	println("=====================Function Delegate Test=====================")
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
}

func getSequence2() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func FuncClosureTest() {
	println("=====================Function Closure Test=====================")
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence2()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence2()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
