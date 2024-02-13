package main

import (
	"fmt"
	"unicode"
)

func main() {
	rev("Aku makan padang")
}

func rev(str string) {
	pisah := []rune(str)
	var s []string
	var res string
	min := 0
	for i := 0; i < len(pisah); i++ {
		if unicode.IsSpace(pisah[i]) {
			s = append(s, string(pisah[min:i]))
			min = i + 1
		}
		if i == len(pisah)-1 {
			s = append(s, string(pisah[min:i+1]))
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			res = s[i] + " "
		} else if i == 0 {
			res = res + " " + s[i]
		} else {
			res = res + s[i] + " "
		}
	}
	fmt.Println(res)
}
