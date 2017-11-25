package foo
const (
	MAX = 100
	internal_const = 1
)

//　名前の1文字目が大文字なので外部から参照可能
func FooFunc(n int) int{
	//関数に入ってしまえばこのgoファイルの関数にアクセスできる
	return internal_Func(n)
}

func internal_Func(n int) int{
	return n+n
}