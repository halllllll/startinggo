package main
import (
	"fmt"
	"reflect"
)

func main(){
	// 要素数が異なればまったく違う型として認識される
	ar1 := [3]string{"カバさん", "アリクイさん", "あんこう"}
	ar2 := [5]string{"うさぎさん", "カメさん", "レオポンさん", "カモさん", "アヒルさん"}
	ar1type, ar2type := reflect.TypeOf(ar1), reflect.TypeOf(ar2)
	if ar1type != ar2type{
		fmt.Println(ar1type, ar2type)
	}
}