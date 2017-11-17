package main
import (
	"fmt"
	"./animals" //こんな感じで同一ディレクトリにあるanimalsフォルダ内部のgoファイルを使えるようになるぽい
)

func main(){
	// 呼び出し先で関数は大文字でなければならない（クソドハマリした）
	// app.goから関数を呼び出す
	fmt.Println(AppName())
	fmt.Println(animals.WaniFeed())
	fmt.Println(animals.HipoFeed())
	fmt.Println(animals.GoriFeed())
}
