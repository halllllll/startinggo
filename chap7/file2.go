package main
import (
	"fmt"
	"os"
	"log"
)

func main(){
	f, err:=os.Open("/Volumes/ExtremeDrive/GitHub/book/startinggo/chap7/memo.md")
	if err!=nil{
		log.Fatal(err)
	}
	defer f.Close()

	// []byte型のスライスにファイルの内容を読み込ませる
	bs:=make([]byte,1028)
	n, err2:=f.Read(bs)
	if err2!=nil{
		log.Fatal(err2)
	}
	// nは単なる読み込んだバイト数
	fmt.Println(n)
	// 元データはbsの中にある
	fmt.Println(string(bs[:]))
	// 途中から読み込むときはReadAtで何バイト目から読み込むか指定する
	// とかいって下のように書くとなんかしらんけどちゃんと1028byteぶん読み込まれてて意味不明
	nn, err3 := f.ReadAt(bs, 990)
	if err3!=nil{
		log.Fatal(err3)
	}
	fmt.Println(nn)

	// ファイルのステータスを取得
	// os.FileInfo型とエラーが渡される
	fi, err:=f.Stat()
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.Mode())
	fmt.Println(fi.ModTime())
	fmt.Println(fi.IsDir())
}