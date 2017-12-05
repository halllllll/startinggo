package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	//わざと受信を失敗させる
	ch1 <- 100
	ch2 <- 200
	select {
	case <-ch1:
		fmt.Println("ch1あります")
	case c2 := <-ch2:
		fmt.Println("c2あります ", c2)
	case ch3 <- 3:
		fmt.Println("ch3へ送信しますた ", <-ch3)
	default:
		//すべてのcase節が実行不可能の場合のみ実行される
		fmt.Println("yoyoyo")
	}
}
