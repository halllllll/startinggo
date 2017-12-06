package main

import (
	"fmt"
)

func main() {
	//未定義のポインタはnil値になる
	var p *int
	fmt.Println(p == nil)

	//アドレス演算子を使えば任意の型からそのポインタ型を生成できる
	var i int
	ptr := &i
	i = 5
	fmt.Println(*ptr) //ポインタ型変数ptrの参照先
	*ptr *= 2         //そのまんま直接操作もできる
	fmt.Println(i)
	fmt.Printf("%T\n", ptr) //intのポインタ型
	ptrptr := &ptr
	fmt.Printf("%T\n", ptrptr) //intのポインタのポインタ型

	ii := 1
	pp := &ii
	some(pp)
	fmt.Println(ii, *pp)

	//配列をいじる
	arry := [3]int{1, 2, 3}
	ary_bai(&arry)
	fmt.Println(arry)
	ary_bai(&arry)
	fmt.Println(arry)
	ary_bai(&arry)
	fmt.Println(arry)

	//Goではいちいちデリファレンスであることを明示的に示さなくてもいい
	strings := [3]string{"smoke", "weed", "everyday"}
	stp := &strings
	fmt.Println((*stp)[0], stp[0], strings[0]) //ほんとは二番目のやつはイレギュラーだがGoではコンパイラが自動的にデリファレンスからの配列の要素の参照だと判断してくれる
	for i, v := range stp {
		fmt.Println(i, v)
	}
}

// intのポインタ型を受け取って直接操作する
func some(i *int) {
	*i *= 999
}

// [3]intの配列のポインタ型を受け取って操作する
func ary_bai(ap *[3]int) {
	for i := 0; i < 3; i++ {
		//(*ap)[i] *= (*ap)[i] Cだとこう書くしかない（らしい
		//Goだとそのまま書いても自動的にポインタ型のデリファレンスから配列型の要素の参照に展開してくれる
		ap[i] *= ap[i]
	}
}
