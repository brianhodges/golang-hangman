package game

import (
	"fmt"
	"golang-hangman/pkg/util"
	"strings"
)

//Board is used to keep track of current game state - hangman letters
var Board string

//PrintGameBoard prints current game state
func PrintGameBoard(solution string, attempts []string) {
	Board = ""
	for _, letter := range solution {
		if util.StringInSlice(string(letter), attempts) {
			Board = Board + " " + strings.ToUpper(string(letter)) + " "
		} else {
			Board = Board + " _ "
		}
	}
	fmt.Println(Board)
}

//PrintHangMan prints the current hangman state
func PrintHangMan(tries int) {
	fmt.Println()
	switch tries {
	case 0:
		fmt.Println("	    +-----+")
		fmt.Println("	    |     |")
		fmt.Println("	          |")
		fmt.Println("	          |")
		fmt.Println("	          |")
		fmt.Println("	          |")
		fmt.Println("	         ---")
	case 1:
		fmt.Println("	    +-----+")
		fmt.Println("	    |     |")
		fmt.Println("	    O     |")
		fmt.Println("	          |")
		fmt.Println("	          |")
		fmt.Println(" 	          |")
		fmt.Println(" 	        ---")
	case 2:
		fmt.Println("	    +-----+")
		fmt.Println("	    |     |")
		fmt.Println("	    O     |")
		fmt.Println("	    |     |")
		fmt.Println("	          |")
		fmt.Println("	          |")
		fmt.Println("	         ---")
	case 3:
		fmt.Println("	    +-----+")
		fmt.Println("	    |     |")
		fmt.Println("	    O     |")
		fmt.Println("	    |     |")
		fmt.Println("	   /      |")
		fmt.Println("	          |")
		fmt.Println("	         ---")
	case 4:
		fmt.Println("	    +-----+")
		fmt.Println("	    |     |")
		fmt.Println("	    O     |")
		fmt.Println("	    |     |")
		fmt.Println("	   / \\    |")
		fmt.Println("	          |")
		fmt.Println("	         ---")
	case 5:
		fmt.Println("	    +-----+")
		fmt.Println("	    |     |")
		fmt.Println("	    O     |")
		fmt.Println("	    |--   |")
		fmt.Println("	   / \\    |")
		fmt.Println("	          |")
		fmt.Println("	         ---")
	case 6:
		fmt.Println("	    +-----+")
		fmt.Println("	    |     |")
		fmt.Println("	    O     |")
		fmt.Println("	  --|--   |")
		fmt.Println("	   / \\    |")
		fmt.Println("	          |")
		fmt.Println("	         ---")
	}
	fmt.Println()
}
