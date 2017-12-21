## Goのツール
### Goのツール群について
* <code>go run</code>, <code>go build</code>以外のコマンドとかみていくよ
#### goコマンドのとこで気になったやつ
* <code>go</code>だけだとヘルプになる
*
| (goに続く)コマンド | 意味 |
|:------------:|:------------|
| version | Goのバージョン <br> go version go1.9.2 darwin/amd64 |
| env | 環境変数とかGOPATHとか出てくる |
| fmt | litみたいなもん？コードを望ましい形に整形する。importとかインデントとか見てくれる<br>オプションでファイルを書き換えるかどうかとか表示するかとか指定できる<br>内部的に<code>gofmt</code>とかいうのを使っているらしく、さらに細かいオプションを指定できるらしい。通常は<code>go fmt</code>で事足りる |
| doc | パッケージのドキュメントを参照する。メソッド一覧とかも出る<br><code>パッケージ名/サブディレクトリのパッケージ名.そこに含まれる関数名や構造体やフィールド</code>でさらに細かく見れる(ex: <code>go doc math/rand.Zipf</code>) |
| build | ビルドするやつ。オプションがたくさんあってそれによっていろんな動作をするらしい<br>細かいしめんどいし割愛してつかうときになったらまたもどってこればいいよね |
| install | パッケージやら実行ファイルをGOPATHの場所にインストールする<br>パッケージ構成とかもう忘れているしなにいってんのかもはやわからんのだけど、ビルド対象mainパッケージが含まれている場合とかでなんか挙動が違ってくるらしい。まったくわからん<br> > <strong>Goのビルドシステムを理解する上で何よりも重要なのは、環境変数 GOPATHがどのような構成をとり、どのような役割を果たすのかについての理解です。プログラムのビルドとインストール、および環境変数 GOPATHの関係についてしっかりと理解できれば、Goのビルドシステムを十分に使いこなせるようになるでしょう</strong>,とのこと。ならもっとわかりやすく教えてくれ|
| get | 外部パッケージのDLとインスコをまとめて実行する。インスコの部分で内部的に上述の<code>install</code>コマンドを叩くらしい<br>GitHubとかBitBucketとかGoogleCodeProjectHostingとかからのダウンロードに対応している |
| test | テストするやつ。テストに関してもまっっっっったく知識がないのでぜんぜんわからんし読むの無駄っぽい。必要になったら戻ってくるわ |

### ベンダリング
* ひとつのGo環境で複数のアプリを開発中にそれらのいくつかが共通、ただしバージョンの違うパッケージに依存しているみたいなの（そもそもなんでそんな状況になったんだ）が問題になる
* GOPATHをいちいち変えればいいのだがめんどくさすぎる
* そこでベンダリングですよ
#### ベンダリングの実例のとこで気になったやつ
* <strong>パッケージに含まれる</strong>venderなるディレクトリが特別な意味を持つ。パッケージってとこがなんか気になる
* 開発しているパッケージが依存するパッケージを、venderディレクトリ以下におくことでGOPATHによる読み込みより優先的に読み込まれるらしい。
    * じゃあどんどん同じパッケージ（の異なるバージョン）が増えていく一方じゃないか。。。
    * pythonみたいに仮想環境使い捨てでやったれや
        * 調べたらgoenvとかいうのがあるらしい
        * でもまだgolangの開発環境ベストプラクティス的なのは確立されてないっぽい？dockerでまるっと使い捨てがいいような気がするんだが