package main
import (
	"fmt"
	"net/url"
	"reflect"
)

func main(){
	// これポインタでやらないと結果が違ってくるからな
	u:=&url.URL{}
	u.Scheme = "https"
	u.Host = "soundcloud.com"
	u.Path = "djjasrac"
	// Query()の戻り値はurl.Values型で、これはmap[string][string]のエイリアスであるらしい
	q := u.Query()
	// なんかmapがSet()もつとか調べてもどこにも出てこp無かったんだけどまあもういいや、Setできるってことだろ
	q.Set("key1", "yo")
	q.Set("key2", "muri")
	// url.Encode()でいい具合に=と&でつなげてくれるってことだろ
	fmt.Println(q.Encode())
	u.RawQuery = q.Encode()
	fmt.Println(u)
	fmt.Println(q)
	fmt.Println(reflect.TypeOf(q))
}