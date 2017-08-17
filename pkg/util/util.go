package util

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

//CheckError logs error
func CheckError(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}

//StringInSlice detects if array contains string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//ClearConsole clears terminal window
func ClearConsole() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
