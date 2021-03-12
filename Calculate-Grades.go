package main

import "fmt"

func main() {
	fmt.Printf("Enter your five subject marks : ")
	var sub1, sub2, sub3, sub4, sub5 int
	// fmt.Scanf("%d %d %d %d %d",&sub1,&sub2,&sub3,&sub4,&sub5)
	fmt.Scan(&sub1,&sub2,&sub3,&sub4,&sub5)
	var avg=(sub1+sub2+sub3+sub4+sub5)/5
	fmt.Println("avg is = ",avg)

}