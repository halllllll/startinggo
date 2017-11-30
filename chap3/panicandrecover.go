package main
import (
	"fmt"
)

func testRecover(src interface{}){
	// deferなのでpanicでも実行される
	defer func(){
		if x:=recover(); x!=nil{
			switch v:=x.(type){
				case int:
					fmt.Printf("panic: int=%v\n", v)
				case string:
					fmt.Printf("panic: stirng=%v\n", v)
				default:
					fmt.Println("あわわわあわわわわあわああわわわわわっわああわわわあ")
			}
		}
	}()//deferに登録するのは基本的に関数呼び出しの形
	//ここでいわれのないパニックを起こす
	panic(src)
	return
}	

func main(){
	testRecover(1222)
	testRecover("yo")
	testRecover([0]int{})
}