package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	// 標準入力からなんか受取る
	scanner:=bufio.NewScanner(os.Stdin)
	// 入力が成功する間はTrue
	for scanner.Scan(){
		// Text()は指定した入力先の入力を文字列として受け通る
		fmt.Println("スキャン: ", scanner.Text())
	}

	// 入力の読み取りに失敗したとき
	if err:=scanner.Err(); err!=nil{
		fmt.Fprintln(os.Stderr, "読み込みエラー: ", err)
	}
}