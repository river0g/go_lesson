package main

import (
	"time"
	"math/rand"
)

func main() {
	/* 
		Q1. 変数の利用
		出力する文字列を変数に格納する。
		println("Hello world!")
	*/
	var msg string = "Hello world!"
	// msg := "Hello world!"
	println(msg)

	
	// 数字の区切りは_で行う => 数値リテラルで_は無視される
	// 可読性を上げるために導入された。
	num1 := 500_000_000
	println(num1)

	
	// 定数式
	// 定数(100, "hoge", 'A' などのリテラルのこと)の計算
	println(100+300) // 四則演算
	println("100"+"100") // 文字列結合
	println(1 << 2) // シフト演算
	println(!(100 == 200)) // 論理演算/関係演算

	// 名前付き定数 - 定数に名前をつけて宣言する
	// 型があったりなかったりすることができる。
	const n int = 100

	// 定数式の利用もできる。
	const s = "Hello" + "世界"


	/*
		100などの数値リテラルに型がない。デフォルトの型はある(100ならint, "AA"はstring, 'A'はrune など)
		この型なしの定数が必要な理由として、100 * 0.2 は int * float なので型変換が必要になるが、それぞれ型なしにすることでただの数字の計算(型変換がいらない)ができる。
	*/


	/*
		Q2. 定数の利用
		println("Hello, world")
		を定数として定義してから出力する。
	*/
	const cmsg = "Hello, world"
	println(cmsg)


	// まとめて代入する	- 名前付き定数定義の右辺の省略
	// 以下のa,b,cは同じ値が入る。
	const (
		a = 1 + 2
		b
		c
	)
	println(a, b, c)


	// iota - 連番定数生成、型なしの整数
	const (
		aio, bio = iota ,iota // 0, 0
		cio = iota // 1
		dio, eio = iota, iota // 2, 2  
	)

	// iotaを使用した連番
	const (
		StatusOK = iota // 0
		StatusNG // 1
	)


	/*
		代入演算
		= 変数への代入 a = 100
		:= 変数の初期化と代入 a := 100
		+=, -= 演算と代入 i+=2
		++, -- インクリメントとデクリメント。ちなみにこれは式ではなく文

		ビット演算
		| 論理和, & 論理積, 
		^ 否定　^0xc, ^ 排他的論理話 0xc^0x3
		&^ 論理積の否定, << 左に算術シフト, >> 右に算術シフト

		論理演算
		|| or , a || b
		&& and, a && b
		! 否定, !a

		比較演算
		== 等しいかどうか
		!= 等しくないか
		< aはbより小さい a<b, <= aはb以下 a<=b
		> aはbより大きい a>b, >= aはb以上 a>=b

		アドレス演算 - 他のポインタを扱える言語(Cなど)と違い、ポインタ演算はできない。
		& ポインタを取得 &a
		* ポインタがサス値を取得 *b

		チャネル演算
		<- チャネルへの送受信 ch<-100, <-ch
	 */


	// 制御構文
	// 条件分岐 if
	// 条件に()がいらない。
	x := 1
	if x == 1 {
		println("xは1")
	} else if x == 2 {
		println("xは2")
	} else {
		println("xは1でも2でもない")
	}

	// 代入文を書く
	if a := 10; a > 0 {
		println("aは0より大きい")
	} else {
		println("aは0より小さい")
	}


	// 条件分岐 switch
	// jsと違い、breakがいらない。
	switch a {
	case 1, 2:
		println("a is 1 or 2")
	default:
		println("default")
	}

	// caseに式が使える -> これにより上のswitchでは==のみだったものが評価式を使うことができる。
	switch {
	case a == 0:
		println("a is 0")
	case a == 1:
		println("a is 1")
	default:
		println("default")
	}


	/*
		繰り返し for
		whileはない。Goでは繰り返しはforのみ。
		
		for 初期化文; 継続条件式; 後処理文 のように書く。
		for i:=0; i<=100; i=i+1 {
		}

		初期化文, 継続条件式, 後処理文の記述は任意。

		初期化文のみ
		for sum0:=0;; {
			sum0 ++
			if sum0 > 10 {
				break
			}
		}

		継続条件式のみ
		sum1 := 0
		for ;sum1<10; {
			sum1 ++
		}

		後処理文のみ。継続条件を書かないか、breakしないと無限ループになるので注意
		sum2 := 0
		for ;;sum2 = oneUp(sum2) {
			if sum2 > 10 {
				break
			}
		}

		無限ループ
		for {
		}

		継続条件飲みの時は;;を省略できる。
		for i<=100 {
		}

		range を使って繰り返し
		for i, v := range []int{1,2,3} {
		}

		名前付きのbreak(ラベル指定のbreak)
		LOOP:
			for sum3:=0;;sum3=oneUp(sum3) {
				if sum3 > 5 {
					break LOOP
				}
			}

	*/



	/*
		Q3. おみくじプログラムを作ろう
		- サイコロを転がして出ためによって運勢を占うプログラム
			- 6: 大吉
			- 5,4: 中吉
			- 3,2: 吉
			- 1: 凶
	*/

	t := time.Now().UnixNano()
	rand.Seed(t)
	diceNum := rand.Intn(10)
	switch diceNum {
	case 6:
		println("大吉")
	case 5,4:
		println("中吉")
	case 3,2:
		println("吉")
	case 1:
		println("凶")
	default:
		println("サイコロの目にありません。エラー。")
	}
}