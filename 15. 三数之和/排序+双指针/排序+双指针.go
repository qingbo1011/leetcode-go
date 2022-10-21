package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))

}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums) // 先排序
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 { // 后面再也不可能出现三数之和为0的三元组了
			return res
		}
		if i > 0 && nums[i] == nums[i-1] { // 保证nums[i]不会在结果集中重复
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 { // 需要更新left
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 { // 需要更新right
				right--
			} else { // 找到结果集了
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] { // 对nums[left]进行去重
					left++
				}
				for left < right && nums[right] == nums[right-1] { // 对nums[right]进行去重
					right--
				}
				// 去重后收缩left和right，让for left < right循环可以找到出口
				left++
				right--
			}
		}
	}
	return res
}
