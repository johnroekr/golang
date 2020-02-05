package main

import (
	"fmt"
)

func MyCal(a int, b int) (int, int) {
	return (a + b), (a - b)
}

func main() {
	temp_1, _ := MyCal(10, 10)

	if temp_1 == 1 {
		fmt.Println(temp_1, "Check1")
	} else if temp_1 == 10 {
		fmt.Println(temp_1, "Check2")
	} else if temp_1 == 20 {
		fmt.Println(temp_1, "Check3")
	} else {
		fmt.Println("No Match!!")
	}

}
