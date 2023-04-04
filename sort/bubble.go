package algorithm

// BubbleDemo
// 冒泡排序（Bubble Sort）是一种简单的排序算法，它依次比较相邻两个元素的大小关系
// 如果顺序不对则交换位置。通过多次遍历，每次遍历将未排好序的最大元素移动到已排好序的部分的末尾
// 直到所有元素都排好序。时间复杂度为 O(n^2)
// 虽然冒泡排序时间复杂度较高 但它实现简单易懂，适用于小规模数据的排序。
func BubbleDemo(data *[10]int) {
	lens := len(data)
	for i := 0; i < lens-1; i++ {
		for j := 0; j < lens-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// ShakerBubble
// 鸡尾酒排序，又称为定向冒泡排序，是对基本冒泡排序算法的一种改进。
// 鸡尾酒排序也是从左到右依次比较相邻的两个元素，然后从右到左依次比较相邻的两个元素
// 在排序的过程中，元素像鸡尾酒一样不断地往左和往右来回摇晃，因此也被称为“摇摆排序”
// 在鸡尾酒排序中，每一轮排序都通过从左到右和从右到左两个方向分别找到当前未排序部分的最大值和最小值，并将它们放到正确的位置上
// 在一些情况下，这种方法可以更快地达到排序的目的，例如当待排序序列中的大部分元素已经有序时，鸡尾酒排序可以减少比较和交换的次数。
func ShakerBubble(data *[10]int) {
	leng := len(data)
	left, right := 0, leng-1
	for left < right {
		sorting := false
		for i := left; i < right; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				sorting = true
			}
		}
		right--
		for i := right; i > left; i-- {
			if data[i] < data[i-1] {
				data[i], data[i-1] = data[i-1], data[i]
				sorting = true
			}
		}
		left++
		if !sorting {
			return
		}
	}
}

// OddEven 它是一种并行算法，与鸡尾酒排序类似，但是只能用于具有某些特殊性质的序列
// 例如序列中的元素必须是可比较的。奇偶排序将序列拆分成奇数位和偶数位两部分，然后在这两部分内分别进行比较和交换，最终将序列排序。
func OddEven(data *[10]int) {
	leng := len(data)
	soring := true
	for soring {
		soring = false
		for i := 0; i < leng-1; i += 2 {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				soring = true
			}
		}
		for i := 1; i < leng-1; i += 2 {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				soring = true
			}
		}
	}
}

// QuickBubble 这个快速冒泡排序算法与普通冒泡排序类似，但是增加了一个检查 sorted 变量的步骤来判断是否需要继续排序
// 在第一轮中，从左到右比较和交换相邻的元素，如果发生了交换，则将 sorted 变量置为 false。如果没有发生交换，则说明序列已经排好序，可以提前退出循环
// 在第二轮中，从右到左比较和交换相邻的元素，然后再次检查 sorted 变量，以确定是否需要继续排序。
func QuickBubble(data *[200]int) {
	leng := len(data)
	sorting := true
	for sorting {
		sorting = false
		for i := 0; i < leng-1; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				sorting = true
			}
		}
		for i := len(data) - 2; i >= 0; i-- {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				sorting = true
			}
		}
	}
}

// EnhancedBubble
// 改进的冒泡排序算法记录了最后一次交换的位置，从而缩小了排序范围，减少了比较和交换的次数
// 算法首先遍历整个序列，并在每一轮中比较和交换相邻的元素。如果发生了交换，则将 swapped 变量置为 true
// 在每一轮结束时，检查 swapped 变量，如果为 false，则说明序列已经排好序，可以提前退出循环
// 如果 swapped 变量为 true，则将最后一次交换的位置记录下来，下一轮只需要比较到该位置即可
func EnhancedBubble(data *[10]int) {
	leng := len(data)
	for i := 0; i < leng-1; i++ {
		sorting := false
		for j := 0; j < leng-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				sorting = true
			}
		}
		if !sorting {
			break
		}
	}
}
