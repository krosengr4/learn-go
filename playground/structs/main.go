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
		position: "Right Field",
		team:     "New York Yankees",
		homeruns: 355,
		nickname: "Aarson Judge",
	}
	player2 := BaseballPlayer{
		name:     "Pete Crow-Armstrong",
		age:      23,
		position: "Center Field",
		team:     "Chicago Cubs",
		homeruns: 38,
		nickname: "PCA",
	}
	player3 := BaseballPlayer{
		name:     "Cal Raleigh",
		age:      28,
		position: "Catcher",
		team:     "Seattle Mariners",
		homeruns: 143,
		nickname: "Big Dumper",
	}

	fmt.Println("Which player do you want to know more about?")
	fmt.Println("1 -", player1.name, "\n2 -", player2.name, "\n3 -", player3.name)
	var userChoice int
	fmt.Scanln(&userChoice)

	switch userChoice {
	case 1:
		fmt.Println(player1.String())
	case 2:
		fmt.Println(player2.String())
	case 3:
		fmt.Println(player3.String())
	}
}

func (bp BaseballPlayer) String() string {
	return fmt.Sprintf("%s plays %s for The %s. In his MLB career, %s has hit %d homers. He is currently %d years old and has been nicknamed %s.",
		bp.name, bp.position, bp.team, bp.name, bp.homeruns, bp.age, bp.nickname)
}
