package mytest

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func SliceTest() {
	println("=====================Slice Test=====================")
	var numbers = make([]int, 3, 5)

	printSlice(numbers)
	//
	var numbers2 []int

	printSlice(numbers2)

	if numbers2 == nil {
		fmt.Printf("切片是空的")

		//
		/* 创建切片 */
		numbers3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		printSlice(numbers3)

		/* 打印原始切片 */
		fmt.Println("numbers3 ==", numbers3)

		/* 打印子切片从索引1(包含) 到索引4(不包含)*/
		fmt.Println("numbers3[1:4] ==", numbers3[1:4])

		/* 默认下限为 0*/
		fmt.Println("numbers3[:3] ==", numbers3[:3])

		/* 默认上限为 len(s)*/
		fmt.Println("numbers3[4:] ==", numbers3[4:])

		numbers4 := make([]int, 0, 5)
		printSlice(numbers4)

		/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
		number5 := numbers4[:2]
		printSlice(number5)

		/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
		number6 := numbers4[2:5]
		printSlice(number6)

		var numbers7 []int
		printSlice(numbers7)

		/* 允许追加空切片 */
		numbers7 = append(numbers7, 0)
		printSlice(numbers7)

		/* 向切片添加一个元素 */
		numbers7 = append(numbers7, 1)
		printSlice(numbers7)

		/* 同时添加多个元素 */
		numbers7 = append(numbers7, 2, 3, 4)
		printSlice(numbers7)

		/* 创建切片 numbers1 是之前切片的两倍容量*/
		numbers8 := make([]int, len(numbers7), (cap(numbers7))*2)

		/* 拷贝 numbers 的内容到 numbers1 */
		copy(numbers8, numbers7)
		printSlice(numbers8)

	}

}
