package main

import (
	"fmt"
)

type Point struct {
	// あとから知ったけどフィールドは頭文字大文字がいいらしい
	x, y, z int
}

type Friends struct {
	Name   string
	Appear Appear //フィールド変数名 構造体名 （わかりづらいけど仕様 マジで頭文字小文字でいいんじゃないか？？？）
}

type Appear struct {
	Serif string
	Ep    int
}

//異なる構造体間で共通の性質を別の構造体として規定することで両方の構造体からアクセスする
type Data struct {
	Genre    []string
	Director string
}

type Anime struct {
	Title string
	Ep    int
	Data
}

type Movie struct {
	Title string
	Min   int
	Data
}

func main() {
	var pt Point
	// 構造体は値型なので各フィールドは初期値をとる
	fmt.Println(pt.x, pt.y, pt.z)
	pt.x = 3
	pt.y = 4
	pt.z = 5
	fmt.Println(pt.x, pt.y, pt.z)
	//複合リテラル 2種類
	pt1 := Point{10, 3, 5}   //定義された順に入れる
	pt2 := Point{z: 8, x: 4} //明示的に指定（初期化してないやつは初期値になる）
	fmt.Println(pt1, pt2)

	//構造体を含む構造体でも複合リテラルが使える
	sarval := Friends{
		Name: "サーバル",
		Appear: Appear{
			Serif: "食べないよ!!",
			Ep:    1, //カンマ必須
		}, //カンマ必須
	}
	fmt.Println(sarval)
	//内側の構造体にアクセス
	fmt.Println(sarval.Appear.Serif)

	//異なる構造体の共通部分を別の構造体として切り出して処理
	kemofre := Anime{
		Title: "けものフレンズ",
		Ep:    12,
		Data: Data{
			Genre:    []string{"百合", "ハートフル", "SF"},
			Director: "たつき",
		},
	}
	leon := Movie{
		Title: "LEON",
		Min:   110,
		Data: Data{
			Genre:    []string{"ロリ"},
			Director: "Luc Besson",
		},
	}

	fmt.Println(kemofre.Genre)
	fmt.Println(leon.Genre)
}
