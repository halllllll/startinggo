package main
import (
	"fmt"
)

func main(){
	// float64型は正の整数を0で割るとINFに、負の整数だと-INFになる
	// 0.0を.0で割るとNanになる
	zero:= .0
	one := 1.0
	fmt.Println(one/zero)
	fmt.Println(-one/zero)
	fmt.Println(zero/zero)
	// 先頭に0をつけると8進数として解釈される
	fmt.Println(0755)
	// でも浮動小数点ついてるとそんなことにはならない
	fmt.Println(0734.2)
	// 整数型に変換するとfloorになる
	f:=29.99931
	f2:=-20.02
	fmt.Println(int(f), int(f2))
	f*=-1
	fmt.Println(int(f))
}