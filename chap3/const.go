package main
import (
	"fmt"
	"math"
)

const (
	A = 10
	B	//自動的に10
	C = "I am 定数"
	D	//自動的に I am 定数
	N = 10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000	//定数は整数型ではないので巨大な数でも代入可能。ただし具体的な変数に代入できるわけではない。fmt.Println(N)とかは出来ない。
)

func main(){
	// Eはmathパッケージでめちゃくちゃ精度高く定義されているが具体的に出力しようとすると浮動小数点の限界float64まで精度が落ちる
	fmt.Println(math.E)
	fmt.Printf("%T", math.E)
}