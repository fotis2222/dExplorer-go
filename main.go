package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	hp int
}

func (p Player) Attack(enemyHP int) {
	p.hp -= rand.Intn(enemyHP + 1)
}

func main() {
    player := Player{
		hp: 100,
	}

	fmt.Println(player.hp)
}
