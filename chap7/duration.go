package main
import (
	"fmt"
	"time"
	"reflect"
)

func main(){
	// time.HourとかMinuteとかはtime.Duration型
	fmt.Println(reflect.TypeOf(time.Hour))
	// ほかにもあるけどとりあえずこんなもん
	// それぞれの最小値を表す
	fmt.Println(time.Hour)
	fmt.Println(time.Second)
	fmt.Println(time.Nanosecond)
	//文字列から生成
	d, _ := time.ParseDuration("2h36m113s")
	// ちゃんとこのように60秒以上で書くと分に繰り上げしてくれる
	fmt.Println(d)

	// time.Time型と組み合わせる。Addを使ってみる
	t:= time.Now()
	fmt.Println(t)
	// いまから2時間15分22秒後
	t = t.Add(2*time.Hour+15*time.Minute+22*time.Second)
	fmt.Println(t)
	// 差分をとるときはSub
	tokyo2020:=time.Date(2020, 7, 24, 0, 0, 0, 0, time.Local)
	dif := tokyo2020.Sub(t)
	fmt.Println(dif)

	// 数値指定で増減
	// 1年後
	t1:=t.AddDate(1, 0, 0)
	fmt.Println(t1)
	// 20年と4ヶ月と40日前
	t2:=t.AddDate(-20, -4, -40)
	fmt.Println(t2)

}