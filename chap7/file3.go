package main
import (
	"fmt"
	"os"
)

func main(){
	f,_:=os.Create("/Volumes/ExtremeDrive/GitHub/book/startinggo/chap7/foo.txt")
	//いま作ったファイルのステータスを見てみる
	fi, _ := f.Stat()
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.Mode())
	fmt.Println(fi.ModTime())
	fmt.Println(fi.IsDir())

	// []byte型で書き込む
	f.Write([]byte("おはよおぉぉぉ輝夜月だよおおぉぉぉぉぉ!!!!!!\n"))
	// 末尾にシークを移動
	// シークがなんなのかはしらんけど
	f.Seek(0, os.SEEK_END)
	// でも、てっきりこれでシーク位置を末尾で初期化できると思ってたんだけどなんか違うみたい。教科書はこれで終わっているので役に立たない

	// これでシーク位置が取れるらしい
	spos, _ := f.Seek(0, 1)
	fmt.Println(spos)
	// シーク位置から（上のコードでシークを取得したんでそれを設定して）始める
	f.WriteAt([]byte("あばばばばばばｂ"), spos) 
	// WriteStringを使えばバイトを介さずそのまま文字列を入れれる
	//f.WriteString("あばばばばばばば")
	// 何バイトから書き込むか指定してみる
	// f.WriteAt([]byte("YOOOO"), 120)

}