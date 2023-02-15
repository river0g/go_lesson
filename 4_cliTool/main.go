package main

import (
	"fmt"
	"os"
	"flag"
	// "strings"
	"bufio"
	"path/filepath"
)

// var msg = flag.String("msg", "デフォルト値", "説明")
var n int
func init() {
	// flag.IntVar(&n, "n", 1, "回数")
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
	// flag.Parse()
	// fmt.Println(strings.Repeat(*msg, n))


	/* flagパッケージとプログラム引数
		- flag.Args関数を用いる
			- os.Argsだとフラグも含まれる ./main -msg=hello hi -> [./main -msg=hello hi]
			- flag.Argsだとフラグの分は除外される。
	*/
	// fmt.Println(flag.Args())



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
	/* nilでなければ(errorであれば)異常終了させる。
		if err := f(); err != nil {
			log.Fatal(err)
		}
	*/


	/* ファイルを扱う
		- osパッケージを用いる
		// 読み込み用にファイルを開く
		sf, err := os.Open(src)
		if err != nil {
			return err
		}
		// 関数終了時に閉じる
		defer sf.Close()

		- 書き込み用にファイルを開く
		df, err := os.Create(dst)
		if err != nil {
			return err
		}
		// 関数終了時に閉じる
		defer func() {
			if err != df.Close(); err != nil {
				rerr = err
			}
		}()
	*/

	/* 追記用のファイル
		todo...

	*/



	/* defer
		- 関数の遅延実行
			- 関数終了時に実行される。
			- 引数の評価はdefer呼び出し時
			- スタック形式で実行される(最後に呼び出したものが最初に実行)

		msg := "!!!"
		defer fmt.Println(msg) // 3番目に実行される。ここではmsg==!!!
		msg = "world"
		defer fmt.Println(msg) // 2番目に実行される。ここではmsg==world
		fmt.Println("hello") // 1番目に実行される。

		上記の出力結果
		hello
		world
		!!!
	*/
	// msg1 := "!!!"
	// defer fmt.Println(msg1)
	// msg1 = "world"
	// defer fmt.Println(msg1)
	// fmt.Println("hello")

	/* for内のdeferは避けよう
		- 予約した関数呼び出しはreturn時に実行される
			- forを関数に分ければよい
		for _, fn := range fileNames {
			f, err := os.Open(fn)
			if err != nil {
				// エラー処理
			}
			defer f.Close()
			
			// fを使った処理
		}

		** for内でdeferを避けるので、関数化する。**

		func main() {
			for _, fn := range fileNames {
				err := readFile(fn)
				if err != nil { // エラー処理 }
			}
		}
		func readFile(fn string) error {
			f, err := os.Open(fn)
			if err != nil { return err }
			defer f.Close()

			// fを使った処理
		}

		// なぜこれがいいのかわからないので動画で確認する。
	*/



	/* 入出力関連の便利パッケージ
		- encoding: JSONやXML, CSVなどのエンコードを扱うことができる
		- strings: 文字列周りの処理がある
		- bufio: Scannerが便利
		- strconv: 文字列への変換を行う関数を提供
		- unicode: Unicode周りの処理を提供
	*/


	/* 1行ずつ読み込む
		- bufio.Scannerを使用する
		// 標準入力から読み込む
		scanner := bufio.NewScanner(os.Stdin)
		// 1行ずつ読み込んで繰り返す
		for scanner.Scan() {
			// 1行分を出力する
			fmt.Println(scanner.Text())
		}
		// まとめてエラー処理をする
		if err := scannerErr(); err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
		}
	*/
	
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	text := scanner.Text()
	// 	if text == "exit" {
	// 		fmt.Println("see you!")
	// 		break
	// 	}
	// 	fmt.Println(text)
	// }


	/* ファイルパスを扱う
		- path/filepathパッケージを使う
			- OSに寄らないファイルパスの処理が行える // winだと \main でmacは /main など...
		// パスを結合する。
		path := filepath.Join("dir", "main.go")
		// 拡張子を取る。
		fmt.Println(filepath.Ext(path))
		// ファイル名を取得
		fmt.Println(filepath.Base(path))
		// ディレクトリ名を取得
		fmt.Println(filepath.Dir(path))
	*/

	/* ディレクトリをウォークする
		// Goファイルを探し出す
		err := filepath.Walk("dir",
			func(path string, info os.FileInfo, err error) error {
				if filepath.Ext(path) == ".go" {
					fmt.Println(path)
				}
				return nil
			}
		)
		if err != nil {
			return err
		}
	*/



	/* 
		- Q1. catコマンドを作ろう
			- 作成するcatコマンドの仕様
				- 引数でファイルパスの一覧を貰い、そのファイルを与えられた順標準出力順に標準出力にそのまま出力するコマンドを作ってください
				- また、-n オプションを指定すると、行番号を各行につけて表示されるようにしてください。
				- なお、行番号は全てのファイルで通し番号にしてください。
			例. $ mycat -n hoge.txt fuga.txt
			1: hoge
			2: hoge hoge
			3: fuga
			4: fugafuga
	*/
	// mycat()

	fmt.Println("********************************")
	const dir = "/Users/0g/myapp/go_lesson/4_cliTool/resource"
	// フラグを受け取る
	var qn bool
	flag.BoolVar(&qn, "qn", false, "行番号付与")

	flag.Parse()

	// ファイル名を受け取る
	var args []string = flag.Args()

	// ファイルを出力
	for idx, fn := range args{
		filePath := filepath.Join(dir, fn)
		rf, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
		}
	
		scanner := bufio.NewScanner(rf)
		for scanner.Scan() {
			if qn {
				fmt.Printf("%v: ", idx)
			}
			fmt.Println(scanner.Text())
		}
	}

	

	fmt.Println("********************************")
	//
}

func f() any {
	return "error"
}

func mycat() {
	flag.Parse()

	// フラグを受け取る。
	


}