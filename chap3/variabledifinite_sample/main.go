package main
import (
	"fmt"
)
// go runで,mainパッケをインポートしているすべてのファイルを指定してビルドすると、nがsub.goでも使えることがわかる
// → nがmainパッケージのパッケージ変数になってしまっている
var n int=100
func main(){
	fmt.Println(n)
	fmt.Println(Output())
}