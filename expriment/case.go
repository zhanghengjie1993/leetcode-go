package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 = []string{"hello", "liqing.io"}
	var s2 = s1
	fmt.Println("s2= ", strings.Join(s2, " "))
}
