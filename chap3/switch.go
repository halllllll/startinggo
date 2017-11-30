package main
import (
	"fmt"
)

func main(){
	/*
	switchによる分岐で定数と式(bool)の混在はエラーを招く
	*/
	switch x:=1; x{
		case 1, 2, 3:
			fmt.Println("xは1,2,3のどれか")
		case x%2==0:
			//そもそもフォールスルーが働かないのでスルーされる
			fmt.Println("xは偶数")
		default:
			fmt.Println("あばばばばばばばばば")
	}
}