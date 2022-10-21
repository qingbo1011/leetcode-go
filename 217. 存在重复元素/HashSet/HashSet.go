package main

func main() {

}

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{}, 0)
	for _, num := range nums {
		if _, ok := set[num]; ok { // 出现了重复数据
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
