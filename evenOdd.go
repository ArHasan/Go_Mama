package main
import "fmt"

func main(){
	var num int
	fmt.Print("Enter the any Number: ")
	fmt.Scanf("%d", &num)
	 odd:=0 
	 even:=0
	for i:=1; i<=num; i++{
		if i%2==1 {
			fmt.Println(i,"= Odd")
			odd++
		} else if i%2==0 {
			fmt.Println(i,"= Even")
			even++
		}
	}
	fmt.Println("Total Odd number is = ",odd)
	fmt.Println("Total Even number is = ",even)
}