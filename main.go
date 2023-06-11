package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 入力
	fmt.Print("入力：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	v := strings.Split(scanner.Text(), " ")
	var a, b *float64
	var s string
	for i, v2 := range v {
		if i%2 == 0 {
			// iが偶数の時は数字
			a = b
			c, err := strconv.ParseFloat(v2, 64)
			if err != nil {
				fmt.Println("数字じゃないよ！")
				return
			}
			b = &c
			if a != nil {
				switch s {
				case "+":
					*b = *a + *b
				case "-":
					*b = *a - *b
				case "/":
					*b = *a / *b
				case "*":
					*b = *a * *b
				default:
					fmt.Println("エラー！")
					return
				}
			}
		} else {
			// iが偶数じゃないとき(奇数の時)
			switch v2 {
			case "+":
				fallthrough
			case "-":
				fallthrough
			case "/":
				fallthrough
			case "*":
				s = v2
			default:
				fmt.Println("記号じゃないよ！")
				return
			}
		}
	}
	fmt.Printf("答えは%f", *b)
}
