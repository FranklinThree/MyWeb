package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Uint2String(number uint) string {
	return strconv.Itoa(int(number))
}
func Confirm(action string) bool {
	var massage string
	fmt.Print("Confirm to", action, "? (y/n) ")
	for {
		n, err := fmt.Scanf("%s", &massage)
		if !CheckErr(err) {
			return false
		} else if n != 1 {
			fmt.Println("Incorrect input number :", strconv.Itoa(n), "\n expect number: 1")

		} else {
			switch strings.ToLower(massage) {
			case "y":
				return true
			case "n":
				return false
			case "r":
				fmt.Print("Confirm ", action, "? (y/n) ")
			default:
				fmt.Println("Incorrect input value:", massage, "\nPlease input 'y' or 'n'.")
			}

		}
	}
}
