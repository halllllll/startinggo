package main
import (
	"fmt"
)

//string型の引数を受け取って戻り値のない関数を返す関数
func retfunc(a string) func(){
	return func(){
		fmt.Println(a)
	}
}

//string型の引数を受け取って「string型を返す関数」を返す関数
func retfunc2(a string) func()string{
	return func()string{
		return a+"!!!!!!!!!!!"
	}
}

func main(){
	f:=retfunc("おっぱい")
	f()
	f2:=retfunc2("いっぱい")
	fmt.Println(f2)	//retfunc2で返される関数のメモリアドレス
	fmt.Println(f2()) //retfunc2で返される関数を呼び出す
	fmt.Println(retfunc2("あばばばばっば")())
}