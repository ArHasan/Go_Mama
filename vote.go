package main

import "fmt"

func main() {
	fmt.Println("Enter your age : ")
	var age float64
	fmt.Scanf("%f",&age)
	if(age>=18){
		fmt.Printf("You are Eligible to Vote")
	}else{
		fmt.Printf("You are not permitted to vote ")
	}


}