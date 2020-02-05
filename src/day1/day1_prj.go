package main

import (
	"fmt"
	"math/rand"
	"time"
)

func remove(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	card := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10}
	// fmt.Println(card)

	p1 := make([]int, 2)
	p2 := make([]int, 2)
	// fmt.Println(p1)
	// fmt.Println(p2)

	rand_num := make([]int, 4)
	idx := 20

	fmt.Println("======================================")
	fmt.Printf("ining\tr_idx\trandom\tcard\n")

	for i := 0; i < 4; i++ {
		r1_num := r1.Intn(idx)
		// fmt.Println("%i", r1_num)

		rand_num[i] = card[r1_num]
		// fmt.Println(idx, rand_num)

		// fmt.Println("before : ", card)
		card = remove(card, r1_num)
		// fmt.Println("after  : ", card)
		fmt.Printf("%d\t%d\t%d\t%d\n", i+1, r1_num, rand_num[i], card)

		idx--
	}

	p1[0] = rand_num[0]
	p1[1] = rand_num[1]
	p2[0] = rand_num[2]
	p2[1] = rand_num[3]

	p1_out := (p1[0] + p1[1]) % 10
	p2_out := (p2[0] + p2[1]) % 10

	fmt.Println("======================================")
	fmt.Printf("player 1 : ( [%2d] + [%2d] ) %% 10 ---> %d\n", p1[0], p1[1], p1_out)
	fmt.Printf("player 2 : ( [%2d] + [%2d] ) %% 10 ---> %d\n", p2[0], p2[1], p2_out)
	fmt.Println("======================================")

	var p1_pair bool = false
	var p2_pair bool = false

	if p1[0] == p1[1] {
		p1_pair = true
	}
	if p2[0] == p2[1] {
		p2_pair = true
	}

	if p1_pair && p2_pair {
		if p1[0] > p2[0] {
			fmt.Println("pair! player 1 wins!")
		} else {
			fmt.Println("pair! player 2 wins!")
		}
	} else if p1_pair {
		fmt.Println("pair! player 1 wins!")
	} else if p2_pair {
		fmt.Println("pair! player 2 wins!")
	} else {
		if p1_out > p2_out {
			fmt.Println("player 1 wins!")
		} else if p1_out < p2_out {
			fmt.Println("player 2 wins!")
		} else {
			fmt.Println("draw!")
		}
	}
}
