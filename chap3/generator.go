package main
import (
	"fmt"
)

//呼び出すたび倍にしていくようなジェネレータ（的なやつ）を実装する
func bai() func()uint64{
	var i uint64 = 1
	return func()uint64{
		tmp:=i
		i*=2
		return tmp
	}
}

func main(){
	b :=bai()
	for i := 0; i < 20; i++ {
		fmt.Println(b())
	}
}