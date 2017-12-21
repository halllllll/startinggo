package main

import (
	"fmt"
)

// 文字列化できることを示すインターフェース
type Stringable interface {
	ToString() string
}

//構造体Person
type Person struct {
	Name string
	Age  int
}

//構造体PersonのメソッドToString
func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

// 構造体Car
type Car struct {
	Number string
	Model  string
}

// 構造体CarのメソッドToString
func (c *Car) ToString() string {
	return fmt.Sprintf("%s(%s)", c.Number, c.Model)
}

//Stringableインターフェースさえ実装していればどんな型からも呼び出せるようにする汎用性高いPritln()をつくる
func Println(s Stringable) {
	fmt.Println(s.ToString())
}

//ついでに適当な構造体をもひとつ追加
type Kuso struct {
	Type   string
	Color  string
	Weight int
}

func (k *Kuso) ToString() string {
	return fmt.Sprintf("%s:%s色(%dkg)", k.Type, k.Color, k.Weight)
}

func main() {
	/* 異なる型を共通のStringableインターフェースにまとめる */
	vs := []Stringable{
		&Person{
			Name: "誰か",
			Age:  222,
		},
		&Car{
			Number: "oppai-CCCCCC00120",
			Model:  "C3PO",
		},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	fmt.Println("-----------------------")
	//汎用性高いほうをつかってみる
	nanika := []Stringable{
		&Kuso{
			Type:   "巻き",
			Color:  "Brown",
			Weight: 1000,
		},
		&Person{
			Name: "誰か",
			Age:  222,
		},
		&Car{
			Number: "oppai-CCCCCC00120",
			Model:  "C3PO",
		},
	}
	for _, v := range nanika {
		Println(v)
	}
}
