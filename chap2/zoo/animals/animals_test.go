package animals

//testingモジュール？パッケージ？をインポート
import (
	"testing"
)


//Test+パッケージ内のgoファイルの関数名
//よくわからん引数を取る
func TestWaniFeed(t *testing.T){
	//想定する戻り値？
	except:="ワニのくうようなやつ"
	//実際の戻り値
	actual:=WaniFeed()

	if except!=actual{
		t.Errorf("%s != %s", except, actual)
	}
}

func TestGoriFeed(t *testing.T){
	except:="ゴリラの主食はマシュマロです"
	actual:=GoriFeed()

	if except!=actual{
		t.Errorf("%s != %s", except, actual)
	}
}

func TestHipoFeed(t *testing.T){
	except:="カバが食べそうなやつ"
	actual:=HipoFeed()

	if except!=actual{
		t.Errorf("%s != %s", except, actual)
	}
}