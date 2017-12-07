package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

//Point型用のメソッドを実装してみる
func (p *Point) Render() {
	fmt.Printf("x:%d, y:%d\n", p.X, p.Y)
}

//Point間の距離を求めるよくあるやつ
func (p *Point) Dist(ap *Point) float64 {
	x, y := p.X-ap.X, p.Y-ap.Y
	return math.Sqrt(float64(x*x + y*y))
}

//typeでエイリアスを作れば基本型をベースにした型でもメソッドを追加できる
// この場合レシーバーに*は不要
func (is Integars) RuisekiSum() []int {
	//ふつうにlenもrangeも使えるんですね
	res := make([]int, len(is))
	res[0] = is[0]
	sum := res[0]
	for i := 1; i < len(is); i++ {
		sum += is[i]
		res[i] = sum
	}
	return res
}

type Integars []int

//型のコンストラクタで初期化処理してみる
type Yome struct {
	ID   int
	Name string
	Sex  int
	Tags []string
}

func NewYome(id int, name string) *Yome {
	//最初に自分自身のなんつーかインスタンス的なやつを作る
	y := new(Yome)
	y.ID = id
	y.Name = name
	y.Sex = 0
	y.Tags = []string{"ここに素晴らしい属性が入ります"}
	//で最後に返す
	return y
}

func main() {
	p := &Point{X: 3, Y: 10}
	p.Render()
	p2 := &Point{X: 2, Y: 4}
	dist := p.Dist(p2)
	fmt.Println(dist)

	lis := Integars{-4, 2, 20, 5, -8, 3, 5, 2, 2, 4}
	fmt.Println(lis.RuisekiSum())
	myyome := *NewYome(0, "谷川柑菜")
	fmt.Println(myyome)
}
