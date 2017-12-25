package main
import (
	"fmt"
	"flag"
)

func main(){
	// オプションの各値を格納する変数
	var (
		max int
		msg string
		x bool
	)

	// コマンドラインオプションの定義
	// 変数のポインタ、コマンド、初期値、ヘルプしたときのドキュメント？
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	flag.StringVar(&msg, "m", "", "処理メッセージだお（ ＾ω＾）")
	flag.BoolVar(&x, "x", false, "拡張オプションってなんだよ")
	// コマンドをパース
	flag.Parse()
	fmt.Println("処理数の最大値::: ", max)
	fmt.Println("処理メッセージ::: ", msg)
	fmt.Println("拡張オプション::: ", x)
}