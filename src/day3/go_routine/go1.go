package main

import (
	"fmt"
	"sync"
)

func main() {
	defer fmt.Println("Main end")
	var wait sync.WaitGroup

	for i := 0; i < 3; i++ {
		wait.Add(1) // one thread is added (cnt++)
		temp := "Hi"

		go func(n int) {
			defer wait.Done() // thread is done (cnt--)
			fmt.Printf("%s, %dth thread\n", temp, n)
		}(i)

		fmt.Println("Main thread")
	}
	wait.Wait()
}
