package main
import (
	"fmt"
)

const (
	A = iota
	B	//自動的にiotaが使われる（のでインクリメントされる）
	C
	D = "ちくわ大明神"
	E = iota	//途中でぶったぎってもiotaは保存される
)

func main(){
	fmt.Println(A, B, C, D, E)
}