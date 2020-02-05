package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup

	defer fmt.Println("main go routine ends!")

	myChannel_1 := make(chan int)
	myChannel_2 := make(chan string)

	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("Func1")
		time.Sleep(time.Second)

		myChannel_1 <- 1
		close(myChannel_1)
	}()

	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("Func2")

		val, flag := <-myChannel_2
		if !flag {
			fmt.Println("[Func2] channel closed")
		} else {
			fmt.Println("[Func2] ", val, flag)
		}
	}()

	val, flag := <-myChannel_1
	if !flag {
		fmt.Println("[main] channel closed")
	} else {
		fmt.Println("[main] ", val, flag)
	}

	myChannel_2 <- "hi!!"
	close(myChannel_2)

	wait.Wait()
}
