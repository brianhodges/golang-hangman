package main

import (
	"bufio"
	"fmt"
	"golang-hangman/pkg/game"
	"golang-hangman/pkg/util"
	"os"
	"strings"
)

func main() {
	var attempts []string
	solution := strings.ToLower("WINNER")
	var tries int = 0
	game.PrintHangMan(tries)
	game.PrintGameBoard(solution, attempts)

	for tries < 6 {
		fmt.Println()
		fmt.Print("Guess: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		guess := strings.TrimSpace(strings.ToLower(string([]byte(input)[0])))
		if util.StringInSlice(guess, attempts) {
			fmt.Println("Already tried: ", strings.ToUpper(guess))
			continue
		}
		attempts = append(attempts, guess)

		if !strings.Contains(solution, guess) {
			tries++
		}

		game.PrintHangMan(tries)
		game.PrintGameBoard(solution, attempts)
		if !strings.Contains(game.Board, "_") {
			tries = 6
			continue
		}
	}
	fmt.Println("\nGame Over")
	os.Exit(1)
}
