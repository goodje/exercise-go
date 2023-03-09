package dp

import "testing"

type TestLongestCommonSubstrCase struct {
	text1 string
	text2 string
	want  int
}

func TestLongestCommonSubstr(t *testing.T) {
	cases := []TestLongestCommonSubstrCase{
		TestLongestCommonSubstrCase{"a", "b", 0},
		TestLongestCommonSubstrCase{"a", "abc", 1},
		TestLongestCommonSubstrCase{"abc", "abc", 3},
		TestLongestCommonSubstrCase{"abcd", "abc", 3},
		TestLongestCommonSubstrCase{"abc", "abcd", 3},
		TestLongestCommonSubstrCase{"abcabcd", "abcd", 4},
		TestLongestCommonSubstrCase{"abcdabck", "abckabcd", 4},
	}
	for _, c := range cases {
		got := longestCommonSubstr(c.text1, c.text2)
		if got != c.want {
			t.Errorf("got(%d) != want(%d), test case: text1=%s, text2=%s", got, c.want, c.text1, c.text2)
		}
	}
}
