package main
import (
	"fmt"
)

func main(){
	var c complex128 = 1.0002983+1887298i
	// 複素数（とか指数）は書式指定子上では%gか%e
	fmt.Printf("%g type is:%T\n", c, c)
	// complexで作ると型はcomplex128になる
	c2 := complex(100.4, 20)
	fmt.Printf("%g type is:%T", c2, c2)
}