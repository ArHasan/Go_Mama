package main

import "fmt"

func main(){
	fmt.Println("Enter character : ")
	var input string
	fmt.Scanln(&input)

	switch (input) {
	case "a":
		fmt.Print("Vowel")
	case "e":
		fmt.Print("Vowel")
	case "i":
		fmt.Print("Vowel")
	case "o":
		fmt.Print("Vowel")
	case "u":
		fmt.Print("Vowel")
	default:
		fmt.Print("Consonant")
	}
}