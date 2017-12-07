package main

import "fmt"

//typeでまとめてエイリアスを定義
type (
	i_slice []int
	rgba    [3]float64
	i_chan  chan int
	s_chan  chan string
	Bai     func(i int)
)

func sum(is []int, bai Bai) []int {
	for _, v := range is {
		bai(v)
	}
	return is
}

func main() {
	pair := i_slice{-1, 3, 2, 5}
	col := rgba{0.334, 0.233, 0.734}
	fmt.Println(pair, col)

}
