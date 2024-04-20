package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	hp int
	xp int
}

func (p *Player) Attack(enemyHP int) {
	damage := rand.Intn(enemyHP + 1)
	p.hp -= damage
}

func (p *Player) Heal(amount int) {
	heal := rand.Intn(amount)
	p.hp += heal
}

func main() {
	player := Player{
		hp: 100,
		xp: 0,
	}

	fmt.Println("Welcome to the dungeon of Supremium!")
	fmt.Println("Press Enter to begin!")
	fmt.Scanln()
	var see int
	var option string
	for {
		see = rand.Intn(14)
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
					player.xp += 5
					fmt.Printf("You now have %d hp\n", player.hp)
					fmt.Printf("You now have %d xp\n", player.xp)
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
					player.Heal(5)
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
			} else if see == 3 {
				fmt.Println("You see a snake!\nOptions:")
				fmt.Println("(1) Attack")
				fmt.Println("(2) Dodge")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.Attack(10)
					player.xp += 10
					fmt.Printf("You now have %d hp\n", player.hp)
					fmt.Printf("You now have %d xp\n", player.xp)
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
			} else if see == 4 {
				fmt.Println("You found a super potion!\nOptions:")
				fmt.Println("(1) Drink")
				fmt.Println("(2) Pass")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.Heal(10)
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
			} else if see == 5 {
				fmt.Println("You see an angry cat!\nOptions:")
				fmt.Println("(1) Attack")
				fmt.Println("(2) Dodge")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.Attack(15)
					player.xp += 15
					fmt.Printf("You now have %d hp\n", player.hp)
					fmt.Printf("You now have %d xp\n", player.xp)
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
			} else if see == 6 {
				fmt.Println("You found a shiny potion!\nOptions:")
				fmt.Println("(1) Drink")
				fmt.Println("(2) Pass")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.hp += 5
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
			} else if see == 7 {
				fmt.Println("You found a shiny super potion!\nOptions:")
				fmt.Println("(1) Drink")
				fmt.Println("(2) Pass")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.hp += 10
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
			} else if see == 8 {
				fmt.Println("You see an old man!")
				player.xp += 30
				fmt.Printf("You now have %d xp\n", player.xp)
				fmt.Println("Press enter to continue")
				fmt.Scanln()
			} else if see == 9 {
				fmt.Println("You see a super old man!")
				player.xp += 50
				fmt.Printf("You now have %d xp\n", player.xp)
				fmt.Println("Press enter to continue")
				fmt.Scanln()
			} else if see == 10 {
				fmt.Println("You see a super spider!\nOptions:")
				fmt.Println("(1) Attack")
				fmt.Println("(2) Dodge")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.Attack(30)
					player.xp += 20
					fmt.Printf("You now have %d hp\n", player.hp)
					fmt.Printf("You now have %d xp\n", player.xp)
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
			} else if see == 11 {
				fmt.Println("You see a super snake!\nOptions:")
				fmt.Println("(1) Attack")
				fmt.Println("(2) Dodge")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.Attack(50)
					player.xp += 35
					fmt.Printf("You now have %d hp\n", player.hp)
					fmt.Printf("You now have %d xp\n", player.xp)
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
			} else if see == 12 {
				fmt.Println("You see a super angry cat!\nOptions:")
				fmt.Println("(1) Attack")
				fmt.Println("(2) Dodge")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.Attack(75)
					player.xp += 75
					fmt.Printf("You now have %d hp\n", player.hp)
					fmt.Printf("You now have %d xp\n", player.xp)
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
			} else if see == 13 {
				fmt.Println("You found an ultra potion!\nOptions:")
				fmt.Println("(1) Drink")
				fmt.Println("(2) Pass")
				fmt.Print("Option >> ")
				fmt.Scanln(&option)
				if option == "1" {
					player.Heal(30)
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
