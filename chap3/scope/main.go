package main
import (
	"fmt"
	//同一ディレクトリにあるfileフォルダのgoファイルを扱う
	"./file"
)

func main(){
	fmt.Println(foo.MAX)	//使える
	fmt.Println(foo.FooFunc(10))	//これも使える
	//fmt.Println(foo.internal_consta)　無理
}