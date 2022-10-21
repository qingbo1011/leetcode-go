package main

import "sort"

func main() {

}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for k := 0; k < len(nums); k++ {
		if target >= 0 && nums[k] > target { // 后面不可能再出现结果集了（注意这里要有target >= 0的条件）
			break
		}
		if k > 0 && nums[k] == nums[k-1] { // 保证nums[k]不会在结果集中重复
			continue
		}
		for i := k + 1; i < len(nums); i++ {
			if target >= 0 && nums[k]+nums[i] > target { // 后面不可能再出现结果集了（注意这里要有target >= 0的条件）
				break
			}
			if i > k+1 && nums[i] == nums[i-1] { // 保证nums[k]不会在结果集中重复
				continue
			}
			// 这里的逻辑就跟三数之和差不多了
			left, right := i+1, len(nums)-1
			for left < right {
				if nums[k]+nums[i]+nums[left]+nums[right] < target { // 更新left
					left++
				} else if nums[k]+nums[i]+nums[left]+nums[right] > target { // 更新right
					right--
				} else { // 找到结果集了
					res = append(res, []int{nums[k], nums[i], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] { // 对nums[left]进行去重
						left++
					}
				}
			}
		}
	}

	return res

}
