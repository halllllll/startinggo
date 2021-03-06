### goの実行ファイルのとこで気になったやつ
* goはOSのAPIに依存しないので指定したパッケージの機能すべてが組み込まれた実行ファイルになるのでコンパイルするとサイズ大きめの実行ファイルになる
* シンボル情報を読む <code>readelf</code>
    * Linux系の実行ファイルをELFファイルといい、それを読むコマンドが<code>readelf</code>らしい
    * [デバッグ情報の歩き方(Qiita記事)](https://qiita.com/mhiramat/items/8df17f5113434e93ff0c)
        > デバッグ情報とは、コンパイル言語においてソースコードから生成された実行可能なマシン語バイナリに対して、元々のソースコードとの対応付や、プログラム実行中のスタック解析を助けるための情報です。
    * 結論からいうと、『みてわかるとおりこれだけ多くのシンボルを組み込んでいるから実行ファイルが大きくなるんだね』でFA

### パッケージと構成のとこで気になったやつ
* export（呼び出し先）からは<b>頭文字を大文字で</b>関数を指定しないと<code>cannot refer to unexported name ~</code>で怒られる [^1]
    * もちろん呼び出し元(main)でもちゃんと大文字にしないといかん
    * [Goでpackageに定義した変数を参照できない](http://horie1024.hatenablog.com/entry/2014/08/25/012123)
* 自作パッケージを入れた同名フォルダと同じディレクトリ（main.goを含む）をカレントディレクトリにして<code>go build</code>するとカレントディレクトリと同名の実行ファイルが生成される
    * とくに指定のない場合の挙動
    * 勝手に名付けられるし勝手に.goファイルをビルドする
    * 曰く、goの想定したビルド構成に従えば最小限の手間でビルドできる
    * <code>go build main.go</code>とすると実行ファイル名は<code>main</code>となる
    * あとは思いつくようなよくあるオプションが用意されているっぽい。実行ファイルを生成するファイルパスを指定する<code>-o</code>とか

### mainパッケージの分類のとこで気になったやつ
* <code>main.go</code>と同じディレクトリにある.goファイルが<code>fmt</code>パッケージをインポートしているとビルドできない（<code>undefined: 関数名</code>とかで怒られる）
    * <code>fmt</code>パッケージをインポートしているファイルを全部列挙して<code>go run</code>するか、ワイルドカードを使う(<code>go run \*.go</code>)
    * <code>go build</code>はカレントディｋレクトリ内のgoファイルをすべてビルドするのでオプションをつけなければ通る

### パッケージとディレクトリのとこで気になったやつ
* <b>1つのディレクトリには1つのパッケージ定義のみ</b>。複数は<code>found package ~ and ~ </code>とかって怒られる

### パッケージのテストのとこで気になったやつ
* パッケージ内に<code>パッケージ名_test.go</code>なファイルを作る
    * これは予約語みたいな扱いで、この書式のファイルはパッケージをテストするためのおファイルとして扱われる
    * これ知らなくてハマったわ <b>goでなんか適当なの作るときに_test.goだとあかんのでsampleとかそういうのにしよう</b>。忘れそうだけど
* （べつにいま作ったファイルにすべてのテストを書く必要はなく、分けてもいいしそっちのほうが管理しやすそう）
* <code>testing</code>パッケージをインポート
* <code>Testテストする対象の関数名</code>というfuncを作る（ただしそこらへんに転がってるコードは別に頭にTestってついているだけであとはわりとどうでもよさそうな感じ??）
    * 引数は <code>TestFunc(t \*Testing T)</code>みたいな感じ

### テストの実行のとこで気になったやつ
* <code>go test テストしたいパッケージのディレクトリ</code>
    * あといろいろ考えつくようなよくあるオプションがあるっぽい

### mainパッケージのテストのとこで気になったやつ
* 外部パッケージのやつはその中に直接<code>パッケージ名_test.go</code>を作ったがmainパッケージのテストは<code>main.go</code>のあるディレクトリに、<code>mainパッケージをインポートしているファイル_test.go</code>で作る
* main.goがある場所をカレントディレクトリとして、<code>go test</code>
* あとは同じ

[^1]: スコープのとこででてくるが、識別子の名前の1文字目が大文字のものは外部に公開されているという話だった

### まとめ
* ファイル構成がクソ丁寧でめちゃくちゃわかりやすい、シンプルな図式で助かる
* <b>1つのディレクトリには1つのパッケージ定義のみ</b>。大事なことなので二回言いました
* テストが間違えないようがないってくらいちゃんと機能が用意されていてえらい
* エラーがなんか不親切じゃないだろうか、Go
