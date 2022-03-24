package main

import (
	"fmt"
)

type State int
type CharType int

const (
	STATE_INITIAL State = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
	STATE_END
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_SPACE
	CHAR_ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		STATE_INITIAL: {
			CHAR_SPACE:  STATE_INITIAL,
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
			CHAR_SIGN:   STATE_INT_SIGN,
		},
		STATE_INT_SIGN: {
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
		},
		STATE_INTEGER: {
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_EXP:    STATE_EXP,
			CHAR_POINT:  STATE_POINT,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT: {
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT_WITHOUT_INT: {
			CHAR_NUMBER: STATE_FRACTION,
		},
		STATE_FRACTION: {
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_EXP: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SIGN:   STATE_EXP_SIGN,
		},
		STATE_EXP_SIGN: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
		STATE_EXP_NUMBER: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SPACE:  STATE_END,
		},
		STATE_END: {
			CHAR_SPACE: STATE_END,
		},
	}
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_NUMBER || state == STATE_END
}

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	status := make([]int, (2*k+1)*len(prices))
	for i := range dp {
		dp[i] = status[:2*k+1]
		status = status[2*k+1:]
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	ans := maxProfit(2, []int{3, 2, 6, 5, 0, 3})
	fmt.Printf("%d", ans)
}
