package main
import (
	"fmt"
	"os"
	"log"
)

func main(){
	f, err:=os.Open("なんか適当なファイル")
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println(f)
	}
	// ファイルを閉じる処理をdeferに登録しておけば確実にストリームを切れる
	defer f.Close()
}