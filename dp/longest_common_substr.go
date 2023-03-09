package dp

import "fmt"

func longestCommonSubstr(text1, text2 string) int {
	var longest uint8

	dp := make([]uint8, len(text2)+1)
	for _, r1 := range text1 { // r stands for rune
		p := []rune(text2)
		for j := len(p) - 1; j >= 0; j-- {
			// iterate from the last the first
			// because we reuse the array dp
			// the later ones will refer to the earlier ones
			// this it to prevent being overwritten
			if r1 == p[j] {
				fmt.Println(dp[j])
				dp[j+1] = dp[j] + 1
				if dp[j+1] > longest {
					longest = dp[j+1]
				}
			} else {
				dp[j+1] = 0
			}
		}
		fmt.Println(dp)
	}
	fmt.Println()
	return int(longest)
}
