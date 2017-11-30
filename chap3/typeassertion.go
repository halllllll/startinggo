package main
import (
	"fmt"
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
			fmt.Printf("%#dはbool型", v)
		case int:
			fmt.Printf("%#dはint型", v)
		case string:
			fmt.Printf("%#dはstring型", v)
		default:
			fmt.Println("なんだお前")
	}
}