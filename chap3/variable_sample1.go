package main
import (
	"fmt"
)
func main(){
	var(
		x, y, z, w int
		s, t, u, v string
	)
	// 両辺は異なる型が混じっていても代入できる？ →　型宣言が正しい限りできる
	x, y, s, t = 100, 200, "帰りたい", "眠い"
	fmt.Println(x, y, s, t)
	// python的なアンパック代入的なやつができる？ →　できない
	/*
	ints := []int{10, 11}
	strs := []string{"アメ舐めたい", "ヤル気ない"}
	z, w = ints
	u, v = strs
	fmt.Println(ints, strs)
	*/
	fmt.Println(z, w, u, v)
}