package main
import (
	"fmt"
)

func main(){
	fruits:=[3]string{"mikan", "ringo", "yuzu"}
	for i, v := range fruits{
		fmt.Printf("[%d]=%s\n", i, v)
	}
} 