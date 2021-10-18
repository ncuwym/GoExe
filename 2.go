package main

import "fmt"

func main() {
	demo_switch()
}

func demo_switch() {

	// var n = 3
	// n := 3

	switch n := 3; n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("0")
	}
}

func testSwitch() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}
