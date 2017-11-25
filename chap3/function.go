package main
import (
	"fmt"
)


func main(){
	fmt.Println(plus(10, 20))
	void()
}

//足すだけのやつ
func plus(x, y int) int {
	return x+y
}

//戻り値なしでもreturn必要
//そんなことはなかった
func void(){
	fmt.Println("void")
}