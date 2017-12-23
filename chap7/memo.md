## Goのパッケージ
* 使用頻度の高そうな標準パッケージをみていくよ
### OS
* <b>goが動作する各OSに依存したAPIを扱う。</b>プラットフォームに独立したAPIを提供するやつ
* <code>os.Exit(ステータスコードを示すint型)</code>は<code>defer</code>すら置き去りにしてその場でプログラムを終了する
    * 引数には終了時に出力されるステータスコードを与える
* <code>log.Fatal(interface{}型)</code>はエラー発生時にエラーメッセージを出力しつつプログラムを終了する
    * > log.Fatalをとくに設定することなく使用 すると、引数として与えたinterface{}型の値 を標準出力へ出力したあとでos.Exit(1 が実行されます。

        とあったが、それを実現するコードのサンプルがなく、適当にlgo.Fatal()をつかわないだけだとなんのエラーもなくふつうに実行されてしまい確認できなかった
* <code>os.Args</code>はosパッケージのパッケージ変数
    * string型のスライス
    * コマンドライン引数が格納される
    ```
    package main
    import (
        "fmt"
        "os"
    )

    func main(){
        // os.Argsの要素数を表示
        // os.Argsは実行時に入力されたコマンドライン引数が入ってるらしい
        fmt.Printf("length: %d\n", len(os.Args))
        // 中身をぜんぶ出力
        for _, v:=range os.Args{
            fmt.Println(v)
        }
    }
    ```
    として<code>go build -o (バイナリで保存ｓたい名前) （このソースファイルの名前）</code>でbuildし、<code>（バイナリで保存した名前） 引数1 引数2 引数3...</code>とコマンドライン引数を与えてやると想像したとおりの出力が得られる