package main

import (
	"fmt"
	"runtime"
	"sync"
)

func ThreadTest(wait *sync.WaitGroup) {
	defer wait.Done()

	var temp int = 0
	for i := 0; i < 5; i++ {
		temp = temp + i
		// fmt.Println(temp, i, "Out Thread")
		fmt.Println("[+]", temp, "th Out Thread wait:", wait)
	}
}

func main() {
	runtime.GOMAXPROCS(8)
	fmt.Println(runtime.GOMAXPROCS(0))

	wait := new(sync.WaitGroup)
	defer fmt.Println("wait:", wait)

	for i := 0; i < 5; i++ {
		wait.Add(1) // test after deleting this line
		go ThreadTest(wait)

		wait.Add(1)
		fmt.Println(wait)

		go func() {
			defer wait.Done()

			var temp int = 0
			for i := 0; i < 500000; i++ {
				temp = temp + i
				// fmt.Println(temp, i, "In Thread")
				fmt.Println("[-]", temp, "th In Thread wait:", wait)
			}
		}()
	}

	fmt.Println("Main Thread")
	wait.Wait()
}
