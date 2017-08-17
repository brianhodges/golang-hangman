package main

import (
	"bufio"
	"fmt"
	"golang-hangman/pkg/game"
	"golang-hangman/pkg/util"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var wordList []string
	var attempts []string
	var tries int

	//Get random word for game solution from word list file
	rand.Seed(time.Now().Unix())
	file, err := os.Open("./words.txt")
	util.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	solution := strings.ToLower(wordList[rand.Intn(len(wordList))])
	game.PrintHangMan(tries)
	game.PrintGameBoard(solution, attempts)

	for tries < 6 {
		fmt.Print("\nGuess: ")
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
