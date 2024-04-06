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
    fmt.Println("Welcome to the dungeon of Supremium!")
    fmt.Println("Press Enter to begin!")
    fmt.Scanln()
    var see int

    while true {
        see = rand.Intn(2)
        if player.hp <= 0 {
            break
        } else {
            if see == 0 {
                fmt.Println("You see nothing, press Enter to continue")
                fmt.Scanln()
            } else {
                // TODO: insert code
            }
        }

    }
}
