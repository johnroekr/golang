package main

import (
	"fmt"
	"sync"
	"time"
)
// 1~100까지 난수 생성, 같은 난수 나오면 기존 난수 모두 더해 리턴

func receiveRand(wait *sync.WaitGroup, c <-chan int) sum int {
	defer wait.Done()

	randNum := make([]int)
	var val, sum int

	for {
		val <- c
		for i range randNum {
			if randNum[i] == val {
				return
			} else {
				sum += <- c
			}
		}
	}
}

func createRand(wait *sync.WaitGroup, c chan<- int) {
	defer wait.Done()

	for {
		
	}
}


func main() {
	var wait sync.WaitGroup

	defer fmt.Println("main go routine ends!")

	myChannel := make(chan int)

	wait.Add(1)
	go createRand(myChannel)

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
