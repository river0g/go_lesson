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
	// ゼロ値で初期化
	var array1 [5]int
	fmt.Println(array1)

	// 配列リテラルで初期化
	var array2 = [5]int{10, 20, 30, 40, 50}
	fmt.Println(array2)

	// 要素数を値から推論
	array3 := [...]int{10,20,30,40,50}
	fmt.Println(array3)

	// インデックスを指定して値を入れることもできる。
	// 6番目が10、11番目が100の要素数が11の配列
	array4 := [...]int{5: 50, 10: 100}
	fmt.Println(array4)

	// 要素にアクセス
	fmt.Println(array4[3])

	// 長さ - Pythonと同じ
	fmt.Println(len(array4))

	// スライス演算
	// [n:m] => n以降からmの一つ手前までの要素を示す。
	fmt.Println(array4[1:2])



	/* スライス
		- 配列の一部を切り出したデータ構造
			- 配列と同じで要素の方はすべて同じ。
			- 要素数は方情報に含まない。 => 長さが途中で変えられるので追加、削除ができる。
			- スライスの背後には配列が存在する。

		- スライスの初期化
			- スライスリテラル([]intのようなもの)での初期化
				- 要素数の指定は不要
			- make関数を使って初期化
				- 長さと容量を指定する。
				- make([]int, 長さ, 容量)
	*/

	// make関数を使って初期化
	// ゼロ値が代入される
	ns1 := make([]int, 3, 10)
	fmt.Println(ns1)

	// スライスリテラルを使用して初期化
	var ns2 = []int{10,20,30}
	fmt.Println(ns2)

	// 配列と同じくインデックス指定で値を入れられる
	// 6番目が50、11番目が100で他の要素が0の要素数11のスライス
	ns3 := []int{5:50, 10:100}
	fmt.Println(ns3)
	

	/* スライスを配列の関係
		- スライスはベースとなる配列が存在している
		- スライス→配列で大体同じ処理を以下に示す

		スライス
		ns := make([]int, 3, 10)
		配列
		var array [10]int
		ns := array[0:3] // or array[:3]

		スライス
		ms := []int{10, 20, 30, 40, 50}
		配列
		var array2 = [...]int{10,20,30,40,50}
		ms := array2[0:5]
	*/


	/* スライスの操作
		- 要素にアクセス
			- ns[2]
		- 長さ
			- len(ns)
		- 容量
			- cap(ns)
		- 要素の追加
			- append関数を使う
			- append(要素を増やしたいスライス, 追加する要素)
			- 追加する要素はカンマ区切りで可変に指定できる。
	*/ 
	// 要素の追加
	// 容量が足りない場合は背後の配列が再確保される
	ns4 := []int{10,20,30,40,50}
	ns4 = append(ns4, 60, 70)
	fmt.Println(ns4)
	/* appendの挙動
		- 容量が足りる場合
			- 新しい要素をコピーする
			- lenを更新する
		- 容量が足りない場合
			- 元のおよそ2倍の容量の配列を確保し直す
				- 1024を超えた場合は、およそ1/2ずつ増える
			- 配列へのポインタを貼り直す
			- 元の配列から要素をコピーする
			- 新しい要素をコピーする
			- lenとcapを更新する。
	*/
	

	// 配列・スライスのスライス演算
	ns5 := []int{10,20,30,40,50}
	n, m := 2, 4

	// n番目以降のスライスを取得する。
	fmt.Println(ns5[n:]) // [30 40 50]

	// 先頭からm-1番目までのスライスを取得する
	fmt.Println(ns5[:m]) // [10 20 30 40]


	/* スライスの要素をfor文で取得する 
		- スライスをfor rangeのrangeに指定することで各要素を取得できる。
			- for i, v := range(スライス) {}
			- iはスライスの要素のインデックス、vはスライスの要素
	*/
	ns6 := []int{10,20,30,40,50}
	for i, v := range(ns6) {
		fmt.Println(i, v)
	}

	// Slice Tricks
	// カット
	ns7 := []int{10,20,30,40,50}
	ns7 = append(ns7[:2], ns7[3:]...)
	fmt.Println(ns7)

	// 削除
	ns8 := []int{10,20,30,40,50}
	ns8 = append(ns8[:3], ns8[4:]...) // or ns8 = ns8[:4+copy(ns8[3:], ns8[4:])]
	fmt.Println(ns8)

	/* x/exp/slicesパッケージ
		- スライスに関する便利なパッケージ
		- 将来標準ライブラリに入るかもしれない
		ns := []int{10,20,30,40,50}
		// 削除
		ns = slices.Delete(ns, 1, 3) // [10 40 50]
		// 挿入
		ns = slices.Insert(ns, 1, 60, 70) // [10 60 70 40 50]
	*/


	/* Q2. 以下のコードを3つの変数しか使わないコードにしてください。
		n1 := 19
		n2 := 86
		n3 := 1
		n4 := 12
		
		sum := n1 + n2 + n3 + n4
		println(sum)
	*/

	var q2_sum int
	nslice := []int{19, 86, 1, 12}
	for _, v := range(nslice) {
		q2_sum += v
	}
	println(q2_sum)



	/* マップ
		- キーと値をマッピングさせるデータ構造
			- キーと値の型を指定できる。
			- キーには「==」で比較できる型しかだめ
			- ゼロ値はnil

		var m map[string]int


		- マップ初期化
			- make関数での初期化
				- make(map[string]int)
			- リテラルでの初期化
				- map[string]int{"x": 10, "y": 20}
			- 空の場合
				- map[string]int{}
	*/

	/* マップの操作
		m1 := map[string]int{"x": 10, "y": 20}
		- キー指定でアクセス
			- m["x"]
		- キー指定で入力
			- m["z"] = 30
		- 存在確認, 存在しない場合はゼロ値とfalseを返す。
			- n, ok := m["z"]
		- キー指定で削除
			- delete(m, "z")
	*/
	m1 := map[string]int{"x": 10, "y": 20}
	fmt.Println(m1["x"])
	me, ok := m1["z"] // 0, false 
	fmt.Println(me, ok)


	/* マップの要素をfor文で取得する 
		- マップをfor rangeのrangeに指定することで各要素を取得できる。
			- for k, v := range(マップ) {}
			- kはマップの要素のキー、vはマップの要素の値
	*/
	m2 := map[string]int{"x": 10, "y": 20, "z": 30}
	for k, v := range(m2) {
		fmt.Println(k, v)
	}


	
	/* コンポジット型を要素にする
		- コンポジット型を要素として持つコンポジット型
			- スライスの要素がスライスの時(2次元スライス)
				- [][]int
			- マップの値がスライスの場合
				- map[string][]int
			- 構造体のフィールドの形が構造体
				- struct {
					A struct {
						N int
					}
				}
	*/



	/* ユーザー定義型
		- typeに名前を付けて新しい型を定義する
		- type 型名 型
			- 組み込み型を基にする
				- type MyInt int
			- 他のパッケージの型を基にする
				- type MyWriter is.Writer
			- 型リテラルを基にする
				- type Person struct {
					Name string
				}
	*/


	/* Underlying type
			
	*/


	/* ユーザ定義型の特徴
	
	*/


	/* 型エイリアス (Go 1.9以上)
		- 型のエイリアス(コピーみたいなもん)を定義できる
			- 完全に同じ型
			- キャスト不要
				- type Applicant = http.Client
		- 型名を出力する%Tが同じ元の型名を出す。
	*/

}