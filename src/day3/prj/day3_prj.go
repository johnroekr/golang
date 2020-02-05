package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DealerInterface interface {
	addPlayer()
	createCard()
	shuffleCard()
	distributeCard()
	startGame()
	checkGame()
}

type Dealer struct {
	name string
		
}

type Player struct {
	name string
}

func remove(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}

func createCard(num int) []int {
	arr := make([]int, num*10)

	idx := 0
	for i := 0; i < num; i++ {
		for j := 0; j < 10; j++ {
			arr[idx] = j + 1
			idx++
		}
	}

	return arr
}

func shuffleCard(arr []int, ining int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	// fmt.Println(arr)

	return arr
}

func main() {

	card := createCard(2)
	fmt.Println("created :", card)

	card = shuffleCard(card, 5)
	fmt.Println("shuffled :", card)

	player := make([][]int, 2, 2)
	cnt := 0

	fmt.Println("======================================")
	fmt.Printf("ining\tr_idx\trandom\tcard\n")

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(cnt + 1)
			r_num := r.Intn(len(card) - 1)
			player[i][j] = card[r_num]
			card = remove(card, r_num)

			// fmt.Printf("%d\t%d\t%d\t%d\n", len(card), r_num, player[i][j], card)
			cnt++
		}
	}

	/*



		p1_out := (player[0][0] + player[0][1]) % 10
		p2_out := (player[1][0] + player[1][1]) % 10

		fmt.Println("======================================")
		fmt.Printf("player 1 : ( [%2d] + [%2d] ) %% 10 ---> %d\n", player[0][0], player[0][1], p1_out)
		fmt.Printf("player 2 : ( [%2d] + [%2d] ) %% 10 ---> %d\n", player[1][0], player[1][1], p2_out)
		fmt.Println("======================================")

		if p1_out > p2_out {
			fmt.Println("player 1 wins!")
		} else if p1_out < p2_out {
			fmt.Println("player 2 wins!")
		} else {
			fmt.Println("draw!")
		}*/
}
