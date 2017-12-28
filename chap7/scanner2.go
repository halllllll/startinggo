package main
import (
	"fmt"
	"bufio"
	"strings"
)

func main(){
	// bufio.Scannerは標準で行単位で読み込むが、これを変更するにはbuifo.Scanner型に用意されているSplit関数を使う

	// 改行するのでrow文字で試す
	s:=`ABC DEF G
	HIJK LM NOP Q
	RS
	TU V
	W X Y       Z!! !   !!`
	// strings.NewReader(文字列)はstrings.Reader型を返す
	// strings.Readerは文字列を読みこむためにio.Readerとかを実装しているらしい
	r := strings.NewReader(s)
	// bufio.NewScannerの引数はio.Reader型である必要がある(io.Readerを実装していればよい)
	scanner := bufio.NewScanner(r)
	// 行ではなく文字単位（タブとか空文字とか行）で区切るようにする
	// デフォルトではScanLineだがこれをScanWordsにする
	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}