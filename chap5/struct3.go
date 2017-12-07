package main

import (
	"fmt"
	"math"
)

//構造体・参照型・スライス・エイリアス・メソッドなどを組み合わせた複雑なデータをやってみる

// Point構造体のポインタ型のスライスのエイリアスをPointsと命名
type Points []*Point

//int型３つのフィールドをもつ構造体
type Point struct {
	X, Y, Z int
}

//Pointの型のコンストラクタ
func NewPoint() *Point {
	p := new(Point)
	p.X, p.Y, p.Z = 1, 1, 1
	return p
}

//Pointのメソッド
//他のPoint型との距離を出す
func (myP *Point) Distance(otherP *Point) float64 {
	x, y, z := myP.X-otherP.X, myP.Y-otherP.Y, myP.Z-otherP.Z
	return math.Sqrt(float64((x * x) + (y * y) + (z * z)))
}

//[]*Pointのメソッド
//全列挙
func (ps Points) PrintPoints() {
	for i, v := range ps {
		fmt.Printf("index:%d [%v]\n", i, *v)
	}
}

func main() {
	fmt.Println("はい")
	ps := Points{}
	for i := 0; i < 5; i++ {
		ps = append(ps, NewPoint())
	}
	ps.PrintPoints()
	//適当になんか値ぶちこむ（ランダムはまだしりません><）
	for i := 0; i < len(ps); i++ {
		ps[i].X += i * i
		ps[i].Y += i*i + ps[i].X
		ps[i].Z += i + i - ps[i].X
	}
	ps.PrintPoints()
	for _, v := range ps[1:] {
		dis := ps[0].Distance(v)
		fmt.Println(dis)
	}
}
