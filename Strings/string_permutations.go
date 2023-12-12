package Strings

import "fmt"

func Permute() {
	var input = "ABC"
	var result []string
	permuteHelper([]rune(input), 0, &result)
	fmt.Println(result)
}

func permuteHelper(s []rune, start int, result *[]string) {
	if start == len(s)-1 {
		*result = append(*result, string(s))
		return
	}

	for i := start; i < len(s); i++ {
		swap(s, start, i)
		permuteHelper(s, start+1, result)
		swap(s, start, i)
	}
}

func swap(s []rune, i, j int) {
	s[i], s[j] = s[j], s[i]
}
