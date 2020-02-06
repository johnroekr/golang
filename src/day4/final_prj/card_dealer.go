package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	// wait := sync.WaitGroup{}

	dealer := Dealer{name: "dealer"}

	dealer.waitPlayer()
	go dealer.startGame()
}

type DealerInterface interface {
	waitPlayer()
	connectHandler()
	addPlayer(name string)
	startGame()
	playGame()

	checkGame()
	sendPlayerStatus()
}

type Dealer struct {
	name    string
	card    []int
	deck    int
	players []Player

	round 		int
	draw  		int

	bStart 		bool
}

func (d *Dealer) waitPlayer() {

	fmt.Println("[waitPlayer] waiting..")
	myListen, err := net.Listen("tcp", ":5554")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		myListen.Close()
	}()

	for {
		connect, err := myListen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		d.connectHandler(connect)
	}
}

func (d *Dealer) connectHandler(connect net.Conn) {

	go func() {
		recvBuf := make([]byte, 4096) // receive buffer: 4kB

		for {
			n, err := connect.Read(recvBuf) // wait until recv sth
			if err != nil {
				if io.EOF == err {
					fmt.Println("connection is closed from client; %v", connect.RemoteAddr().String())
					continue
				}
				fmt.Println(err)
				continue
			}
			if 0 < n {
				data := recvBuf[:n]
				// fmt.Println(string(data))
				d.addPlayer(string(data))
			}
		}
	}

	go func() {

		writeBuf := make([]byte, 4096)	

		for {
			connect.Write([]byte(<-d.chanWrite))
			fmt.Println("Send Data : ", "HIHI")
			time.Sleep(time.Second * 1)
		}
	}

	
	for {
		
	}
}

func (d *Dealer) addPlayer(playerName string) {

	p := Player{name: playerName}
	d.players = append(d.players, p)
	fmt.Printf("[addPlayer] player %s joined.(total %d)\n", playerName, len(d.players))

	if len(d.players) >= 2 {
		d.bStart = true
	}
}

func (d *Dealer) deletePlayer(playerName string) {

	// need to check
	for i, p := range d.players {
		if p.name == playerName {
			d.players = append(d.players[:i], d.players[i+1:]...)
			fmt.Printf("[deletePlayer] player %s has left.(total %d)\n", playerName, len(d.players))
		}
	}

	if len(d.players) < 2 {
		fmt.Printf("[deletePlayer] stopped playing game. (total %d)\n", len(d.players))
		d.bStart = false
	}
}

func (d *Dealer) startGame() {

	d.deck = 2
	d.makeCard()

	for {
		if d.bStart {
			d.cardShuffle()
			d.playGame()
			d.sendPlayerStatus()
		}

		fmt.Println("[startGame] A new game will be starting in 10 seconds..")
		time.Sleep(time.Second * 10)
	}
}

func (d *Dealer) makeCard() {

	d.card = make([]int, d.deck*10)
	for i := 0; i < d.deck*10; i++ {
		d.card[i] = (i + 1) % 10
		if d.card[i] == 0 {
			d.card[i] = 10
		}
	}
}

func (d *Dealer) cardShuffle() {

	myCard := make([]int, 20)
	s := rand.NewSource(time.Now().UnixNano())

	for index := range myCard {
		r := rand.New(s)
		randomNumber := r.Intn(len(d.card))
		myCard[index] = d.card[randomNumber]
		d.card = append(d.card[:randomNumber], d.card[randomNumber+1:]...)
	}

	d.card = myCard
}

func (d *Dealer) playGame() {

	for index := range d.players {
		d.players[index].mycard = append(d.players[index].mycard, d.card[(index*d.deck):(index*d.deck)+2])
	}
	d.checkGame()
}

func (d *Dealer) checkGame() {

	playerResult := make([]int, 0)

	for index := range d.players {
		playerResult = append(playerResult, (int(d.players[index].mycard[d.round][0]+d.players[index].mycard[d.round][1]))%10)
	}

	max := -1
	draw_check := 0
	max_index := 0

	for i := range playerResult {
		temp := playerResult[i]

		if temp > max {
			max = temp
			max_index = i
		} else if max == temp {
			draw_check++
		} else {
			continue
		}
	}

	if draw_check == 0 {
		d.players[max_index].winHit++
	} else {
		d.draw++
	}
	d.round++
	// fmt.Println(playerResult)
}

func (d *Dealer) sendPlayerStatus() {
	for i := range d.players {
		fmt.Println("[sendPlayerStatus] Player ", i, " WinHit : ", d.players[i].winHit)
	}
	fmt.Println("[sendPlayerStatus] Round : ", d.round)
	fmt.Println("[sendPlayerStatus] Draw Round: ", d.draw)
}

type Player struct {
	name   string
	mycard [][]int
	winHit float32
	chanWrite	chan
}
