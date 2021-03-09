package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter any number: ")
	fmt.Scanf("%d", &a)
	if (a%2==0) {
		fmt.Printf("%d is even number", a)
	}else{
		fmt.Printf("%d is odd number",a)
	}
}
