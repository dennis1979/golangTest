package mytest

import (
	"fmt"
)

func OpTestArithmetic() {
	println("=====================算數符運算=====================")
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}

func OpTestRelation() {
	println("=====================關係運算=====================")
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if a > b {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}
}

func OpTestBoolean() {
	println("=====================邏輯運算=====================")
	var boola bool = true
	var boolb bool = false
	if boola && boolb {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if boola || boolb {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	boola = false
	boolb = true
	if boola && boolb {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(boola && boolb) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}

func OpTestBit() {
	println("=====================位運算=====================")
	var bita uint = 60 /* 60 = 0011 1100 */
	var bitb uint = 13 /* 13 = 0000 1101 */
	var bitc uint = 0

	bitc = bita & bitb /* 12 = 0000 1100 */
	fmt.Printf("第一行 - bitc 的值为 %d\n", bitc)

	bitc = bita | bitb /* 61 = 0011 1101 */
	fmt.Printf("第二行 - bitc 的值为 %d\n", bitc)

	bitc = bita ^ bitb /* 49 = 0011 0001 */
	fmt.Printf("第三行 - bitc 的值为 %d\n", bitc)

	bitc = bita << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - bitc 的值为 %d\n", bitc)

	bitc = bita >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - bitc 的值为 %d\n", bitc)

}

func OpTestAssign() {
	println("=====================運算符運算=====================")
	var a int = 21
	var c int

	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)

	c = 200

	c <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)
}

func OpTestPriority() {
	println("=====================運算符優先級=====================")
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int

	e = (a + b) * c / d // ( 30 * 15 ) / 5
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)

	e = ((a + b) * c) / d // (30 * 15 ) / 5
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n", e)

	e = (a + b) * (c / d) // (30) * (15/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)

	e = a + (b*c)/d //  20 + (150/5)
	fmt.Printf("a + (b * c) / d 的值为  : %d\n", e)
}

func OpTestOther() {
	println("=====================其他運算符=====================")
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
