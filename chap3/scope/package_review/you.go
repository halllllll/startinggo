package main

import "strings"

// mainパッケージならどこからでも使えるやつを定義
// これをパッケージ変数という
// := で定義できない
var Thisispackageparam_you = "うんこ"

// 関数も同じパッケージからなら使える
func Oioioi() string{
	return strings.Repeat("こちらはyou側となります", Thisispackageparam_me)
}