package main
import (
	"fmt"
	"os"
)

func main(){
	// os.Argsの要素数を表示
	// os.Argsは実行時に入力されたコマンドライン引数が入ってるらしい
	fmt.Printf("length: %d\n", len(os.Args))
	// 中身をぜんぶ出力
	for _, v:=range os.Args{
		fmt.Println(v)
	}
}