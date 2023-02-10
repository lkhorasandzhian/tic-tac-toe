package GameLibrary

import (
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	name   string
	number int
	score  int
}

func (player Player) String() string {
	return "Name: " + player.name + "\n" +
		"Score: " + strconv.Itoa(player.score) + "\n"
}

func repeatGame() bool {
	var repeat string
	for {
		fmt.Printf("Do you wanna play another round? (y/n): ")
		_, err := fmt.Fscanln(os.Stdin, &repeat)
		if err == nil && (repeat == "y" || repeat == "n") {
			fmt.Println()
			return repeat == "y"
		}
		fmt.Println("Incorrect answer. Please, try again...")
		fmt.Println()
	}
}

func GetWinner(field [3][3]int) int {
	// Checking straight lines.
	for i := 0; i < 3; i++ {
		if field[i][0] == field[i][1] && field[i][1] == field[i][2] {
			return field[i][0]
		}
		if field[0][i] == field[1][i] && field[1][i] == field[2][i] {
			return field[0][i]
		}
	}

	// Checking diagonals.
	if (field[0][0] == field[1][1] && field[1][1] == field[2][2]) ||
		(field[0][2] == field[1][1] && field[1][1] == field[2][0]) {
		return field[1][1]
	}

	// There is no winner in game.
	return 0
}

func PrintField(field [3][3]int) {
	fmt.Println()

	fmt.Println("  1  2  3")
	fmt.Println("  ———————")
	for i, row := range field {
		fmt.Printf("%d |", i+1)
		for _, element := range row {
			switch element {
			case 1:
				fmt.Print("X|")
			case 2:
				fmt.Print("O|")
			default:
				fmt.Print(" |")
			}
		}
		fmt.Println()
		fmt.Println("  ———————")
	}

	fmt.Println()
}

func MakeMove(players [2]Player, index int) (int, int) {
	var x, y int

	for {
		fmt.Printf("%s's turn (x y): ", players[index%2].name)
		var xStr, yStr string

		_, err := fmt.Fscanln(os.Stdin, &xStr, &yStr)
		if err != nil {
			fmt.Println("Incorrect console input. Please, try again...")
			fmt.Println()
			continue
		}

		x, err = strconv.Atoi(xStr)
		if err != nil || !(x >= 1 && x <= 3) {
			fmt.Println("Incorrect horizontal coordinate. Please, try again...")
			fmt.Println()
			continue
		}

		y, err = strconv.Atoi(yStr)
		if err != nil || !(y >= 1 && y <= 3) {
			fmt.Println("Incorrect vertical coordinate. Please, try again...")
			fmt.Println()
			continue
		}

		break
	}

	return x, y
}

func GameProcess(players [2]Player) {
	var field [3][3]int
	for {
		for i := 0; i < 9; i++ {
			var x, y int = MakeMove(players, i)

			field[y-1][x-1] = i%2 + 1

			PrintField(field)

			var winnerNumber = GetWinner(field)
			if winnerNumber != 0 {
				fmt.Printf("%s has won the game!\n", players[winnerNumber-1].name)
				players[winnerNumber-1].score++
				fmt.Println()
				fmt.Println(players[0])
				fmt.Println(players[1])
				break
			}
		}
		field = [3][3]int{}

		if !repeatGame() {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}

func InitializeNames() [2]Player {
	var players [2]Player
	for i := 0; i < len(players); i++ {
		for {
			fmt.Printf("Write username #%d: ", i+1)
			_, err := fmt.Fscanln(os.Stdin, &players[i].name)
			if err == nil {
				break
			}
			fmt.Println("Incorrect username. Please, try again...")
			fmt.Println()
		}
		fmt.Println()
	}
	return players
}

func PrintGameTitle() {
	fmt.Printf("%s\n", "tic-tac-toe game")
	fmt.Println()
}
