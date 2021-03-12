package main
import "fmt"

func main() {
	for a:=0; a<3; a++ {
		for b:=3; b>0; b-- {
			fmt.Println("First Loop -",a ," ","Second Loop - " ,b)
		}
	}
}