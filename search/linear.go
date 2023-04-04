package algorithm

// LinearDemo
// 线性查找（Linear Search）是一种简单的查找算法，也称为顺序查找
// 它从数组的起始位置开始依次遍历数组中的每个元素，直到找到目标元素或者遍历完整个数组
// 时间复杂度为 O(n)，其中 n 为数组的长度。虽然线性查找的时间复杂度较高
// 但是它适用于小规模数据的查找，并且不要求数据有序。
func LinearDemo(data []int, target int) int {
	for key, value := range data {
		if value == target {
			return key
		}
	}
	return -1
}
