package main
import (
	"fmt"
)
/*
メインのforがおわった時点で終わる？っぽい？？？？
なのでsubと無名関数が10回出力されてなくても終わる？
*/
func sub1(){
	for i:=0; i<10; i++{
		fmt.Println("I am sub")
	}
}

func main(){
	//関数呼び出しのゴルーチン
	go sub1()
	// 無名関数のゴルーチン
	go func(){
		for i:=0; i<10; i++{
			fmt.Println("I am goroutin")
		}
	}()
	for i:=0; i<10; i++{
		fmt.Println("MAIN")
	}
}