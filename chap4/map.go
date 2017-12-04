package main
import (
	"fmt"
)

func main(){
	//マップをmakeを使って作る
	m:=make(map[int]string)
	m[0] = "おっぱい"
	m[10] = "おっぱい"
	m[299] = "おっぱいがいっぱい"
	fmt.Println(m)
	//マップのリテラル
	m2:=map[int]string{1: "富士", 2: "鷹", 3: "ダーーーーーーッ!!!!"}
	//みやすさのため次のような書き方でもいい
	m3:=map[int]string{
		10: "眠い",
		21: "だるい",
		32: "死にたい",	//, 必須
	}
	fmt.Println(m2)
	fmt.Println(m3)
	//存在しない値にもアクセスできてしまう
	v1 := m3[1000]
	fmt.Println(v1=="")
	//隠し戻り値みたいなやつ（正式名称しらん）を利用する
	v2, ok := m3[10]
	fmt.Println(v2, ok)
	v3, ok := m3[999]
	fmt.Println(v3, ok)
}