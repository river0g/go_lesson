package main

import (
	"fmt"
)

func main() {
	/*　インタフェース
	- インタフェースと抽象化
		- 抽象化
			- 具体的な実装を隠し振る舞いによって共通化すること。
			- 複数の実装を同質の物として扱う。
		- インタフェースによる抽象化
			- Goではインタフェースでしか抽象化をすることができない

	例えば...
	具体的な処理Aがあり、ファイル1とファイル2で使用されているとする。
	これ抽象化で処理Aをまとめられる。

	- インタフェース
		- 型TがインタフェースIを実装しているとは、
			- インタフェースで定義されているメソッドを全て持つ。
			- 型TはインタフェースI型として振舞うことができる。
				- var i I = t // tはT型の変数とする。
		型T
			- メソッド1
			- メソッド2
			- メソッド3
			- メソッド4
		↓ インタフェース実装(抽象化)
		インタフェースI
			- メソッド1
			- メソッド2

	- インタフェースの例
		- インタフェースはメソッドの集まり。
			- メソッドのリストがインタフェースで規定しているものと一致する型はインタフェースを実装していることになる。
	*/
	type Stringer interface {
		String() string
		PrintType() string
	}
	// インタフェースを実装していることになる。
	var s Stringer = Hex(100)
	fmt.Println(s.String())
	fmt.Println(s.PrintType())
	/* 解説
	- Stringメソッドを抽象化(インタフェース実装)。
	- この抽象化した型(インタフェース)を変数の型にすることでインタフェースのメソッドが扱える。
	- 当たり前な話、メソッドの型とインタフェースを型に持つ変数の型はあっていないとダメ。
	*/

	/*
		- Q1. インタフェースの必要性
			- どうやって使えそうか
				- パッケージ化するとき
			- なんで必要なのか
				- 抽象化することで使いまわしやすくする。
			- インタフェースがない場合何が大変なのか
	*/

	/* interface{}
	- empty interface
		- メソッドセットが空なインタフェース
		- つまりどの型の値も実装していることになる。
		- JavaのObject型のような使い方ができる。
	*/
	var v interface{}
	v = 100
	v = "hoge"
	fmt.Printf("type: %T, value: %v\n", v, v)

	/* 関数にインタフェースを実装させる
	- 関数にメソッドを持たせる
	type Func func() string
	func (f Func) String() string { return f() }

	func main() {
		var s fmt.Stringer = Func(func() string {
			return "hi"
		})
		fmt.Println(s)
	}
	*/
	var s1 fmt.Stringer = Func(func() string { // 無名関数をFUnc型へキャスト(型変換)している
		return "interface!"
	})
	fmt.Println(s1)

	/* スライスとインタフェース
	- 実装していてもスライスは互換がない
		- コピーするには愚直にforで回すしかない
			- Go1.18からはジェネリクスがリリース。
	*/
	ns := []int{1, 2, 3, 4}
	// 以下はできない
	// var vs []interface{} = ns
	fmt.Println(ns)

	/* インタフェースの実装チェック
	- コンパイル時に実装しているかチェックする
		- インタフェース型の変数に代入してみる
	*/
	var _ fmt.Stringer = Func(nil)

	/* 型アサーション
	- インタフェース.(型)
		- インタフェース型の値を任意の型にキャストする。
		- 第2戻り値にキャストできるかどうかが返る。
		- 第2戻り値を省略するとパニックが発生する。
	*/
	var v2 interface{}
	v2 = 100
	n1, ok := v2.(int)
	fmt.Println(n1, ok)

	s2, ok := v2.(string)
	fmt.Println(s2, ok)

	/* 型スイッチ
	- 型によって処理をスイッチする
		- 代入文は省略可能
	インタフェース.(type)で型スイッチ可能。switch文の中でのみできる。
	*/
	var i1 interface{}
	i1 = 100
	switch v := i1.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "hoge")
	default:
		fmt.Println("default")
	}

	/* インタフェースの設計
	- メソッドセットは小さく
		- 共通点を抜き出して抽象化しない
		- 一塊の振る舞いを一つのインタフェースにする。
		- 型を使うユーザが触れる部分がインタフェースでなくてもよい
			- 内部にエンジンやドライバの形で抽象化した物を持つ
			- http.Client内部のhttp.RoundTripperのような感じ

	- 型階層は作れない
		- Goでは型階層は作れない。
		- 抽象化はすべてインタフェース。
		- 型階層ではなくコンポジットで表現する。
	*/

	/* io.Reader と io.Writer
	- 入出力の抽象化
		- 入出力を抽象化したioパッケージで提供される型
		- それぞれ1つメソッドしか持たないので実装が楽
		- 入出力をうまく抽象化し、さまざまな型を透過的に扱える
			- ファイル、ネットワーク、メモリ etc...
		- パイプのように簡単に入出力を繋げられる
	type Reader interface {
		Read(p []byte) (n int, err error)
	}
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
	*/

	/*
		- Q2. インタフェースを作ろう
			- Stringerインタフェース
				- String() string メソッドを持つインタフェースを作る
				- そして3つ以上Stringerインタフェースを実装する型を作る。
			- インタフェースを受け取る関数
				- Stringerインタフェースを引数で受け取る関数を作る。
				- 受け取った値を上記の3つの具象型によって分岐し、具象型の型名と値を表示する。
		- 解答
			- https://go.dev/play/p/U8y2GX6NWEC

	*/
	var j MyInt = 100
	F(j)

}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func (h Hex) PrintType() string {
	return fmt.Sprintf("%T", h)
}

type Func func() string

func (f Func) String() string { return f() }

type QStringer interface { // Stringerインタフェースを作成する。
	String() string
}
type MyString string
type MyInt int
type MyBool bool

func (s MyString) String() string {
	return "MyString"
}
func (b MyBool) String() string {
	return "MyBool"
}
func (i MyInt) String() string {
	return "MyInt"
}
func F(s QStringer) {
	switch v := s.(type) {
	case MyString:
		fmt.Println(string(v), "MyString")
	case MyInt:
		fmt.Println(int(v), "MyInt")
	case MyBool:
		fmt.Println(bool(v), "MyBool")
	}
}
