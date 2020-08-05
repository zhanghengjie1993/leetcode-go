package main

<<<<<<< HEAD
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

=======
func removeDuplicates(nums []int) int {
	left, right, size := 0, 1, len(nums)
	for ; right < size; right++ {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
>>>>>>> d11fcca760bba99ecf1a8dd4934927b287a13b4e
}
