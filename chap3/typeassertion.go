package main
import (
	"fmt"
	"reflect"
)

func anything(a interface{}){
	fmt.Println(a)
	//fmt.Println(a.(int64))
}

func main(){
	anything([...]string{"oyoyo", "wiwiiw", "lalalal"})
	// 型アサーション
	var x interface{} = 10
	i:=x
	fmt.Println(i.(int))	// > 10
	// ２つの変数に入れると二つ目の戻り値でエラーを吐かずにboolで返る
	j, isString := x.(string)
	fmt.Println(j, isString)//> (空文字) false
	// switch節のみ、以下のが使える
	switch x.(type){
		case string:
			fmt.Println("xは文字列型")
		case int, uint:
			fmt.Println("xはint型")
		default:
			fmt.Println("なんだお前")
	}
	var xx interface{}="おちんちん"
	switch v:=xx.(type){
		case bool:
			fmt.Printf("%#dはbool型\n", v)
		case int:
			fmt.Printf("%#dはint型\n", v)
		case string:
			fmt.Printf("%#dはstring型\n", v)
		default:
			fmt.Println("なんだお前\n")
	}
	// case節で複数の定数を指定した場合の挙動
	/*
	var y interface{} = [...]int{10, 3, 5}
	fmt.Println(reflect.TypeOf(y))
	switch y.(type){
		case int, uint, string, [3]int:
			fmt.Prinltln(y[0])
	}
	*/
	// 同上
	var z interface{} = "おっぱい"
	switch z.(type){
			case int, uint, string:
				fmt.Println(z+z)
	}
}