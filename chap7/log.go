package main
import (
	"log"
	"os"
)

func main(){
	// ログには日付と時刻が自動的に入る
	log.Print("1:\n")
	log.Println("2:")
	log.Printf("%d:\n", 3)
	// ログの出力先を変更できる
	// 標準出力に出してみる
	log.SetOutput(os.Stdout)
	log.Println("標準出力と標準エラー出力の違いがわからん")
	// ファイルを作ってそこに書き込んでみる
	// やっぱりなんかしらんけどファイルが作成されないけど問題なく実行はされる、意味不明につき無視
	f, err:=os.Create("log.txt")
	if err!=nil{
		log.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("こんちゃー輝夜月だよぉおおおお")

	log.SetOutput(os.Stdout)
	// ログのフォーマットを変えてみる
	log.SetFlags(log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	log.Println("おぱい")
	// プレフィックスを設定してみる
	log.SetPrefix("[おいィィィィィィィィィ]")
	log.Println("ぺろぺろ")
}