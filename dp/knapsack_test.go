package dp

import "testing"

type DPCase struct {
	c    int
	w    []int
	v    []int
	want int
}

func TestDP(t *testing.T) {

	cases := []DPCase{
		DPCase{2, []int{}, []int{}, 0},
		DPCase{1, []int{1, 2, 3}, []int{5, 4, 2}, 5},
		DPCase{6, []int{1, 2, 3}, []int{5, 4, 2}, 11},
		DPCase{5, []int{5, 2, 3}, []int{2, 5, 3}, 8},
		DPCase{5, []int{5, 2, 11, 3}, []int{2, 5, 3, 7}, 12},
	}

	for _, c := range cases {
		got := knapsack(c.c, c.w, c.v)
		if got != c.want {
			t.Errorf("got(%d) != want(%d), test case: capacity=%d, weights=%v, values=%v", got, c.want, c.c, c.w, c.v)
		}
	}

}
