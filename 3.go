package main

import "fmt"

func main() {
	// arr()
	// readArr()
	// arrs()
	demo_exe()
}

func arr() {
	//var 数组变量名  [数组大小]元素类型
	//数组的长度必须是常量，长度是数组类型的一部分，一旦定义，长度不能变

	var a [3]int //如果只声明不使用，编译会有问题
	a = [3]int{0, 1, 2}
	fmt.Println(a)
	//声明并初始化
	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	//自动推断数组长度
	c := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(c)
	//根据索引赋值
	d := [5]int{0: 0, 4: 4}
	fmt.Println(d)
}

//遍历
func readArr() {
	d := [5]int{0: 0, 4: 4}
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}

	for i, v := range d {
		fmt.Println(i, v)

	}
}

//多维数组

func arrs() {
	var a [3][2]int
	a = [3][2]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}
	fmt.Println(a)
	//遍历二维数组
	for _, i := range a {
		// fmt.Println("i")
		for _, j := range i {
			fmt.Println(j)
		}
	}
}

//数组是值类型
//数组是引用类型

//练习题
func demo_exe() {
	a := [...]int{1, 3, 5, 7, 8}
	res := 0
	for _, v := range a {
		res += v
		// fmt.Println(res)
	}
	fmt.Println(res)

	result := 8
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if i != j && result == a[i]+a[j] {
				fmt.Println(i, j)
			}
		}
	}
}
