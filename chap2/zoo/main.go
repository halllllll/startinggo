package main
import (
	"fmt"
	"./animals" //こんな感じで同一ディレクトリにあるanimalsフォルダ内部のgoファイルを使えるようになるぽい
)

func main(){
	// 呼び出し先で関数は大文字でなければならない（クソドハマリした）
	fmt.Println(animals.WaniFeed())
	fmt.Println(animals.HipoFeed())
	fmt.Println(animals.GoriFeed())
}
