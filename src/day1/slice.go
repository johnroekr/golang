package main

import "fmt"

func main() {
	var temp string
	temp = "hello"
	fmt.Println(temp)

	var tvt [5]byte
	tt := []byte(temp)

	for index, val := range tt {
		fmt.Println(index, val)
		tvt[index] = val
	}

	// pointer
	fmt.Println(tt)
	checkBuffer(tt)
	fmt.Println(tt)

	// value
	fmt.Println(tvt)
	checkBufferCopy(tvt)
	fmt.Println(tvt)

	ttt := [...]byte{tvt[0], tvt[1], tvt[2], tvt[3], tvt[4]}
	fmt.Println("ttt : ", ttt)
	checkBufferCopy(ttt)
	fmt.Println("ttt : ", ttt)
}

// call by reference
func checkBuffer(temp []byte) { // pointer
	temp[0] = 1
}

// call by value
func checkBufferCopy(temp [5]byte) { // value
	temp[0] = 1
}
