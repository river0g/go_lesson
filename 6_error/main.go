package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
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

	/* エラー処理をまとめる
	- bufio.Scannerの実装が参考になる。
		- 途中でエラーが発生したらそれ以降の処理を飛ばす
		- すべての処理が終わったらまとめてエラーを処理
		- それ以降の処理を実行する必要ない場合に使う
		- エラー処理が1箇所になる
	*/
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := s.Text()
		if text == "exit" {
			break
		}
		fmt.Println(text)
	}
	if err := s.Err(); err != nil {
		// エラー処理
	}

	/*

	 */

	/*
		- Q2. エラ＝処理をまとめる
		- 1コードポイント(rune)ずつ読み込むScannerを作る。
			- 初期化時にio.Readerを渡す
			- bufio.Scannerと似た感じに
			- エラー処理をまとめる
	*/

	s1 := NewRuneScanner(strings.NewReader("Hello, 世界"))

	for {
		r, err := s1.Scan()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%c\n", r)
	}

	/*

	 */

	/* エラーをまとめる
	- https://github.com/uber-go/multierr を使う
		- 成功したものは成功させたい
		- 失敗したものだけエラーとして報告したい
		- N番目エラーはどういうエラーなのか知れる

	var rerr error
	if err := step1(); err != nil {
		rerr = multierr.Append(rerr, err)
	}
	if err := step2(); err != nil {
		rerr = multierr.Append(rerr, err)
	}
	return rerr

	このとき、以下のように扱える(まとめられる)
	for _, err := range multierr.Errors(rerr) {
		fmt.Println(err)
	}
	*/

	/* エラーに文脈を持たせる
	- github.com/pkg/errors を使う
		- エラ〜メッセージが「File Not Found」とかでは分かりづらい
		- 何をしようとした時にエラーが起きたか知りたい
		- どんなパラメータだったのか知りたい
		- errors.Wrapを使うとエラーをラップできる
		- errors.Causeを使うと元のエラーが取得できる

	if err := f(s)l err != nil {
		return errors.Wrapf(err, "f() with %s", s)
	}
	*/

	/* エラーに文脈を持たせる(Go 1.13)
	- fmt.Errorf関数の%wを使う
		- 引数で指定したエラーをラップしてエラーを作る
		- Unwrapメソッドを実装したエラーが作られる
		- errors.Unwrap関数で元のエラーが取得できる
	err := fmt.Errorf("bar: %w", errors.New("foo"))
	fmt.Println(err) // bar: foo
	fmt.Println(errors.Unwrap(err)) // foo
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

type RuneScanner struct {
	r   io.Reader
	buf [16]byte
}

func NewRuneScanner(r io.Reader) *RuneScanner {
	return &RuneScanner{r: r}
}

func (s *RuneScanner) Scan() (rune, error) {
	n, err := s.r.Read(s.buf[:])
	if err != nil {
		return 0, err
	}

	r, size := utf8.DecodeRune(s.buf[:n])
	if r == utf8.RuneError {
		return 0, errors.New("RuneError")
	}

	s.r = io.MultiReader(bytes.NewReader(s.buf[size:n]), s.r)
	return r, nil
}
