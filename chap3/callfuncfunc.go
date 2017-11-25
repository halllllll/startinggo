package main
import (
	"fmt"
)
//戻り値のない関数を引数にとる戻り値のない関数
func callfunc(f func()){
	f()
}

func main(){
	callfunc(func(){fmt.Println("わかりづらい")})
}
