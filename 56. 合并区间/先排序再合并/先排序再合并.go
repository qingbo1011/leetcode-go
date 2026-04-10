package main

import "sort"

func merge(intervals [][]int) [][]int {
	// 按起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间的起始位置start≤ 上一个合并区间的结束位置lastEnd，说明重叠，更新上一个区间的结束位置为 max(lastEnd, end)
		if intervals[i][0] <= merged[len(merged)-1][1] {
			// 有重叠，更新merged的lastEnd
			merged[len(merged)-1][1] = max(intervals[i][1], merged[len(merged)-1][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

func main() {

}
