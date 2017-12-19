package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*
タグとは構造帯のフィールドに付与するメタ情報のこと
フィールドの最後に文字列またはrow文字列で指定
（文字列のほうが好まれる傾向がある）
*/

type User struct {
	id    int    "ID"
	Name  string "名前"
	email string "メール"
}

// タグ「json:"キー"」をrow文字をつかって囲んでjsonパッケージからパースできるようにする

//フィールドの頭文字が小文字だとタグをjsonで出力するときに反映されない
type User2 struct {
	id    int    `elem:"user_ID"`
	email string `elem:"user_email"`
	Name  string `elem:"user_name"`
}

//フィールドの頭文字が大文字だと反映される
type User3 struct {
	Id    int    `elem:"user_ID"`
	Email string `elem:"user_email"`
	Name  string `elem:"user_name"`
}

func main() {
	u := User{id: 1, Name: "toris", email: "ハイボール"}
	t := reflect.TypeOf(u)
	fmt.Println(t)
	//構造帯のすべてのフィールドを回す
	//reflect.TypeOf型のNumFieldメソッドはその構造体のもつフィールドの数を返すらしい
	for i := 0; i < t.NumField(); i++ {
		// reflect.TypeOf型のFieldメソッドは指定したintにおけるインデックスでフィールドの情報を返すらしい
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}
	/*
		encoding/jsonパッケージにはタグを拾ってjsonのキーにする機能がある
	*/

	/*
		User2のフィールドの頭文字に小文字が含まれているとうまくjsonで出力されない
	*/
	u2 := User2{
		email: "yoyoyo",
		id:    1298,
		Name:  "明太バター",
	}
	bs2, _ := json.Marshal(u2)
	fmt.Println(string(bs2))

	//大文字にしたバージョンはちゃんとフィールドが出力される
	u3 := User3{
		Email: "thisisemai",
		Id:    2827,
		Name:  "myname",
	}
	bs3, _ := json.Marshal(u3)
	fmt.Println(string(bs3))
}
