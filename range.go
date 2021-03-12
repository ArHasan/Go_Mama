package main
import "fmt"

func main() {
	food := map[string]string{"1":"mango","2":"apple","3":"banana"}  
	for k, v := range food {  
	   fmt.Printf("%s -> %s\n", k, v)  
	}      
}