package algorithm

// InsertionDemo
// 插入排序（Insertion Sort）是一种简单直观的排序算法，它类似于打扑克牌时的排序过程
// 将待排序元素分为已排序和未排序两部分，每次从未排序部分取出一个元素，与已排序部分从后往前依次比较大小
// 找到合适的位置插入该元素，直到所有元素都被插入到已排序部分中。时间复杂度为 O(n^2)
// 但在对近乎有序的数据进行排序时，插入排序效率较高。
func InsertionDemo(data []int) {
	leng := len(data)
	for i := 1; i < leng; i++ {
		val := data[i]
		j := i - 1
		for j >= 0 && data[j] > val {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = val
	}
}
