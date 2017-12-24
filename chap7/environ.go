package main
import (
	"fmt"
	"os"
)

func main(){
	// 環境変数を取得する
	for _, v := range os.Environ(){
		fmt.Println(v)
	}
}