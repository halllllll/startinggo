package main

import (
	"fmt"
	"time"
)

func receiver(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			//チャネルがクローズかつバッファが空なら終了
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "is over.")
}

func main() {
	//チャネル用意
	ch := make(chan int, 20) //バッファサイズ20でint型のデータを保持するチャネル
	//ゴルーチン登録
	go receiver("1st goroutine", ch)
	go receiver("2nd goroutine", ch)
	go receiver("3rd goroutine", ch)

	//なんかこのループ回すときローカル変数定義してやるやり方って推奨されてんだっけ？？？？？
	i := 0
	for i < 100 {
		//チャネルにiを送信
		ch <- i
		i++
	}
	//チャネルをクローズ
	close(ch)

	//待つ
	time.Sleep(3 * time.Second)
}
