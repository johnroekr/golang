

type Player struct {
	name      string
	mycard    [][]int
	winHit    float32
	dropRound int
}

func (p *Player) receiveCard(receive_card []int) {
	p.mycard = append(p.mycard, receive_card)
}
