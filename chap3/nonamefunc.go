package main
import (
	"fmt"
)

func main(){
	// 無名関数をつくる
	f:=func(x, y int) int {return x+y}
	fmt.Println(f(20, 4))
	// 書式指定子で調べてみると、int2つを引数にとってintひとつを返す関数型、と出る
	fmt.Printf("%T\n", f)
	// 更に書式指定子でデフォルトフォーマットをみてみるとメモリアドレスが出力される
	fmt.Printf("%v\n", f)
	// ちなみにGoリテラルをみてみると型とメモリアドレスが出力される
	fmt.Printf("%#v\n", f)
	fmt.Printf("%v\n", f(19, 4))
}