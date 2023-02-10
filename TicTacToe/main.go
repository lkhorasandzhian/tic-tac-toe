package main

import "TicTacToe/GameLibrary"

func main() {
	GameLibrary.PrintGameTitle()

	var players = GameLibrary.InitializeNames()

	GameLibrary.GameProcess(players)
}
