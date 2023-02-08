package main

import "fmt"


func main() {
	/*
		変数と型
		- 型
			- どういう値かを示すもの。
			- ユーザー定義型という自分で作成することも可能
		- プログラミングには動的型付け(Pythonなど)、静的型付け(Goなど)がある。

		静的型付けの利点
		- 実行前に型の不一致を検出できる
			- コンパイルが通れば方の不一致が起きない。
			- 型の不一致によるバグは見つけづらい問題なので検出することは役立つ
		- 曖昧なものはエラーになる
			- 暗黙の型変換がない。(jsはある)
				1 + "2" => "12"
			- floatとintの演算など見つけづらいバグが起きにくい
		- 型推論がある
			- 型推論があることにより明示的に型を書く必要がない時が多い。

		組み込み型
		- int, float, string, booleanなど
		- int8などbit数を示してつける型もある

		型変換(型キャスト)
		- ある型から別の型に変換すること
		- 変換できない場合はコンパイルエラーになる
			- "hoge"は文字列にしかならないのでintにはできない。
			- 10.1をintに型変換すると10になる。(コンパイラーにはならない)
		var a int = 10
		aa := string(a)
		pritnln(aa) // "10"
		
	*/
	var f float64 = 10.0
	var n int = int(f)
	println(n)


	/*
		Q1. 次のプログラムはコンパイルが通るか否か
		var sum int
		sum = 5 + 6 + 3
		avg := sum / 3
		if avg > 4.5 {
			println("good")
		}

		- コンパイルエラーになる。理由はavgがintになるから。
		- 修正するにはsumとavgに型つけるか、ifの時にavgをfloat32にキャストする。
	*/

	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	if float32(avg) > 4.5 {
		fmt.Println(avg)
	}

	/* コンポジット型について
		コンポジット型(複合型)
		- 複数のデータ型が集まって一つのデータ型になっている
			- 構造体: 型の異なるデータ型を集めたデータ型
			- 配列: 同じ列のデータを集めて並べたデータ型
			- スライス: 配列の一部を切り出したデータ型
			- マップ: キーと値をマッピング(対応)させたデータ型 // Pythonでいう辞書。Jsでいうオブジェクト、PHPでいう連想配列

		コンポジット型のゼロ値
		- スライスやマップはmake関数で初期化が必要なため、nilとなる
			- 構造体: フィールドが全てゼロ値
			- 配列: 要素が全てゼロ値
			- スライス: nil
			- マップ: nil
		
		型リテラル
		- 具体的な型表現
			- []int, map[string]int など
	*/



	/* 構造体
		- 各変数はフィールドと呼ばれる
			- 以下の例では構造体pがフィールドnameとageを持つ。
		- フィールドの型は異なっても良い。
			- 下記だとnameはstring, ageはint
		- フィールドの値には組み込み型以外のコンポジット型やユーザー定義型も使える。
		- 型リテラルはstruct{...}
	*/
	var p1 struct {
		name string 
		age int
	}
	fmt.Println(p1)

	/* 構造体リテラル
		フィールドを指定して初期化(構造体リテラル)
	*/
	p2 := struct {
		name string
		age int
	}{
		name: "Jisoo",
		age: 28,
	}
	fmt.Println(p2)

	/* プログラミングの文法
		var 変数名 型
		var n int
		var p struct {
			name string
		}
		構造はどれも同じ。
	*/

	
	/* フィールドの参照
		「.」でアクセスする。
		参照も代入も.を用いてアクセスする。JSみたいなもの。
	*/
	p3 := struct {
		name string
		age int
	}{ name: "lisa", age: 26 }
	p3.age ++
	fmt.Println(p3.name, p3.age)



	/* 配列
		同じデータ型を集めたデータ構造
		- 要素数は変更できない
		- 型は型リテラルで記述することが多い。
	*/
	var ns [5]int
	fmt.Println(ns)

}