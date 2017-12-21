package main

import (
	"fmt"
)

//Stringerを実装しない適当な構造体
type Nanika struct {
	Name string
	Id   int
}

//適当な構造体に組込みのfmt.Stringerインタフェースを実装 stringを返すString()があればよい
type Tekitou struct {
	Name string
	Id   int
}

func (t *Tekitou) String() string {
	return fmt.Sprintf("<<ID:%d, Name:%s>>", t.Id, t.Name)
}

func main() {
	//Stringerを実装してないやつをそのままPrintlnする
	n := Nanika{
		Name: "はい",
		Id:   917273,
	}
	fmt.Println(n)
	//実装したやつを普通にPrintlnする
	t := &Tekitou{
		Name: "名前",
		Id:   1984723892,
	}
	fmt.Println(t)
}
