package main
import (
	"fmt"
	"strings"
)

func main(){
	ss := []string{"welcome", "to", "ようこそ", "ジャパリパーク"}
	// Joinは結合する
	// 結合する文字は文字列なので何文字でもいい
	sss := strings.Join(ss, "(ｱｰﾊｰﾝ)")
	fmt.Println(sss)
	// Indexは指定した文字列が何文字目から始まるのかインデックスを教えてくれる 複数あれば最初のインデックス、存在しなければ-1
	s1 := "また夜が明ければお別れ夢は遠き幻にあなたを追いかけていた光の中で抱かれるたびあたたかい風をたより"
	fmt.Println(strings.Index(s1, "あ"))
	fmt.Println(strings.Index(s1, "あた"))
	fmt.Println(strings.Index(s1, "輝夜月だよぉおおおお"))
	// LastIndexは最後のインデックス
	fmt.Println(strings.LastIndex(s1, "あ"))
	// IndexAnyは第二引数の文字列に含まれる文字のうち"いずれかが"含まれるときの最初のインデックス" とかいうんか珍しいやつ
	fmt.Println(strings.IndexAny(s1, "ろんぴ悪もへ"))
	// 置換する 第四引数は置換する最大数
	s2 := "ええいああ君からもらい泣き"
	s3 := strings.Replace(s2, "もらい泣き", s2, 10)
	fmt.Println(s3)
	s4 := "すもももももももものうち"
	s5 := strings.Replace(s4, "も", "し", 3)
	fmt.Println(s5)
	// Fieldsはスペースとか区切りの文字列を[]string型にする
	s6 := "28 4 102 3 56 9"
	s7 := strings.Fields(s6)
	for _, s := range s7{
		fmt.Println(s)
	}
}	