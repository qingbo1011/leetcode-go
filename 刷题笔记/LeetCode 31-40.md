# 33. 搜索旋转排序数组

[LeetCode 33. 搜索旋转排序数组](https://leetcode.cn/problems/search-in-rotated-sorted-array/)

## 二分查找

参考[zwjason](https://leetcode.cn/u/zwjago/)的[题解](https://leetcode.cn/problems/search-in-rotated-sorted-array/solution/sou-suo-xuan-zhuan-pai-xu-shu-zu-er-fen-92jot/)。

因为原数组是**严格升序**的，因此旋转后的数组可以分为**两部分的严格升序**，借鉴二分法进行查找。

旋转后的数组，分为左半部分与右半部分，这两部分的大小关系如图所示：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20221001203119.png)

还是采用左闭右闭的区间`[left,right]`。

因为原数组在某一点被旋转了，被分为左右两个部分，在这两个部分中数组是连续的。所以使用二分的必要条件就是：**middle和target必须要在同一部分。**

所以在二分处理的时候我们要进行额外的判断：**判断 `middle` 在哪个部分**。

判断`middle` 在哪个部分，只需要比较`nums[middle]`和`nums[left]`即可：

- 若 `nums[middle] >= nums[left]`，则**middle在左半部分**：
  - 若 `nums[middle] < target`，说明 target**只可能在左半部分**，可以使用二分（此时更新`left=middle+1`）
  - 若 `nums[middle] > target`，说明 target**可能在左半部分也可能在右半部分**：
    - target在左半部分（`target >= nums[left]`），target和middle都在左半部分，可以使用二分，更新`right=middle-1`
    - target在右半部分，middle在左半部分而target在右半部分，不能使用二分。只能更新   `left = middle+1`
- 若 `nums[middle] < nums[left]`，则**middle在右半部分**：
  - 若 `nums[middle] < target`，说明 target**可能在左半部分也可能在有半部分**：
    - target在左半部分（`target >= nums[left]`），不能使用二分，只能更新`right=middle-1`
    - target在右半部分，可以使用二分，更新`left=middle+1`
  - 若 `nums[middle] > target`，说明target只可能在右半部分。可以使用二分，更新`right=middle-1`

> 上述分析，只需要像这样画画草图就能搞清楚了。
>
> ![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20221001213800.png)

有了上面的分析，再解决这题就简单多了。

代码如下：

```go
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right { // 开始二分
		middle := left + (right-left)>>1
		if nums[middle] == target {
			return middle
		}
		// 判断middle在左半部分还是右半部分
		if nums[middle] >= nums[left] { // middle在左半部分
			if nums[middle] < target { // target只可能在左半部分，和middle在同一边，可以使用二分
				left = middle + 1
			} else { // target可能在左半部分也可能在右半部分
				if target >= nums[left] { // target在左半部分，和middle在同一边，可以使用二分
					right = middle - 1
				} else { // target在右半部分，和middle不在同一边，不能使用二分，只能更新left
					left = middle + 1
				}
			}
		} else { // middle在右半部分
			if nums[middle] < target { // target可能在左半部分也可能在有半部分
				if target >= nums[left] { // target在左半部分,和middle不在同一边，不能使用二分只能更新right
					right = middle - 1
				} else { // target在右半部分，和middle在同一边，可以使用二分
					left = middle + 1
				}
			} else { // target只可能在右半部分，和middle在同一边，可以使用二分
				right = middle - 1
			}
		}
	}
	return -1
}
```

> 代码在if else上并没有简化，是为了体现分析到的各种情况。



# 34. 在排序数组中查找元素的第一个和最后一个位置

[LeetCode 34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)

## 二分查找

> 最开始，我自己写的代码如下：
>
> ```go
> func searchRange(nums []int, target int) []int {
> 	left, right := 0, len(nums)-1
> 	res := make([]int, 0)
> 	for left <= right {
> 		middle := (left + right) / 2
> 		if nums[middle] == target { // nums[middle]等于target了，那么我们就以middle为中心向左右移动
> 			l, r := middle, middle
> 			for l = middle; l >= 0; l-- { // 向左移动，找到开始位置
> 				if nums[l] != target {
> 					start := l + 1
> 					res = append(res, start)
> 					break
> 				}
> 			}
> 			if l == -1 { // 说明l走到左边的尽头之后发现都是target，所以start为0
> 				res = append(res, 0)
> 			}
> 			for r = middle; r < len(nums); r++ { // 向右移动，找到结束位置
> 				if nums[r] != target {
> 					end := r - 1
> 					res = append(res, end)
> 					break
> 				}
> 			}
> 			if r == len(nums) { // 说明r走到右边的尽头之后发现都是target，所以end为len(nums)-1
> 				res = append(res, len(nums)-1)
> 			}
> 			return res
> 		} else if nums[middle] < target { // 更新left
> 			left = middle + 1
> 		} else { // 更新right
> 			right = right - 1
> 		}
> 	}
> 	return []int{-1, -1}
> }
> ```
>
> 虽然也可以通过，但是不符合题目要求的：你必须设计并实现时间复杂度为 `O(log n)` 的算法解决此问题。

参考：[代码随想录](https://programmercarl.com/0034.%E5%9C%A8%E6%8E%92%E5%BA%8F%E6%95%B0%E7%BB%84%E4%B8%AD%E6%9F%A5%E6%89%BE%E5%85%83%E7%B4%A0%E7%9A%84%E7%AC%AC%E4%B8%80%E4%B8%AA%E5%92%8C%E6%9C%80%E5%90%8E%E4%B8%80%E4%B8%AA%E4%BD%8D%E7%BD%AE.html)

寻找target在数组里的左右边界，有如下**三种情况**：

1. target 在数组范围的右边或者左边，例如数组{3, 4, 5}，target为2，此时应该返回`[-1, -1]`
2. target 在数组范围中，且数组中不存在target，例如数组{3,6,7}，target为5，此时应该返回`[-1, -1]`
3. target 在数组范围中，且数组中存在target，例如数组{3,6,7}，target为6，此时应该返回`[1, 1]`

> 官方题解是使用了一个二分查找函数，通过设置参数（bool类型）来决定是左边界还是右边界。
>
> 但是刚刚接触二分搜索的话不建议直接用一个二分来查找左右边界，很容易把自己绕进去。还是建议扎扎实实的**写两个二分分别找左边界和右边界**。 

- **左边界**：左边最接近target的位置leftBorder 
- **右边界**：右边最接近target的位置rightBorder

那么结果就是`[leftBorder + 1, rightBorder - 1]`（当然要是在情况2的条件下）

这里采用左闭右闭的区间`[left, right]`。（确定好：计算出来的右边界是不包含target的右边界，左边界同理。）

### 寻找右边界

**右边界**：右边最接近target的位置rightBorder

代码如下：

```go
/*
/ 二分查找，寻找target的右边界（不包括target）
// 考虑到rightBorder没有被赋值的情况（即target在数组范围的左边，例如数组[3,3]，target为2），用来处理情况一
*/
func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	rightBorder := -2 // rightBorder为-2，表示rightBorder没有被赋值（情况一）
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] > target { // 更新right（target在左区间，所以更新区间为[left,middle-1]）
			right = middle - 1
		} else if nums[middle] == target { // 找到了，边界为middle+1（但是此时不能break！要继续循环）
			rightBorder = middle + 1
			left = middle + 1
		} else if nums[middle] < target { // 更新left（target在右区间，所以更新区间为[middle+1,right]）
			left = middle + 1
		}
	}
	return rightBorder
}
```

> 这里的if else处理不是最简洁的，为了方便明白各种情况所以分了三类。其实也可以这么写：
>
> ```go
> if nums[middle] > target {
>    right = middle - 1 // target 在左区间，所以[left, middle - 1]
> } else { // 当nums[middle] == target的时候，更新left，这样才能得到target的右边界
>    left = middle + 1
>    rightBorder = left
> }
> ```
>
> 在寻找左边界的时候就会像上面那样处理。

### 寻找左边界

**左边界**：左边最接近target的位置leftBorder 

代码如下：

```go
/*
  二分查找，寻找target的左边界leftBorder（不包括target）
  考虑到leftBorder没有被赋值的情况（即target在数组范围的右边，例如数组[3,3],target为4），用来处理情况一
*/
func getLeft(nums []int, target int) int {
   left, right := 0, len(nums)-1
   leftBorder := -2 // leftBorder为-2，表示leftBorder没有被赋值（情况一）
   for left <= right {
      middle := left + (right-left)>>1
      if nums[middle] < target { // target在有区间，更新区间为[middle+1,right]
         left = middle + 1
      } else {
         right = middle - 1
         leftBorder = right
      }
   }
   return leftBorder
}
```

### 处理三种情况

左右边界计算完之后，就要开始处理这三种情况了：

1. target 在数组范围的右边或者左边，例如数组{3, 4, 5}，target为2，此时应该返回`[-1, -1]`
   - `leftBorder == -2 || rightBorder == -2`
2. target 在数组范围中，且数组中不存在target，例如数组{3,6,7}，target为5，此时应该返回`[-1, -1]`
   - `else`的情况
3. target 在数组范围中，且数组中存在target，例如数组{3,6,7}，target为6，此时应该返回`[1, 1]`
   - `rightBorder-leftBorder >= 2`

判断代码如下：

```go
	leftBorder := getLeft(nums, target)
	rightBorder := getRight(nums, target)
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	} else if rightBorder-leftBorder >= 2 {
		return []int{leftBorder + 1, rightBorder - 1}
	} else {
		return []int{-1, -1}
	}
```

完整代码如下：

```go
func searchRange(nums []int, target int) []int {
	leftBorder := getLeft(nums, target)
	rightBorder := getRight(nums, target)
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	} else if rightBorder-leftBorder >= 2 {
		return []int{leftBorder + 1, rightBorder - 1}
	} else {
		return []int{-1, -1}
	}
}

/*
  二分查找，寻找target的左边界leftBorder（不包括target）
  考虑到leftBorder没有被赋值的情况（即target在数组范围的右边，例如数组[3,3],target为4），用来处理情况一
*/
func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	leftBorder := -2 // leftBorder为-2，表示leftBorder没有被赋值（情况一）
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] < target { // target在有区间，更新区间为[middle+1,right]
			left = middle + 1
		} else {
			right = middle - 1
			leftBorder = right
		}
	}
	return leftBorder
}

/*
  二分查找，寻找target的右边界（不包括target）
  考虑到rightBorder没有被赋值的情况（即target在数组范围的左边，例如数组[3,3]，target为2），用来处理情况一
*/
func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	rightBorder := -2 // rightBorder为-2，表示rightBorder没有被赋值（情况一）
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] > target { // 更新right（target在左区间，所以更新区间为[left,middle-1]）
			right = middle - 1
		} else if nums[middle] == target { // 找到了，边界为middle+1（但是此时不能break！要继续循环）
			rightBorder = middle + 1
			left = middle + 1
		} else if nums[middle] < target { // 更新left（target在右区间，所以更新区间为[middle+1,right]）
			left = middle + 1
		}
		//if nums[middle] > target {
		//	right = middle - 1 // target 在左区间，所以[left, middle - 1]
		//} else { // 当nums[middle] == target的时候，更新left，这样才能得到target的右边界
		//	left = middle + 1
		//	rightBorder = left
		//}
	}
	return rightBorder
}
```

## go sort包

我看[官方题解](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/solution/zai-pai-xu-shu-zu-zhong-cha-zhao-yuan-su-de-di-3-4/)的思路，跟上面二分查找的思路一样。但是使用了sort包下的`SearchInts`方法。

> go sort包的用法，可以参考我之前的[笔记](https://www.qingbo1011.top/2022/04/28/Go%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95/#go-sort%E5%8C%85)。

在[LeetCode 35. 搜索插入位置](https://leetcode.cn/problems/search-insert-position/)的题解中也提到过了，`sort.SearchInts`方法返回的是按自然顺序排序（升序）的切片中target的插入位置（即并返回目标值的索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置）。

代码如下：

```go
func searchRange(nums []int, target int) []int {
	leftBorder := sort.SearchInts(nums, target)
	if leftBorder == len(nums) || nums[leftBorder] != target { // 说明target不存在数组中
		return []int{-1, -1}
	}
	// target存在数组中,开始找rightBorder
	rightBorder := sort.SearchInts(nums, target+1) - 1
	return []int{leftBorder, rightBorder}
}
```



# 35. 搜索插入位置

[LeetCode 35. 搜索插入位置](https://leetcode.cn/problems/search-insert-position/)

## 二分查找

二分查找，选择**左闭右闭**的区间，那么：

- 循环调价：`for left <= right`
- 更新left：`left = middle + 1`
- 更新right：`right = middle - 1`

代码如下：

```go
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target { // 更新left
			left = middle + 1
		} else { // 更新right
			right = middle - 1
		}
	}
	return left
}
```

> Go sort包：
>
> ```go
> func searchInsert(nums []int, target int) int {
>    return sort.SearchInts(nums, target)
> }
> ```
>
> 这里使用Go sort包是想说明：`sort.SearchInts`方法返回的是按自然顺序排序（升序）的切片中target的插入位置。













