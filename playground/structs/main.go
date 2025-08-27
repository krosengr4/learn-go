package main

import "fmt"

type BaseballPlayer struct {
	name     string
	age      int
	position string
	team     string
	homeruns int
	nickname string
}

func main() {
	player1 := BaseballPlayer{
		name:     "Aaron Judge",
		age:      33,
		position: "RF",
		team:     "New York Yankees",
		homeruns: 355,
		nickname: "Aarson Judge",
	}
	player2 := BaseballPlayer{
		name:     "Pete Crow-Armstrong",
		age:      23,
		position: "CF",
		team:     "Chicago Cubs",
		homeruns: 38,
		nickname: "PCA",
	}
	player3 := BaseballPlayer{
		name:     "Cal Raleigh",
		age:      28,
		position: "C",
		team:     "Seattle Mariners",
		homeruns: 143,
		nickname: "Big Dumper",
	}

	fmt.Println("Which player do you want to know more about?")
	fmt.Println("1 -", player1.name, "\n2 -", player2.name, "\n3 -", player3.name)
	var userChoice int
	fmt.Scanln(&userChoice)

	var player BaseballPlayer

	switch userChoice {
	case 1:
		player = player1
	case 2:
		player = player2
	case 3:
		player = player3
	}

	fmt.Println("Name:", player.name)
	fmt.Println("Age:", player.age)
	fmt.Println("Position:", player.position)
	fmt.Println("Team:", player.team)
	fmt.Println("Homeruns:", player.homeruns)
	fmt.Println("Nickname:", player.nickname)

}
