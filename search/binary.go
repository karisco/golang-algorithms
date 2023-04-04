package algorithm

import "fmt"

// BinaryDemo 二分查找（Binary Search）是一种常见的查找算法，也称为折半查找
// 它通过将有序数组分成两部分来进行查找，每次比较中间元素与目标值的大小关系，并根据大小关系缩小查找范围
// 直到找到目标元素或者确定目标元素不存在。时间复杂度为 O(log n)
func BinaryDemo(data []int, target int) int {

	left := 0
	right := len(data) - 1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(left, right, mid)
		if data[mid] > target {
			right = mid - 1
		}
		if data[mid] < target {
			left = mid + 1
		}
		if data[mid] == target {
			return mid
		}
	}
	return -1
}
