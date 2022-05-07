package app

// Iterator - 1つ1つの要素の処理を繰り返すためのもの
// 書籍では java.util.Iterator のコードが紹介されている
type Iterator interface {
	HasNext() bool
	Next() (*Book, error)
}

// NOTE: Pure な Iterator interface として実装したかったが、
//       Duck typing なので型を指定せざるを得ない。
//       また package も循環参照になってしまうため仕方なくまとめた。
