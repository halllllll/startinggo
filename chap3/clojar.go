package main
import (
	"fmt"
)

//「string型を引数にとってstring型を返す関数」を返す関数
func later() func(string)string{
	var store string
	// 「string型を引数にとってstring型を返す関数」を返す
	return func(next string) string{
		s:=store
		store = next
		return s
	}
}

func main(){
	f :=later()
	fmt.Println(f)
	fmt.Println(f("Java"))
	fmt.Println(f("Python"))
	fmt.Println(f("C++"))
	fmt.Println(f("おっぱい"))
}

