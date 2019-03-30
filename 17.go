package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if (len(digits)) == 0 {
		return nil
	}
	AllDigits := map[byte][]string{
		'0': {" "},
		'1': {""},
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	var cur string
	answer := make([]string, 0)
	dfs(digits, AllDigits, 0, cur, &answer)
	return answer
}
func dfs(digits string, AllDigits map[byte][]string, l int, cur string, answer *[]string) {
	if l == len(digits) {
		*answer = append(*answer, cur)
		return
	}
	for _, n := range AllDigits[byte(digits[l])] {
		cur = cur + n
		dfs(digits, AllDigits, l+1, cur, answer)
		cur = cur[:int(len(cur)-1)]
	}
}
func main() {
	fmt.Print(letterCombinations("23"))
}
