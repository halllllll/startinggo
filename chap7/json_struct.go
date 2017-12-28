package main
import (
	"fmt"
	"encoding/json"
	"time"
	"log"
)

type Nanika struct{
	Id int
	Name string
	Age time.Time
}

func main(){
	u:=new(Nanika)
	u.Id = 2284839042
	u.Name = "輝夜月だよおおぉぉぉぉぉ"
	u.Age = time.Now()

	bs, err:=json.Marshal(u)
	if err!=nil{
		log.Fatal(err)
	}
	// json.Marshalの戻り値は[]byteなのでそｈのままでは出力しても読めない
	fmt.Println(bs)
	// string()を使って文字列化
	fmt.Println(string(bs))

	// 逆にjsonから構造帯を作る
	src := `{
		"Id": 999999,
		"Name": "けもみみのじゃロリ狐娘youtuberおじさん",
		"Age": "2017-12-05T00:00:00.00+09:00"
	}`
	homunkurusu := new(Nanika)
	err2 := json.Unmarshal([]byte(src), homunkurusu)
	if err2!=nil{
		log.Fatal(err2)
	}
	fmt.Printf("%+v\n", homunkurusu)
}