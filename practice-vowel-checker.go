package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Program start successfully")
	checkVowel()
	fmt.Println("Program run successfully")
}

func checkVowel() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	for _, v := range strings.ToLower(name) {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			fmt.Println("Your name contain one or more vowel letter")
			return
		}
	}
	fmt.Println("Your name does not contain vowel letters")
}
