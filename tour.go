package main
// Golangのプログラムは Package(パッケージ) で構成される
// プログラムは main　から実行される
// パッケージ名はインポートパスの最後の要素と同じ名前にしなければならない

// インポートの仕方は以下
import (
    "fmt"
    "math"
)
// もしくは以下でもできるが、上のやり方が推奨される
// import "math"
// import "fmt"

// 変数の宣言
// var をつけて宣言する
var i int
var j int = 10
var c, python, java bool

// 定数の宣言
// const をつけて宣言する
// character, string, bool, numeric のみで定数宣言できる
const Pi = 3.14

// 関数内の使い捨て変数
// var の代わりに i := 1 のように初期化することができ、
// 暗黙的な型宣言ができる。関数の中以外では := は使用できない
func vari_example() (int) {
    i := 10
    return i
}


// 関数の書き方
// 関数は0個以上の引数を取ることができる。引数の方は
// 変数の後に書くこと。2つ以上の引数が同じ方の場合には
// (x int, y int) を (x, y int) のように型をまとめて表せる(省略できる)
func add(x int, y int) (int) {
    return x + y
}

// Exported names
// 大文字で始まる名前は、外部パッケージから参照できる
// エクスポートされた名前
// 下の例だと、math.Piのように、importしたmathから参照できる
// Piのこと
func print_pi() {
    fmt.Println(math.Pi)
}

// 戻り値を複数返すことができる
func swap(x, y string) (string, string) {
    return y, x
}

// 戻り値に名前をつけることができる
// 名前をつけると、関数の始めに定義した変数として扱うことができる
// 戻り値に名前をつけた場合、returnステートメントに何も書かなくても、
// 名前をつけた戻り値をreturnしてくれる(naked returnという)
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    print_pi()
}

//
