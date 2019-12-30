package main

import (
	"strconv"
	"strings"
)

const intmax = 2 ^ 31 - 1
const intmin = -2 ^ 31

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	res, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return res
}

func main() {

}
