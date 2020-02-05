package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main go routine ends!")

	myChannel := make(chan int)
	test := 1

	go func() {
		time.Sleep(time.Second)
		myChannel <- 1 // insert
		test = 2
	}()

	temp := <-myChannel // fetch, wait until values are inserted in queue
	fmt.Println(temp)
	fmt.Println("test:", test)
}
