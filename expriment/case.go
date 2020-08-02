package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我是中国人"
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

}
