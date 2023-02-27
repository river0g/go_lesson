package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/* エラー処理
	- 正常系と委譲系
		- 正常系
			- 使用通りの動作
			- ユーザが意図通りに使った場合の挙動
				- マイクロサービスやライブラリのユーザも含む
		- 異常系
			- 意図しない挙動、発生頻度が低い挙動
			- ユーザが意図通りに使わなかった場合の挙動
			- 外部要因による意図しないエラー
				- ネットワーク、ファイル、ライブラリのバグ
			- バグが起因のエラー
	- エラー処理の必要性
		- エラーは必ず起きる
			- 外部要因で起きる可能性がある
			- 正常系より異常系の方が難しいし、パターンが多いことがある。
		- エラー処理
			- エラーが起きても処理を続けることができる場合もある
			- リトライをかけたり、別の方法をとることもできる。
			- 適切に処理してできる限り処理を続ける。
	*/

	/* エラー
	- errorインタフェース
		- エラーを表す型
		- 最後の戻り値として返すことが多い
	*/
	type error interface {
		Error() string
	}

	/* エラー処理
	- error型で表現する
		- エラーがない場合はnilになる。
		- 実行時エラーを表す
	- nilと比較して絵rーあが発生したかをチェックする
		- スコープを狭くするため、代入付きのifを用いる場合が多い
	*/
	if err := f(); err != nil {
		// エラー処理
	}

	/* エラー処理のよくあるミス
	- err変数を使い回すことによるハンドルミス
		- エラーが発生してもハンドルされずに次に進んでしまう
		- errcheckなどの静的解析ツールで回避できる
	*/
	file, err := os.Open("file.txt")
	if err != nil {
		// エラー処理
	}
	fmt.Println(file)

	f()             // 本来はerr = f()としたつもり
	if err != nil { // 絶対に実行されない
		// エラー処理
	}

	/* エラー処理で大事なこと
	- 必要十分な正しい情報を伝えること
		- エラーの持つ情報でエラーの原因を追える
		- 必要があれば情報を追加する
		- 無駄に情報を増やさない
	- 受け取り手によって伝え方を変える
		- 同じパッケージの別の関数なのか
		- 別のパッケージなのか
		- クライアントなのか
		- エンドユーザーなのか
	*/

	/*

	 */

	/* 文字列ベースで簡単なエラーの作り方
	- errors.Nowを使う
		err := errors.New("Error")
	- fmt.Errorfを使う
		- 書式を指定して、エラーを作る
		err := fmt.Errorf("%s is not found", name)
	*/
	err1 := errors.New("Error")
	fmt.Println(err1)

	/*

	 */

	/* エラー型の定義
	- Errorメソッドを実装している型を定義する。
		- そのエラー特有の情報を保持する。
		type PathError struct {
			Op string
			Path string
			Err error
		}
		func (e *PathError) Error() string {
			return e.Op + " " + e.Path + ": " + e.Err.Error()
		}
	*/

	/*
		- Q1. エラー処理をしてみる
		- Stringerインタフェースに変換する関数を作る
			- 任意の値をStringer型に変換する関数
				- type Stringer {String() string}
			- 引数にempty interfaceを取り、Stringerとエラーを返す。
				- func ToStringer(v interface{}) (Stringer, error)
			- 返すエラー型はerrorインタフェースを実装したユーザー定義型にする
			- 実際に呼び出してみてエラー処理をしてみよう
				- エラーが発生した場合はエラーが発生した旨を表示する。
	*/

}

func f() error {
	return nil
}

type User struct {
	Name string
	Age  int
	Err  error
}

func (e *User) Error() string {
	return e.Name + ": " + e.Err.Error()
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

type Stringer interface {
	String() string
}

func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, MyError("CastError")
}
