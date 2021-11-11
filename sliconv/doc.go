/*
sliconv はスライスの型変換を行うためのパッケージ。
[]bool, []float32, []float64, []int, []int32, []int64, []string の相互変換が可能。

元になるスライスと対応する構造体にスライスをセットし、 ToXXXX() メソッドで変換を行う。

	// サンプル: []bool から []string への変換
	b := []bool{true, false}
	s := Bool{b}.ToString()
	fmt.Printf("%#v -> %#v\n", b, s)
*/
package sliconv
