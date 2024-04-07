package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	hp int
}

func (p *Player) Attack(enemyHP int) {
	damage := rand.Intn(enemyHP + 1)
	p.hp -= damage
}

func main() {
	player := Player{
		hp: 100,
	}

	fmt.Println("Welcome to the dungeon of Supremium!")
	fmt.Println("Press Enter to begin!")
	fmt.Scanln()
	var see int
	var option string
	for {
		see = rand.Intn(2)
		if player.hp <= 0 {
			fmt.Println("You have lost the game, press Enter to exit")
			fmt.Scanln()
			break
		} else {
			if see == 0 {
				fmt.Println("You see nothing, press Enter to continue")
				fmt.Scanln()
			} else if see == 1 {
				fmt.Println("You see a spider!\nOptions:")
				fmt.Println("(1) Attack")
				fmt.Println("(2) Dodge")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.Attack(5)
					fmt.Printf("You now have %d hp\n", player.hp)
					fmt.Println("Press enter to continue")
					fmt.Scanln()
				} else if option == "2" {
					fmt.Println("You decided to dodge, press Enter to continue")
					fmt.Scanln()
				} else {
					fmt.Println("Sorry, but your option was invalid. (No hp was deducted)")
					fmt.Println("Press enter to continue")
					fmt.Scanln()
				}
			} else if see == 2 {
				fmt.Println("You found a healing potion!\nOptions:")
				fmt.Println("(1) Drink")
				fmt.Println("(2) Pass")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.Attack(-5)
					fmt.Printf("You now have %d hp\n", player.hp)
					fmt.Println("Press enter to continue")
					fmt.Scanln()
				} else if option == "2" {
					fmt.Println("You decided to pass, press Enter to continue")
					fmt.Scanln()
				} else {
					fmt.Println("Sorry, but your option was invalid. (No hp was deducted)")
					fmt.Println("Press enter to continue")
					fmt.Scanln()
				}
			}
		}
	}
}
