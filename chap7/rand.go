package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// シード値を固定しているのでコードを実行するたび同じ乱数が生成される
	rand.Seed(122)
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	// シード値を現在時刻にしているのでコードを実行するたび異なる乱数が生成される
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	// 生成する範囲を指定する（最大値しか指定できないっぽいので負値をいれたいときは結果から減算するなりする）
	fmt.Println(rand.Intn(100))

	// デフォルトの疑似乱数発生器はGoのランタイム上のやつを共有しているのでどっかでシード値を書き換えるとすべての場所で使ってる擬似乱数もその影響を受ける
	// このため独自に擬似乱数発生機を生成できる仕組みがある
	// NewSource(ソースとなる値)で新しいソースを生成
	//src := rand.NewSource(time.Now().UnixNano())
	src:=rand.NewSource(221)
	// Newでソースをもとに擬似乱数発生機を生成
	rnd := rand.New(src)
	fmt.Println(rnd.Intn(10))
	fmt.Println(rnd.Intn(10))
	fmt.Println(rnd.Intn(10))
}