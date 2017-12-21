package main

import (
	"fmt"
)

// 組込みのerrorインターフェースを使ってエラーを吐かせてみる

/*エラーごとに構造体を作る？？？？？らしい？？？？？？？これがスタンダードなのかは不明*/
type MyError struct {
	ErrorMessage string
	ErrorCode    int
}

/*errorインターフェースのメソッドをMyError構造体に実装。errorインターフェースはError()stringのみをもつ*/
func (e *MyError) Error() string {
	return e.ErrorMessage
}

func RaiseError() error {
	//わざとエラーを起こす
	//MyError構造体のインスタンスっつーかポインタを返す
	return &MyError{
		ErrorMessage: "raise error test :)",
		ErrorCode:    274823,
	}
}

func main() {
	err := RaiseError()
	// よくある形
	if err != nil {
		fmt.Println(err.Error()) //ErrorCodeが吐かれない理由は知らん
	} else {
		fmt.Println("アルコール")
	}

	//型アサーションを使うと本来の構造体をそのまま取り出せる
	e, ok := err.(*MyError)
	if ok {
		// eは構造体
		fmt.Println(e) //ErrorCodeが吐かれない理由は知らん
		fmt.Println(e.ErrorMessage)
		fmt.Println(e.ErrorCode)
	}
}
