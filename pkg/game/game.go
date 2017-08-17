package game

import (
	"fmt"
	"golang-hangman/pkg/util"
	"strings"
)

var Board string

//PrintGameBoard shows current guessed correct letters/spaces
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

//PrintHangMan shows the current status of the game (hangman)
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
