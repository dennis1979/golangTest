package mytest

import (
	"fmt"
	"unsafe"
)

const (
	d = "abc"
	e = len(d)
	f = unsafe.Sizeof(d)
)

func VarTest() {
	// variable and const
	println("=====================Var & Const Test=====================")
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("area : %d", area)
	println(a, b, c)
}

func EnumTest() {
	// enum
	println("=====================Enum Test=====================")
	println(d, e, f)
}
