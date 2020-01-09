package chapter1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/apbgo/go-study-group/chapter1/lib"
)

// Calc opには+,-,×,÷の4つが渡ってくることを想定してxとyについて計算して返却(正常時はerrorはnilでよい)
// 想定していないopが渡って来た時には0とerrorを返却
func Calc(op string, x, y int) (int, error) {

	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf

	// TODO Q1

	//switch op {
	//case "+":
	//    return 6, nil
	//case "-":
	//    return 2, nil
	//case "×":
	//    return 10, nil
	//case "÷":
	//    return 2, nil
	//default:
	//    return 0, fmt.Errorf("Invalid op=%s", op)
	//}

	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "×":
		return x * y, nil
	case "÷":
		return x / y, nil
	default:
		return 0, fmt.Errorf("Invalid op=%s", op)
	}

}

// StringEncode 引数strの長さが5以下の時キャメルケースにして返却、それ以外であればスネークケースにして返却
func StringEncode(str string) string {
	// ヒント：長さ(バイト長)はlen(str)で取得できる
	// chapter1/libのToCamelとToSnakeを使うこと

	// TODO Q2

	//if str == "sn_ak" {
	//	return "SnAk"
	//}
	//if str == "asHGkdJahf" {
	//	return "as_h_gkd_jahf"
	//}
	//return ""

	if l := len(str); l <= 5 {
		return lib.ToCamel(str)
	}
	return lib.ToSnake(str)
}

// Sqrt 数値xが与えられたときにz²が最もxに近い数値zを返却
func Sqrt(x float64) float64 {

	// TODO Q3

	//if x == 9.0 {
	//	return 3.0
	//}
	//return 6.164414002968976

	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z

}

// Pyramid x段のピラミッドを文字列にして返却
// 期待する戻り値の例：x=5のとき "1\n12\n123\n1234\n12345"
// （x<=0の時は"error"を返却）
func Pyramid(x int) string {
	// ヒント：string <-> intにはstrconvを使う
	// int -> stringはstrconv.Ioa() https://golang.org/pkg/strconv/#Itoa

	// TODO Q4
	//if x == 5 {
	//	return "1\n12\n123\n1234\n12345"
	//}
	//return "1\n12\n123\n1234\n12345\n123456\n1234567\n12345678\n123456789\n12345678910"

	s := ""
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			s += strconv.Itoa(j)
		}
		if i < x {
			s += "\n"
		}
	}
	return s
}

// StringSum x,yをintにキャストし合計値を返却 (正常終了時、errorはnilでよい)
// キャスト時にエラーがあれば0とエラーを返却
func StringSum(x, y string) (int, error) {

	// ヒント：string <-> intにはstrconvを使う
	// string -> intはstrconv.Atoi() https://golang.org/pkg/strconv/#Atoi

	// TODO Q5

	//if x == "2" {
	//	return 12, nil
	//}
	//return 0, fmt.Errorf("Invalid =%s", x)

	a, err1 := strconv.Atoi(x)
	if err1 != nil {
		return 0, err1
	}

	b, err2 := strconv.Atoi(y)
	if err2 != nil {
		return 0, err2
	}

	return a + b, nil
}

// SumFromFileNumber ファイルを開いてそこに記載のある数字の和を返却
func SumFromFileNumber(filePath string) (int, error) {
	// ヒント：ファイルの扱い：os.Open()/os.Close()
	// bufio.Scannerなどで１行ずつ読み込むと良い

	// TODO Q6 オプション
	//return 55, nil

	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sum := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			return 0, err
		}
		v, err := strconv.Atoi(sc.Text())
		if err != nil {
			return 0, err
		}
		sum += v
	}

	return sum, nil
}
