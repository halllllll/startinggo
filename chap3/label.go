package main
import (
	"fmt"
)

func main(){
	// 深いネストから一気にぬけられる
	A:=[6]int{1, 2, 5, 8, 11, 13}
	B:=[8]int{2, 3, 5, 10, 14, 15, 16, 17}
	C:=[4]int{5, 6, 8, 21}
	var ans int
	MYLABEL:
		for _, a:=range A{
			for _, b:=range B{
				for _, c:=range C{
					if (a+b+c)%7==0{
					fmt.Printf("最初に7で割り切れる組み合わせは%d+%d+%d\n", a, b, c)
					ans = a+b+c
					break MYLABEL
					}
				}
			}
		}
	fmt.Println(ans)

	// continueで処理をスルーする
	/*
	例がクソすぎてクソ　参考にするな
	MYLABEL2:
		for _, a:=range A{
			for _, b:=range B{
				for _, c:=range C{
					if a>b || b>c{
						continue MYLABEL2
					}
					fmt.Printf("大小関係がちゃんとできてる組み合わせ %d<%d<%d\n", a, b, c)
				}
			}
		}
	}*/
}