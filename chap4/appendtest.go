package main
import (
	"fmt"
)

func main(){
	// 複数要素をappendすることはできる
	testslicea := []int{0, 1, 2}
	testsliceb := make([]int, 5, 6)
	testslicec := []int{10, 11, 12}
	testslicea = append(testslicea, testsliceb...)
	// スライス同士のappendは複数同時にはできない
	testsliceb = append(testslicea, testsliceb..., testslicec...)
	fmt.Println(testslicea)
}