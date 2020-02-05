package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main go routine ends!")

	myChannel := make(chan int)

	go func() {
		time.Sleep(time.Second)
		myChannel <- 1 // insert
		close(myChannel)
	}()

	go func() {
		val, flag := <-myChannel
		fmt.Println("1)", val, flag)
		val, flag = <-myChannel
		fmt.Println("2)", val, flag)
	}()

	val, flag := <-myChannel
	fmt.Println("3)", val, flag)
}
