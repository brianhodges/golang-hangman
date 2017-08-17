package util

import (
	"log"
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
