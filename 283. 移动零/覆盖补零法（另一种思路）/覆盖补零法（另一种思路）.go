package main

func moveZeroes(nums []int) {
	// 先遍历数组，将所有非零元素按顺序覆盖到数组前面，记录非零元素的个数 index
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	// 然后从 index 开始到数组末尾，全部赋值为零
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {

}
