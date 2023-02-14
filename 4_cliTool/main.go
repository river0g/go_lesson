package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
	"log"
)

var msg = flag.String("msg", "デフォルト値", "説明")
var n int
func init() {
	flag.IntVar(&n, "n", 1, "回数")
}

func main() {
	fmt.Println("hi")
	/* プログラム引数
		- プログラム実行次に渡される引数
			- プログラムに対して外から渡されるデータ・情報
			- コマンドライン引数 = プログラム引数
		echo hello
		hello

		- プログラム引数を取得する
			- os.Argsを使用する
				- プログラム引数が入った文字列型のスライス
				- 要素のひとつめはプログラム名
	*/
	fmt.Println(os.Args) // go run main.go hello -> [/var/.../main hello]
	// ビルドした後、./main [./main hello]



	/* flagパッケージ
		- フラグ(オプション)を便利に扱うパッケージ
		文字列を扱うには、
		flag.Strng("フラグ名", "デフォルトの値", "フラグの説明")
		数値を扱うには、
		flag.IntVar(int型のポインタ, "フラグ名", デフォルト値(int), "フラグの説明")
		このファイルではinit関数を使ってflagの設定を初期化している。

		// 設定される変数のポインタを取得
		var msg = flag.String("msg", "デフォルト値", "説明")
		var n int
		func init() {
			// ポインタを指定して設定を予約
			flag.IntVar(&n, "n", 1, "回数")
		}
		func main() {
			// ここで実際に設定される
			flag.Parse()
			fmt.Println(strings.Repeat(*msg, n))
			// ./main -msg=こんにちは -n=2
		}
	*/
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))


	/* flagパッケージとプログラム引数
		- flag.Args関数を用いる
			- os.Argsだとフラグも含まれる ./main -msg=hello hi -> [./main -msg=hello hi]
			- flag.Argsだとフラグの分は除外される。
	*/
	fmt.Println(flag.Args())



	/* 入出力
	- 標準入力と標準出力
		- osパッケージで提供されている*os.File型の変数
			- さまざまな関数やメソッドの引数として渡せる。
			- エラーを出力する場合は標準エラー出力に出力する。
				- 標準入力: os.Stdin
				- 標準出力: os.Stdout
				- 標準エラー出力: os.Stderr
	*/

	/* fmt.Fprintln関数
		- 出力先を指定して出力する
			- 末尾に改行をつけて表示する
			- os.Stdoutやos.Stderrに出力できる
			- ファイルにも出力できる
	*/
	fmt.Fprintln(os.Stderr, "エラー") // 標準エラー出力に出力
	fmt.Fprintln(os.Stdout, "Hello") // 標準入力に出力


	/* プログラムの終了
		- os.Exit(code int)
			- 終了コードを指定してプログラムを終了
			- プログラムの呼び出し元に終了状態を伝えられる
			- 0: 成功(デフォルト)
	*/
	fmt.Fprintln(os.Stderr, "エラー")
	// os.Exit(1) // これをコメント解除すると下のPrintlnは出力されない(こっから下のプログラムは実行されない)
	fmt.Println("hi")
	

	/* プログラムの終了 (エラー)
		- log.Fatal
			- 標準エラー出力(os.Stderr)にエラーメッセージを表示
			- os.Exit(1)で異常終了させる
			- 終了コードがコントロールできないためあまり多用しない
	*/
	if err := f(); err != nil {
		log.Fatal(err)
	}
}

func f() any {
	return "error"
}