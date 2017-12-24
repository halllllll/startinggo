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
* <code>os.Open()</code>はファイルをreadonlyモードで開く
    * つかうときは最後にストリームを切るために<code>defer filename.Close()</code>しとくと良いらしい
* <code>os.File</code>は ~~なんかファイルとかディレクトリの操作とかにばんばん使う感じのやつらしい~~ Openしたりしたときに返る構造体でとりあえず思いつく感じのファイル操作ができる。CreateとかRenameとかRemoveAllとかそういう感じ
    * os.File型のファイルへのIOは<code>[]byte</code>型をつかう
    * なのでファイルを扱うときはとりあえずまずそれ作って、あとは元のファイルのどこからbyteを読み込むかとかそういう感じ
    * <code>os.Fle.Stat()</code>で返る<code>os.FileInfo</code>型でファイル名やファイルサイズやファイルのモードやファイルの最終更新時間やディレクトリかどうかなんてのがわかる
    * [写経した](https://github.com/halllllll/startinggo/blob/master/chap7/file2.go)ものの、ReadAtとか動かし方が違うのか思った動作してないしゴミ

* 読み取り・書き取りやファイルが無い場合に新規作成する、みたいなモードを指定してファイルをオープンするときには<code>os.OpenFile</code>という別の関数を使う
    ```
    f, err := os.OpenFile("ファイルの名前", オプション, パーミッション)
    ```
    * オプションは

    | オプション（フラグ） | 意味 |
    |:----------------:|:----------------|
    | O_RDONLY | 読み取り専用 |
    | O_WRONLY | 書き取り専用 |
    | O_RDWR | 読み書き可能 |
    | O_APPEND | ファイルの末尾に追加 |
    | O_CREATE | ファイルが存在しなければ新規作成 |
    | O_TRUNC | ファイルの内容をオープン時に空にする（?)<br>TRUNCはtrancate,切り取るって意味らしい。DB界隈ではよく使われるっぽい？まあリストのスライスみたいなもんで、 <i>~~ググったけどろくな情報がなくて全然わからんかった</i>ので無視しよう、無視~~ 内容を削除して開くって感じらしい。試してないので知らんけど|

    * なお、オプションはパイプ演算子でつなげることができる。たとえば<code>O_RDWR | O_CREATE | O_APPEND</code>は、読み書き可能でファイルが存在しなかったら新規作成して存在してたらその末尾に追加、ってこと
    * 第三引数のパーミッションが意味わからんかった（教科書にはその説明がまったく載っていなくてスルーされてる）ので調べると[網羅的なやつが出てきた。](http://waman.hatenablog.com/entry/2017/10/01/130330)正直これがないと直上のTRUNCもわからんかった（試してないので結局わかってないけど）
        * しらんけどPOSIXたらいうUnix起源ななんかカーネルをうんたらする用のAPIらしい。全然わからんけど
* <code>os.Getwd</code>で現在のカレントディレクトリを取得できる
    ```
    package main
    import (
        "fmt"
        "os"
        "log"
    )

    func main(){
        dir, err := os.Getwd()
        if err!=nil{
            log.Fatal(err)
        }
        fmt.Println(dir)
    }
    ```
    * あと<code>os.Chdir("存在するディレクトリ")</code>でカレントディレクトリを変更できる、らしい
* <code>*os.File</code>型のメソッド<code>os.Readdir()</code>でディレクトリ以下のファイル情報を読み取れる
    * 試してない
    * 引数に0だとすべてのファイル、正の整数だと読み込む最大値を指定
* <code><os.MkdirMk</code>でディレクトリを作成できる
    ```
    err := os.Mkdir("ディレクトリ名", パーミッション)
    ```
    という感じ
    * 深い階層を一度に作るために<code>so.MkdirAll</code>もあるらしい。例によって試してないので知らんけど
#### osパッケージのその他のファイル操作
* テンポラリディレクトリを取得する<code>os.TempDir</code>
* シンボリックリンクを作成する<code>os.Symblink</code>
* シンボリックリンクを読み込む<code>os.ReadLink</code>
* ホスト名を取得する<code>os.Hostname</code>
* 環境変数名を[]stringで取得する<code>os.Environ</code>
    * 
    ```
    package main
    import (
        "fmt"
        "os"
    )

    func main(){
        // 環境変数を取得する
        for _, v := range os.Environ(){
            fmt.Println(v)
        }
    }
    ```
* 環境変数名が存在するか確認する<code>os.Getenv("環境変数名")</code>
* 環境変数をセットする<code>os.SetEnv("環境変数名")</code>
* 環境変数をすべてリセットする<code>os.Clearenv</code>
* 環境変数をチェックしつつその値を（列挙せず）逐次参照する<code>os.LookupEnv("環境変数名")</code>
* などなど
* 実行中のプロセスを取得する<code>os.Getpid()</code>
    * win機にはそもそも<b>ユーザーIDやグループIDは取得できない</b>らしい。取得できない場合は常に-1が返るらしい
### time
* 日付とか時刻ですね
* time.Now
    * お察しのとおり現在時刻が取得できる。<code>time.Time</code>型の構造体。
    * <code>time.Time</code>構造体はtimeパッケージにおける中心的なデータ構造
* time.Data
    * 指定した日時を表す<code>time.Time</code>型を任意に生成できる
    * 年を表す<code>time.Time.Year()</code>はじめ、<code>YearDay</code>,<code>Month</code>,<code>Weekday</code>,<code>Day</code>,<code>Hour</code>,<code>Minutee</code>,<code>Second</code>,<code>Nanosecond</code>,<code>Zone</code>みたいなのが取れる
    * 詳しくは教科書にあるんだけど、わざわざ書くまでも無いので都度ググれ
    * ちなみによくある言語仕様と異なり、<strong>
    月（Month）は1から始まるらしい</strong>。へー
