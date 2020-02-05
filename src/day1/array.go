package main

import "fmt"

func main() {
	buff := [5]int{1, 2, 3, 4, 5}
	fmt.Println(buff)

	buff_1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(buff_1)

	buff_2 := [5][5]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}
	fmt.Println(buff_2)

	buff2 := make([]int, 5)
	buff2[0] = 1
	buff2[1] = 2
	fmt.Println(buff2)

	buff2_1 := make([][]int, 5, 5)
	buff2_1[0] = buff2
	buff2_1[1] = []int{3, 4, 5, 6, 7}
	fmt.Println(buff2_1)

	for i, val := range buff2_1 {
		// fmt.Println(val) // index
		fmt.Println(i, val)
	}
}
