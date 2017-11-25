package main
import (
	"fmt"
)
func main(){
	var x interface{}
	// てっきり方はinterface{}になると思ったらnilだた
	fmt.Printf("%#v type is %T", x, x)
}