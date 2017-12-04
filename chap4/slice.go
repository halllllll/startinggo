package main
import (
	"fmt"
	"reflect"
)

func main(){
	s:=make([]int, 10)
	fmt.Println(s)
	// 配列と見た目は同じに見える
	a:=[10]int{}
	fmt.Println(a)
	fmt.Printf("%v, %v\n", s, a)
	fmt.Printf("%#v, %#v\n", s, a)
	fmt.Println(reflect.TypeOf(s), reflect.TypeOf(a))
	// 代入や参照は配列と同じ文法
	s[0] = 100200120
	s[1] = 392
	fmt.Println(s)
	fmt.Println(s[8])
	// 存在しない要素にアクセスするとパニック
	//fmt.Println(s[10])

	//組み込み関数len, cap
	fmt.Println(len(s))
	fmt.Println(cap(s))
	//要素数5, 容量10のスライス
	s2:=make([]int, 5, 10)
	fmt.Println(len(s2), cap(s2))
	
	//簡易スライス式でスライスを作成
	array:=[5]int{10, 4, 2, -5, 3}
	slice:=array[2:]	//pythonと似た感じ,ただし負値を使って末尾からn番目のアクセスみたいなやつはできない
	fmt.Println(array, slice)

	//ちなみに簡易スライスは文字列にも使えるが、バイト列としてみなすことに注意
	st1:="aiueo"
	st2:="おいうえあ"
	fmt.Println(st1[1:3])
	fmt.Println(st2[3:9])	//utf8では日本語は一文字につき3byte使う

	//appendを使って拡張する
	testslice := make([]int, 3, 5)
	testslice = append(testslice, 10, 3)
	fmt.Println(testslice, cap(testslice))
	testslice = append(testslice, 1, 1, 1)
	fmt.Println(testslice, cap(testslice))	//容量を越えたので倍の容量を確保している

	//スライス同士のappendも可能
	slicea := make([]int, 4)
	sliceb := []int{3, 3, 0}
	slicec := append(slicea, sliceb...)
	fmt.Println(slicec, cap(slicec))

	//スライスをコピーする
	// なぜか「コピーが実行された要素数」が返る
	n := copy(sliceb, slicea)
	// 
	fmt.Println(n, sliceb)

	//完全スライス式
	slicex := make([]int, 5) //要素5容量5のスライス
	slicey := slicex[1:5]
	fmt.Println(cap(slicey))
	slicez := slicex[1:3:5]
	fmt.Println(cap(slicez))

	fmt.Println(sum(5, 67, 3, 2))
	fmt.Println(sum(-4))
	fmt.Println()
	//スライスを可変長引数として使う
	integers:=[]int{4, 2, 3, 6, 0, 10, -20}
	fmt.Println(sum(integers...))
}

//可変長引数をとる関数
func sum(slices ...int)int{
	ans:=0
	for _,v := range slices{
		ans+=v
	}
	return ans
}