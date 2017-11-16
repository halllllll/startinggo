### 導入のとこで気になったやつ
* 相変わらずGOPATHがよくわからんが、とりあえず空のフォルダを作ってそこを指定しておく。
    * <code>go help gopath</code>
    * 外部のパッケージをインストールしたりビルドするのにかならず必要
    * GOPATHに指定した各ディレクトリは所定の構造になっていないといけない。srcディレクトリ、pkgディレクトリ、binディレクトリ。
    *<code>
    GOPATH=/home/user/gocode

    /home/user/gocode/
        src/
            foo/
                bar/               (go code in package bar)
                    x.go
                quux/              (go code in package main)
                    y.go
        bin/
            quux                   (installed command)
        pkg/
            linux_amd64/
                foo/
                    bar.a          (installed package object)
</code>
    * そもそもver1.8から勝手にGOPATHが設定されるようになったので考える必要なし←
    * <code>go env</code>でいろいろ確認できる
    * winだとC直下にGOROOTに指定されたフォルダができており、GOPATHはuser直下になっていた（見えてない？）
        * 最初に<code>go get</code>で適当なやついれたら作ってくれたっぽい
    * winでVSCでやるときにそのままRunnerだと動かない（そりゃそうだ）
        * setting.jsonからGOPATHを設定する →　できない(?????
        * 再起動してまたRunしようとしたらgopkgを入れてね的なやつがでてきたのでAll Install
        * →go-outlineとかgo-symbolsとかいろいろが勝手にインストールされるのを待つ
        * <b>All tools successfully installed. You're ready to Go :).</b>
        * → ダメ
        * ターミナルから<code>go run</code>するか、デバッグならイケたのでRunnerを使うのは諦めよう。。使い方いまいちわからんし