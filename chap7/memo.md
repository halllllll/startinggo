## Goのパッケージ
* 使用頻度の高そうな標準パッケージをみていくよ
* クソ多いのでざっといくよ（気づいたのだいぶあとなので急に内容が加速します）
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
* time.Month型、time.Weekday型の<code>String</code>メソッドは月名や曜日名を文字列で取得できる
* <code>time.Duration</code>型で<b>時間の間隔を表現</b>できる
    * <code>time.Hour</code>とか<code>time.Nanosecond</code>とかで取れる。最小値が初期値。時間だったら<code>"1h0m0s"</code>とかでナノ秒だったら<code>"1ns"</code>とか
    * <code>time.Time</code>型と組み合わせることができる。time.Time型は<code>time.Now()</code>とか<code>time.Date(時間指定)</code>とかの返り値
    ```
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
    ```
    * あと、時間の前後関係をboolで返す<code>Before</code>と<code>After</code>とか同じかどうかを返す<code>Equal</code>とかある
* <code>AddDate</code>で年月日を増減
    ```
	// 数値指定で増減
	// 1年後
	t1:=t.AddDate(1, 0, 0)
	fmt.Println(t1)
	// 20年と4ヶ月と40日前
	t2:=t.AddDate(-20, -4, -40)
	fmt.Println(t2)
    ```
* <code>time.Parse(フォーマット, 指定の日時)</code>でフォーマットに沿った日時を生成する
    * 教科書なにいってんのかわからんので無視
* <code>time.Format(フォーマット)</code>で任意のフォーマットにできる
    * 無視
* 以下だいたい意味不明なので無視

### math

* だいたいよくあるような定数とかMaxInt的なやつは当然揃っている
* 平方根はよくあるけど<strong>立方根</strong>もある。<code>math.Cbrt()</code>
* <code>math.Trunc()</code>で小数点以下を切捨て。負値もふつうに切り捨てるので0のほうに丸まる
* Ceil, Floorもふつうにある。負値の取扱に注意
* やたらと三角関数が充実している
* 対数もLogだけじゃなくて<code>Log10</code>, <code>Log1p</code>, <code>Log2</code>とかある
* <strong>非数</strong>なんて久しぶりに聞いたが非数かどうか判定する<code>IsNaN</code>がある
* 非数とか無限大も生成できる。<code>math.NaN</code>, <code>math.inf(0)</code>

### math/rand
* みんな大好き疑似乱数
    ```
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Float64())
    ```
    よくあることだがシード値が固定だと実行のたび毎回同じ乱数になる
* デフォルトでは範囲は0から1
* 範囲は最大値だけ指定できるので負値がほしけりゃあとで減算するなりする
    ```
	fmt.Println(rand.Intn(100))
    ```
* デフォルトの疑似乱数発生器はGoのランタイム上のやつを共有しているのでどっかでシード値を書き換えるとすべての場所で使ってる擬似乱数もその影響を受ける・このため独自に擬似乱数発生機を生成できる仕組みがある
    * <code>rand.NewSource(ソースとなる値)</code>でソースを生成
    * <code>rand.New(ソース)</code>で新たに擬似乱数発生器を生成

### flag
* シンプルなコマンドラインを作れる(?)
    * ~~コマンドラインツールじゃなくてコマンド自体を作る？？~~ まあなんかやっぱParseArgと同じと思っていいっぽい
* ArgParse的なもんかと思ったが、<code>os.Args</code>を使って作るとコマンドの順番とかそういうのを考慮するのが非常にめんどくさいらしいのでこれがあるらしい？？
* 一応つくったけどメモするまでもねぇっつぅかまた知りたくなったら必要なときに戻って来ればいいかな

### fmt
* ふつうの標準出力の<code>fmt.Print()</code>系と文字列としてフォーマットして生成する<code>fmt.Sprintf()</code>系とほかのio.wirter（ファイルとか）に出力する<code>fmt.Fprintf()</code>系がある（誤解を生む表現）
* 書式指定子がふつうに豊富にあるけどまとめて書くのめんどくさすぎる。とりあえずいつも通り0埋めとと左詰めと2, 8, 16進数で表現するやつと小数点以下どこまでって指定するやつくらいはかければいいや
* 書式指定子<code>%v</code>
    * いろんな型を柔軟に出力できる。基本型はもちろん<code>interface{}</code>のような型が不定であるやつでもマップでも配列でもスライスでも構造体でも。
    * 構造体に<code>%+v</code>を使うとフィールドについても出力し、<code>&#v</code>だとフィールドに加えて型についても出力する、とか
### log
* 簡単なログを標準エラー出力に出力する。標準出力と標準エラー出力の違いは知らん
* ログの出力先を変更できる。標準入力とか別のファイルとか。（VSC上で試したけどどっちもわからん、後者に至ってはファイル生成されないし）
* <code>log.SetFlags(フォーマット)</code>でログをフォーマットできる。フォーマットは論理和？パイプ演算子でつなげて複数指定できる
    ```
	log.SetFlags(log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	log.Println("おぱい")
    ```
    とかにすると、日付、時間、マイクロ秒、ファイルの絶対パスも合わせて出力される
* <code>log.SetPrefix(プレフィックス)</code>でプレフィックスを設定できる
* ロガーを生成できるらしい。例によって何書いてあんのかさっぱりわからんので無視
### strconv
* 基本型とstringの相互変換みたいなやつらしい。bool値や整数値、floatなんかをstringに変換できる
    * intを変換するときに基数を指定できる。つまりn進数で表現できる
    * floatはなんかオプションがいろいろあって指数表現だったり実数表現だったり小数点以下n桁までの指定できたり精度をビット数で指定したりするもうわけわからんね
* 文字列からbool, int, floatそれぞれに変換するときも可能であったり不可能であったりするような書き方がある
    * TRUEや1やtでもtrueとして扱われる、とか
    * この場合もintに変換するときに基数を設定できる
    * ただし<strong>負値は使えない</strong>

### unicode
* > unicodeはrune型が表現するUnicodeコードポイント処理のためのユーティリティーがまとめられたパッケージです。
* 意味不明
    * 特に<strong>Unicodeコードポイント処理</strong>
* 文字か数値か、といった判定が可能
    * <code>IsDigit()</code>とか<code>IsLetter()</code>とか
### strings
* 文字列操作。検索や置換や結合やetc
* <code>strings.Join()</code>は<code>[]string</code>型に含まれる文字列を結合する。第二引数につなげる文字を指定できる（指定しないとダメなのでいらんときは空文字で）
* <code>Index</code>とか<code>Join</code>とか<code>Contains</code>とか<code>Repeat</code>とか<code>ToLower</code>とか<code>ToUpper</code>は当たり前として、なんか珍しいのがある
    * <code>IndexAny(検索対象となる文字列, 含まれるか調べる文字列)</code>は第二の文字のうち第一引数に最初に含まれるもののインデックスを返す
    * <code>HasPrefix(ry, ry)</code>は第一引数が第二引数から始まる文字列かどうかをboolで返す
    * <code>HasSuffix</code>は逆
    * <code>Contains(ry, ry)</code>は部分文字列が含まれるかどうかをboolで返すが、<code>ContainsAny(ry, ry)</code>は部分文字列じゃなくていずれかの文字が含まれるかどうかを返す。こんなん第二引数スライスか配列でいいような気がするが
    * <code>Count(ry,ry)</code>はよくある数えるやつだけどなぜか<strong>空文字で検索するとすべての文字が対象になる</strong>ので注意(len(s1)になる)
    * <code>Replace(対象となる文字列, 含まれるかどうか検索する文字列, 置換する文字列, 置換最大数)</code>でそのまんまの感じ。長い。
    * <code>TrimeSpace(対象となる文字列)</code>は空文字やスペースやタブや改行なんかを削除する
    * <code>Fileds(対象となる文字列)</code>はスペースほか改行とかタブで区切られた文字列を分割して[]string型にする
### io
* > このパッケージ自体にはほとんど機能はありませんが、他のパッケージにおける入出力処理などを 取り扱う場合に、ioパッケージについての知識が必要 になります。

#### io.Reader
* なんらかのバイト列の入力を抽象化するinterface。
* メソッドは<code>Read</code>だけ
* <code>os.File</code>が実装してるのでそこでよく見かけるらしい
#### io.Writer
* 出力処理を抽象化したinterface
* 同じく<code>Write</code>だけもつ

### bufio
* 基本的なIO処理に<b>バッファ処理を付加した機能</b>がまとめられたパッケージ    
    * ？？？？？
* 多機能
* 手順として、
* <code>bufio.NewScanner(io.Readerを実装した何か)</code>で適当に<code>os.Stdin</code>なり<code>strings.NewReader</code>なりの入力ソースを与えてやり、スキャナを生成
* スキャナの<code>Scan()</code>関数はスキャンが成功する限りtrueを返す
* スキャナの<code>Text()</code>関数はスキャンした内容を文字列にする
* という感じ
* スキャナを生成した際デフォルトでは<b>行区切り</b>でスキャンすることになっているが、スキャナの<code>Split(区切り方)</code>関数でどういう挙動にするか変更できる
    * 区切り方は<code>SplitFunc</code>型とかいうやつらしいが、よくわからんので以下のやつを使うといい
    * <code>bufio.ScanLines</code>は改行区切り
    * <code>bufio.ScanWords</code>は空文字やスペースや行区切り
    * あと<code>bufio.ScanRunes</code>と<code>bufio.ScanBytes</code>もあるらしいが知らん
* バッファ使うと早くなる。出力でバッファ使うときは最後に<code>Flush</code>しないと出力されない

### io/ioutil
* 入出力をサポートする機能をまとめたもの
* いままでのIOとどう違うっていうんだ。。。
* たとえば入力全体を読み込める
    * <code>ioutil.ReadAll(io.Reader型)</code>
    * 戻り値は[]byteで、あんまり巨大なファイルとかには向いてない
* ファイル全体を読み込む
    * <code>ioutil.ReadFile(ファイル名)</code>
    * 上と同じじゃねぇか（上のは入力はなんでもいいけど）
    * こっちを使うとわざわざファイルを開く手続きが必要なくなる
* テンポラリファイルを作れる
    * <code>ioutil.TempFile(ディレクトリのパス, ファイル名のプレフィックス)</code>
    * なんに役に立つのかしらん

### regexp
* どうみても正規表現
* ダルいので無視
* なんかフラグの立て方が特殊らしいよ　どうでもいいけど
### json
#### 構造体をJSON式に
* <code>json.Marshal</code>で<code>フィールド名:値</code>って形にしてくれる
    * <code>json.Unmarshal</code>でその逆もできる
* いつか出てきた（忘れた）構造体のタグ機能で<code>タグ:値</code>って形にもできる
* ほかにももっとあるやろとおもったらそれしか書かれてなくて草

### net/url
* URL文字列をパースする
* pythonで見たぞ requestsだっけ？
* <code>url.Parse(URL文字列)</code>の戻り値である<code>url.URL</code>型のフィールドからurlの各要素を取れる
    * <code>Scheme</code>とか<code>RawQuery</code>とか<code>Fragment</code>とか色々
* <code>url.URL</code>から逆にurlを生成したりもできる。各クエリも<code>Encode()</code>でうまい具合にしてくれるっぽい
    ```
	// これポインタでやらないと結果が違ってくるからな
	u:=&url.URL{}
	u.Scheme = "https"
	u.Host = "soundcloud.com"
	u.Path = "djjasrac"
	// Query()の戻り値はurl.Values型で、これはmap[string][string]のエイリアスであるらしい
	q := u.Query()
	// なんかmapがSet()もつとか調べてもどこにも出てこなかったんだけどまあもういいや、Setできるってことだろ
	q.Set("key1", "yo")
	q.Set("key2", "muri")
	// url.Encode()でいい具合に=と&でつなげてくれるってことだろ
	fmt.Println(q.Encode())
	u.RawQuery = q.Encode()
	fmt.Println(u)
	fmt.Println(q)
	fmt.Println(reflect.TypeOf(q))
    ```
    出力
    ```
    key1=yo&key2=muri
    https://soundcloud.com/djjasrac?key1=yo&key2=muri
    map[key1:[yo] key2:[muri]]
    url.Values
    ```
### net/http
* httpクライアント/サーバー用
* めんどｋさいので飛ばすか・・・
### sync
* 非同期処理。ゴルーチンとなにが違うのか
* > チャネルは同期処理のあらゆる局面の解決策であるというわけではありません。
    * 具体的にどういう局面がダメなのか一切の説明がない
* > Goのチャネルを使った解決策も考えられますが、ここではsync.Mutex型によって提供されるミューテックス機構による排他制御を加えてみましょう。
    * ミューテックスがなんなのかについての説明は一切ない
* <code>sync.Mutex</code>型の<code>Lock</code>と<code>Unlock</code>関数で単一のゴルーチンのみ処理させ、別のゴルーチンが来たときはいまのゴルーチンが終了するまで待たせる、っていうやつする
    * LockとUnlockに挟まれたブロックというか間は単一のゴルーチンが走る
* <code>sync.WaitGroup</code>型で任意のゴルーチンの<b>終了を待つ</b>仕組みを提供している
    * <code>Add(いくつのゴルーチンを待つか int)</code>で処理を完了するまで待つゴルーチンの数を指定
    * 待ち受けたい箇所に<code>Wait()</code>を仕掛ける
    * <code>Done()</code>が指定した回数実行されるまで待つ
        * なので教科書ではgo funcの後にDone()しているのだがこれが果たして汎用的な場所なのかは不明

### crypto
* MD5, SHA-1, SHA-256, SHA-512などのハッシュ値を生成
* なるほど(???????????????)
