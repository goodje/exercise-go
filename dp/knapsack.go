package dp

import "sort"

func knapsack(capacity int, weights, values []int) int {
	if capacity < 1 || len(weights) < 1 {
		return 0
	}

	// sort by weight
	wv := make([][2]int, len(weights))
	for i, w := range weights {
		wv[i] = [2]int{w, values[i]}
	}
	sort.Sort(WVSlice(wv))

	dp := make([][]int, len(weights))

	for i := 0; i < len(weights); i++ {
		dp[i] = make([]int, capacity+1)
		for c := 0; c <= capacity; c++ { // TODO c might be able to be initialized b with min(weight)
			pOld := 0
			pOldMWi := 0
			if i != 0 {
				pOld = dp[i-1][c]
				if c-wv[i][0] > 0 {
					pOldMWi = dp[i-1][c-wv[i][0]]
				}
			}
			dp[i][c] = max(pOld, pOldMWi+wv[i][1])
		}
	}
	return dp[len(weights)-1][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type WVSlice [][2]int

func (x WVSlice) Len() int           { return len(x) }
func (x WVSlice) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x WVSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
