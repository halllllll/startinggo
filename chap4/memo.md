## 参照型とは
* よくある参照型の認識で構わない？
* goではスライス、マップ、チャネルのデータ構造が標準で相当する
### make関数
* スライス、マップ、チャネルはいずれも<code>make</code>関数によって作成される。
### スライス
* いちばん使うデータ構造。可変長配列のための型みたいなもん
* <code>s:=make([]int, 10)</code>という感じ。要素数10のint型のスライス
* 容量を指定する場合は
    ```
    s:=make([]int, 5, 10)
    ```
    という形。 これで要素数5, 容量10のスライスが定義される（後述）
* また、配列ライクに定義することも可能。この場合は要素数=容量となる
    ```
    s:=[]int{1, 2, 3, 4, 5}
    ```
* 見た目はあまり配列と変わらん。配列数が表示されるかどうか
    * <code>reflect.TypeOf</code>や<code>fmt.Prinltf("%#v", slice)</code>で型を表示
* 代入も参照も配列と同じ文法で実現できる
    * スライスの要素数を超過した要素へのアクセスはパニックを起こす
    * <code>panic: runtime error: index out of range</code>
* <code>len</code>で要素数を得られる
* <code>cap</code>で容量capacityを得られる
* 容量はスライスが将来使うかもしれない領域を予め確保しておいたもの
    * 可変長配列とはいえGo的には新たにメモリ領域を確保し直すのは負荷が高い処理らしい
    * したがって予めどの程度の容量が必要になるかの見積もりもパフォーマンスのためには必要
* 配列やスライスをもとにして新たなスライスを作ることができる。<b>簡易スライス式</b>とかいうらしい
    ```
    array:=[5]int{10, 4, 2, -5, 3}
    slice:=array[2:]
    ```  

* 簡易スライス式は文字列にも使えるが、配列のとこでやったように<b>文字列をバイト列としてみなしたうえでのインデックスを指定する必要がある</b>。utf8だと日本語は１文字につき3byte必要なので3区切りのインデックスを指定してやらねばならない
* <code>append</code>でスライスを拡張する
    * <code>newslice := append(slice, sliceの要素型の要素1, 要素2, ....)</code>
    * また、スライス同士の結合も可能。その場合は第二引数の末尾に...をつける
    ```
    slicea := make([]int, 4)
	sliceb := []int{3, 3, 0}
	slicec := append(slicea, sliceb...)
	fmt.Println(slicec, cap(slicec))
    >> [0 0 0 0 3 3 0] 8
    ```
    ちなみに3つ以上のスライスの結合はどうやるのかわからんしできるとも限らない

#### copyのとこで気になったやつ
* スライスにスライスの値をコピーするときの組み込み関数
    ```
    slicea := make([]int, 4)
    sliceb := []int{3, 3, 0}
    n := copy(slicea, sliceb)
    fmt.Println(n, slicea)
    >> 3 [3 3 0 0]
    ```
    戻り値はなぜか<b>「コピーされた要素数」</b>であることに注意。実際には第一引数のスライスに第二引数のスライスを上書きコピーする。
* コピー先の要素数のほうが大きい場合はそのまま先頭から順にコピーされるだけだが、逆にコピー元のほうが大きい場合はコピー先の容量を越えて増えることはない。上のコードでsilceaとslicebを逆にしてcopyすると
    ```
    >>> 3 [0 0 0]
    ```
    になる
#### 完全スライス式のとこで気になったやつ
* 配列やスライスからスライスを作るときにその容量をコントロールする
* 通常、要素10容量10の配列またはスライスから簡易スライス式を使って部分スライス[2:4]を抽出すると要素数2, 容量8のスライスになる
    ```
    関係: 0<=low<=high<=max<=cap(もとの配列またはスライス)
    完全スライス式: a[low:high:max]（aは配列またはスライス）
    簡易スライス式: a[low:high]

    ①簡易スライス式で新たに作られたスライスの容量=len(a)-low
    ②完全スライス式で新たに作られたスライスの容量=max-low
    ```
    すなわち、①より、要素10容量10から[2:4]として抽出たスライスの容量は10-2=8となる。
    完全スライス式で作るときはこの8を越えない範囲で容量の最大値を指定できるというわけ
* 教科書の図を見たほうが速い

#### スライスとforのとこで気になったやつ
* rangeを使った範囲節for。配列とおなじ
    * スライスの各要素に対して操作を行うようなループだと、たとえばループの条件節に要素数をiが越えないみたいな感じにしたうえでループ内でappendしたりすると無限ループになる（わかりにくい
    * rangeを使ったループだとそういうことにはならない。rangeの場合はループ開始時点での要素数を対象とするのでスライスをループ内で変更しても影響はない（わかりにくい
* 要は、<b>スライスをループしたいときはrangeを用いた範囲節forでやれ</b>ということ

#### スライスと可変長引数のとこで気になったやつ
* <code>fmt.Printf()</code>とかは可変長引数、それにスライスが使われている
* 可変長引数をとる関数
    ```
    func sum(slices ...int)int{
        // ぜんぶ足すだけ
        ans:=0
        for _,v := range slices{
            ans+=v
        }
        return ans
    }
    ```
    <code>slices ...int</code>のように型名の前に<code>...</code>をつけているところがポイント。
* 可変長引数は引数の最後にひとつだけ指定できる。pythonと似たようなもん
* スライスを可変長引数に渡す場合は変数名の前に<code>...</code>をつける
    ```
    //スライスを可変長引数として使う
    integers:=[]int{4, 2, 3, 6, 0, 10, -20}
    fmt.Println(sum(integers...))
    ```
#### 参照型としてのスライスのとこで気になったやつ
* 基本的にはごく普通のいわゆる参照型の扱いとおなじでいい
* Goに特徴的な基本型と参照型の違いは、<b>nilを値として取りうるかどうか</b>である、らしい
    * と思ったけどべつにGoに特徴的ってことはないな
#### スライスの落とし穴のとこで気になったやつ
* 配列からスライスを抽出する場合、コピーではなくそのまま配列の要素を参照する。よってこの段階で配列の要素に代入したりするとスライスのほうもその影響を受ける
* ただし、appendによる拡張などによって新たにメモリ領域が割り当てられ、スライスの参照先が変化した場合、当然もとの配列を操作してももうスライスには影響しない
* あーいわれてみれば確かに、という挙動。要注意

### マップ
* 連想配列
* makeを使って生成する場合
    ```
    m:=make(map[int]string)
    ```
    int型をキーに、string型を値にもつマップ,となる
* 容量の概念はないが、makeを使って生成するときに第二引数に「だいたいこんくらいの要素必要っぽい」って整数を与えてやるとランタイムに優しいらしい
    * 少ない要素数では無意味らしいけど
#### マップのリテラルのとこで気になったやつ
* 宣言
    ```
	m2:=map[int]string{1: "富士", 2: "鷹", 3: "ダーーーーーーッ!!!!"}
	m3:=map[int]string{
		10: "眠い",
		21: "だるい",
		32: "死にたい",	//, 必須
	}
    ```
    という感じで作れる。ふたつめのやつで最後にカンマをつけないと<code>unexpected newline, expecting comma or }</code>って怒られる
#### 要素の参照とのとこで気になったやつ
* ふつうに変数名[キー]で値にアクセスできる
* 存在しない値へアクセスしようとすると、変数名[キー]の段階で初期値が与えられてしまう
    ```
	//存在しない値にアクセスすると値が初期化され定義される
	v1 := m3[1000]
	fmt.Println(v1=="")
    >>> true
    ```
* という問題があるので、次のようにする
    ```
	v2, ok := m3[10]
	fmt.Println(v2, ok)
    >>> 眠い true
	v3, ok := m3[999]
	fmt.Println(v3, ok)
    >>>  false
    ```
* Goのマップは値にnilを指定できる。値が参照型であるようなマップを取り扱うときは要注意
    * なのでokイディオムを使うやつを積極的に多用する
* <b>lenが使える</b>。地味に便利そう
* 消すときは組み込み関数<code>delete</code>を使う。
    ```
    delete(マップ, キー)
    ```
    ちなみにキーが存在しなくてもとくにエラーは吐かない

#### マップとforのとこで気になったやつ
* rangeを使って範囲節forをかける
    * 配列や文字列やスライスだとrangeから返るのはインデックス、値という順番だが、マップではキー、値という感じ
* 当然ながら順序は保証されない

### チャネル
* ゴルーチンでデータの受け渡しをするときに使う
#### チャネルの型のとこで気になったやつ
* <code>chan 型</code>で定義する
    ```
	// int型のチャネルchを定義
	var ch chan int
    ```
* サブタイプ？とかいうやつで、受信専用か送信専用かを表現できる。  
    ```
	// 受信専用のチャネルukeを定義
	var uke <-chan int
	// 送信専用のチャネルsemeを定義
	var seme chan<- int
* サプタイプのついていない、ただの<code>chan</code>は、サブタイプつきの<code><-chan</code>（受信専用）<code>chan<-</code>（送信専用）のどちらにも代入可能
    * それ以外の組み合わせの代入は不可能
    * 基本的に送信専用or受信専用という用途で使われるのがいいらしい
#### チャネルの生成と送受信のとこで気になったやつ
* <code>make</code>で生成
    ```
    //ふつうのやつ（バッファサイズは0になる）
    ch:=make(chan int)
    //バッファサイズを指定したバージョン
    ch:=make(chan int, 10)
    ```
* 生成時に第二引数にチャネルのバッファサイズを整数で指定できる。指定しないと0
    * GoにおけるチャネルとはQueueであるらしく、どれだけQueueにためていられるか、がバッファサイズ
* <code>len</code>でデータの個数を得られる
* <code>cap</code>でバッファサイズを得られる

#### チャネルとゴルーチンのとこで気になったやつ
* 基本的にゴルーチン間でデータの媒介となるのがチャネル
    * なのでそもそもゴルーチンをもたないソース（mainのみ）ではデッドロックを起こしたりする
* レシーバーとなる関数に受信用のチャネルを引数として渡し、<code>go</code>キーワードでゴルーチンに登録
    ```
    func receiver(uke <-chan int){
        // 受信専用でint型を格納するチャネルを引数にとる
        // 当然送受信可能のチャネル型でもいい
        for {
            //i := <-uke
            //そのまま受信してもいい
            fmt.Println(<-uke)
        }
    }
    //どっかでチャネルを作っとく
    ch := make(chan int)
    //どっかでゴルーチンに登録する。あと引数にこの関数とやり取りするためのチャネルを渡す
    go receiver(ch)
    //あとは適当にチャネルにデータ送信
    chtest <- 999
    >>> 999
    ```
    <strong>ちなみになぜか何もない<code>for</code>で囲ってやらないと動かない。意味不明</strong>
* ループ内でやるパターン、パターンってほどでもない
    ```
    //関数でチャネルいじる
    func receiver2(i int, ch chan int){
        ch <- i*i
    }

    //forrrrrrrrrrrrrrrrrrrrrrrrrr
    for i:=0; i<100; i++{
        //ゴルーチンに登録してチャネルにiを送信
        go receiver2(i, chtest2)
        // <-chtest2　の部分で受信している
        fmt.Printf("chtest2:%d\n", <-chtest2)
    }
    ```
* 教科書の図解が一番わかりやすい

* <b>バッファサイズが0またはバッファ内が空のチャネルからの受信</b>および<b>バッファサイズが0またはバッファに空きがないチャネルへの送信</b>はデッドっロックを引き起こす
    ```
    deadlockch := make(chan rune, 3)
    deadlockch <- 'A'
    deadlockch <- 'B'
    deadlockch <- 'C'
    // deadlockch <- 'D'これすると死

	//空のバッファを受信するとデッドロック
	killmebaby := make(chan int)
	fmt.Println(<-killmebaby)
    ```
#### closeのとこで気になったやつ
* チャネルはopen/closeなる状態をもつ
    * <code>make</code>で生成されたチャネルはopen
    * <code>close</code>でチャネルを閉じる
* close状態のチャネルに送信しようとするとパニックを引き起こす
* close状態のチャネルから受信は可能。
    * バッファが空になっても受信するたびチャネルの型の初期値を返し続ける（パニックにならない）
* 受信<code><-</code>したときに左辺に２つの変数を割り当てるようにすれば第二引数でチャネルがopenかcloseかわかる
    ```
    ch := make(chan int)
    close(ch)
    i, ok := <- ch // i == 0, ok == false
    ```
    * 厳密には<b>「バッファが空かつclose状態」</b>のときにok==falseとなる
* テスト [Goroutineとclose](https://github.com/halllllll/startinggo/blob/master/chap4/goroutineandclose.go)
* rangeを用いた範囲節forを使ってチャネルからデータを取り出す感じのやつもできる
    * ただしcloseのタイミングが検知できない。当然空のチャネルから受信しようとするとデッドロックをキメることになる
#### selectのとこで気になったやつ
* <code>selet</code>とは
    * 制御構文
    * 節
    * 複数のチャネルを効率的にコントロールするやつ
    * たとえば1つのチャネルが受信に失敗したときに別のチャネルが受信を待ち続ける感じになってしまう、のを防ぐ
    * switch的な感じでつかう
    * のだけど、良い例が一切思いつかんし、ググっても（わかるくらいシンプルな）良いサンプルが一切無いのでサンプルソースを書きようがない。書こうとしたけど2時間くらいわけわからんとこでハマってやる気をなくした
        * なので剽窃
    ```
    select{
        case e1 := <- ch1 /*ch1からの受信が成功した場合の処理 */
        case e2 := <- ch2 /*ch2からの受信が成功した場合の処理 */
        default: /* case節の条件が成立しなかった場合 の処理 */
        }
    ```
    * このあともう一個exampleで「複数のゴルーチンを使ったselect」ってやつがあるんだけど、「case節で実行可能なやつからランダムに処理する」とか言ってたくせしてどうみても順番に処理していては？？？？？？？？？？？となったのでもうどうでもいいいです

### まとめ
* スライスは可変長配列
* スライスは要素数と容量という概念を持つ
* 容量を越えた範囲でスライスを拡張しようとすると、自動的に容量を増やして確保してコピーする
    * これはなかなかコストが高い処理
* スライスをループするときはrange使え
* appendで拡張したときに参照先変わるよ
* Goのマップは値にnilを指定できる。値が参照型であるようなマップを取り扱うときは要注意（二回目
* チャネルには送受信できるやつと送信用、受信用がある
* チャネルは様々なゴルーチンで共有されうるのでたとえ<code>len</code>やら<code>cap</code>で取った値は瞬間的に動的に変化し続ける可能性がある
* ゴルーチンに登録する関数内でチャネルをつかうときは<strong>なぜか何もないforで囲ってやらないとdeadlockとかっつって怒られる</strong>（二回目）。まことに意味不明で1時間以上ハマった
* チャネルのopenとcloseを意識しようね
* <code>select</code>節で複数のチャネルを管理
* selectとかチャネルとかゴルーチンはクソ（わからんという意味で）