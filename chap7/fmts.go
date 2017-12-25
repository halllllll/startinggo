package main
import (
	"fmt"
	"os"
	"log"
)

func main(){
	// ふつうに標準出力に出す
	n:=4
	fmt.Printf("%d*%d=%d\n", n, n, n*n)
	// 文字列として生成 ついでに小数点以下どこまでを取るかってやつ
	s:=fmt.Sprintf("[%.4f]\n", 1.283722)
	fmt.Printf(s)
	// ファイルへ出力
	// なぜかどこにもファイルが作られていない（実行はされる）ので謎　もう無視しよう
	f, err:=os.Create("ababa.txt")
	if err!=nil{
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "|%05d|%05d|\n", 121, 33)
}