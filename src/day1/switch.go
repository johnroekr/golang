package main

import (
	"fmt"
)

func MyCal(a int, b int) (int, int) {
	return (a + b), (a - b)
}

func main() {
	temp_1, _ := MyCal(10, 10)

	switch {
	case temp_1 > 1:
		fmt.Println(temp_1, "check1")
		fallthrough
	case temp_1 > 5:
		fmt.Println(temp_1, "check2")
	default:
		fmt.Println(temp_1, "default")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i, _ := MyCal(1, 1); i < 10; i++ {
		fmt.Println(i)
	}

	for {
		if temp_1 > 30 {
			break
		}
		fmt.Println(temp_1)
		temp_1 = temp_1 + 1
	}
}
