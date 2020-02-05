package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1~100까지 난수 생성, 같은 난수 나오면 기존 난수 모두 더해 리턴
func main() {
	var wait sync.WaitGroup
	mutex := sync.RWMutex{}

	defer fmt.Println("main go routine ends!")

	chanRand := make(chan int)
	chanRandNums := make(chan []int)
	var chanDone chan bool = make(chan bool)

	// create
	wait.Add(1)
	go func() {
		defer func() {
			mutex.Lock() // write protection since wait is pointer
			wait.Done()
			mutex.Unlock()
			fmt.Println("createFunc done")
		}()

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		var val int

		for {
			val = r.Intn(100)
			fmt.Println("[createFunc] rand:", val, "(sent to checkFunc)")
			chanRand <- val

			if <-chanDone {
				fmt.Println("[createfunc] chanDone : true")
				break
			}
		}
	}()

	// check
	wait.Add(1)
	go func() {
		defer func() {
			mutex.Lock()
			wait.Done()
			mutex.Unlock()
			fmt.Println("checkFunc done")
		}()

		var val int
		randNums := make([]int, 0)
		var quit bool = false

		for {
			val = <-chanRand
			fmt.Println("[checkFunc] recv:", val)
			for i, _ := range randNums {
				if val == randNums[i] {
					fmt.Printf("[checkFunc] %d exist!!\n", val)
					chanRandNums <- randNums
					close(chanRandNums)
					quit = true
					break
				}
			}

			if quit {
				fmt.Println("[checkFunc] quit!!")
				chanDone <- true
				break
			} else {
				randNums = append(randNums, val)
				fmt.Println("[checkFunc] buf:", randNums)
				chanDone <- false
			}
		}
	}()

	// print
	wait.Add(1)
	go func() {
		defer func() {
			mutex.Lock()
			wait.Done()
			mutex.Unlock()
			fmt.Println("printFunc done")
		}()

		var sum int
		randNums := make([]int, 0)

		randNums = <-chanRandNums
		for i, _ := range randNums {
			sum += randNums[i]
		}

		fmt.Println("[printFunc] sum : ", sum)
	}()

	wait.Wait()
}
