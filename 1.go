package main

import (
	"fmt"
	"strings"
)

func main() {
	// demo_if()
	demo_for_range()
	chengfa()

}

func demo_if() {
	if age := 17; age > 18 {
		fmt.Println("hello")
	} else if age < 16 {
		fmt.Println("show")
	} else {
		fmt.Println("gun")
	}
}

func demo_str() {
	fmt.Println("Hello, World!")
	str := "好_聪_明"
	s1 := strings.Split(str, "_")
	fmt.Println(s1)
}

func demo_for() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//死循环
	// for {
	// 	fmt.Println("hhh")
	// }
}

func demo_for_range() {
	s := "hello 黑马"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}

func chengfa() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

}
