package main
import (
	"fmt"
)

func receiver(uke <-chan int){
	// 受信専用でint型を格納するチャネルを引数にとる
	// 当然送受信可能のチャネル型でもいい
	for {
		// iにukeが受信した値を代入
		//i := <-uke
		//fmt.Println(i)
		//そのまま受信してもいい
		fmt.Println(<-uke)
	}
}

func receiver2(i int, ch chan int){
	ch <- i*i
}


func main(){
	// int型のチャネルchを定義（送受信可能）
	var ch chan int
	// 受信専用のチャネルukeを定義
	var uke <-chan int
	// 送信専用のチャネルsemeを定義
	var seme chan<- int
	fmt.Printf("ch=%#v, uke=%#v, seme=%#v\n", ch, uke, seme)

	//送信なしに受信したものを出力しようとするとデッドロックを引き起こす
	//fmt.Println(<-uke) 受信待ちの状態になるもののほかのゴルーチンが存在しないので永遠に待ち続けることになるのでエラー
	//ゴルーチン間で試してみる用のチャネル
	// なんでかしらんけどchだとできないんだが?????makeで作らなきゃいかんのかな・・・
	//あと送受信専用のやつって使い方が全然わからん
	chtest := make(chan int)
	chtest2:= make(chan int)
	//　ゴルーチンに登録
	go receiver(chtest)
	chtest <- 999
	for i:=0; i<100; i++{
		//チャネルにiを送信
		chtest <-i
		//チャネルにiを送信
		go receiver2(i, chtest2)
		// <-chtest2　の部分で受信している
		fmt.Printf("chtest2:%d\n", <-chtest2)
	}

	//バッファサイズを越えるデータをチャネルに入れるとデッドロックを起こす
	deadlockch := make(chan rune, 3)
	deadlockch <- 'A'
	deadlockch <- 'B'
	deadlockch <- 'C'
	// deadlockch <- 'D'これすると死
	//空のバッファを受信するとデッドロック
	killmebaby := make(chan int)
	fmt.Println(<-killmebaby)
}
