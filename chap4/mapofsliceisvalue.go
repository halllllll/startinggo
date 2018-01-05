package main
import (
	"fmt"
)

func main(){
	// mapでvalueにスライスを使う
	m1 := map[int] []int{
		1: []int{0, 1, 2},
		2: []int{0, 1},
		3: []int{3, 4, 5},
	}
	// スライスを使う場合は一部省略できる
	m2 := map[int] []int{
		1: {0, 1, 2},
		2: {0, 1},
		3: {3, 4, 5},
	}
	fmt.Println(m1)
	fmt.Println(m2)
}