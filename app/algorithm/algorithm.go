package algorithm

import "fmt"

// SelectSort 选择排序
// 基本思想： 选择排序的原理是，对给定的数组进行多次遍历，每次均找出最大的一个值的索引。
func SelectSort(data []int) []int {
	// 只有一个不用比较
	length := len(data)
	if length <= 1 {
		return data
	}
	// 对给定的数组进行多次遍历
	for i := 0; i < len(data); i++ {
		// 先假设第一元素索引最小
		k := i
		// 遍历找出最小的
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[k] {
				// 找到更小的，则替换索引值
				k = j
			}
		}
		// 替换元素位置
		if k != i {
			data[i], data[k] = data[k], data[i]
		}
	}
	return data
}

// BubbleSort 冒泡排序
// 对给定的数组进行多次遍历，每次均比较相邻的两个数，如果前一个比后一个大，则交换这两个数。
// 经过第一次遍历之后，最大的数就在最右侧了；第二次遍历之后，第二大的数就在右数第二个位置了；以此类推。
func BubbleSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	// 多次遍历
	for i := 0; i < len(data); i++ {
		// 两两比较
		for j := 1; j < len(data); j++ {
			// 前面大于后面，则替换位置
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return data
}

// QuickSort 快速排序
// 根据找到的基准值，把待排序的数组‘平均’分成两组，即大于基准值的为一组，小于基准值的为一组。
// 然后对两个组再次按照上述方法进行操作，最后合并结果。
func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	// 定义基准值
	benchMark := data[0]
	// 小于基准值的组
	leftGroup := []int{}
	// 大于基准值的组
	rightGroup := []int{}
	// 遍历分组
	for i := 1; i < len(data); i++ {
		// 假设第一个元素
		if data[i] > benchMark {
			rightGroup = append(rightGroup, data[i])
		} else {
			leftGroup = append(leftGroup, data[i])
		}
	}
	// 分组后继续操作
	rightGroup = QuickSort(rightGroup)
	leftGroup = QuickSort(leftGroup)
	// 合并结果
	leftGroup = append(leftGroup, benchMark)
	for i := 0; i < len(rightGroup); i++ {
		leftGroup = append(leftGroup, rightGroup[i])
	}
	return leftGroup
}

// InsertSort 插入排序
// 从第二个数开始向右侧遍历,如果左侧的元素比取的数大，则互换位置
func InsertSort(data []int) []int {
	fmt.Println("length=", len(data))
	if len(data) <= 1 {
		return data
	}
	for i := 1; i < len(data); i++ {
		value := data[i]
		// 左侧数据index
		j := i - 1
		for j >= 0 && data[j] > value {
			// 如果左侧的元素比取的数大，则右移
			data[j+1] = data[j]
			data[j] = value
			// 依次和取的数进行比较
			j--
		}
	}
	return data
}

// BinarySearch 二分查找，返回索引值,不存在则返回-1
func BinarySearch(data []int, beginIndex, endIndex, search int) int {
	// 大前提条件
	if beginIndex <= endIndex {
		// 计算中间下标
		middenIndex := (beginIndex + endIndex) / 2
		if data[middenIndex] == search {
			return middenIndex
		} else if search < data[middenIndex] {
			// 如果查找的值，小于中间值
			return BinarySearch(data, beginIndex, middenIndex-1, search)
		} else {
			// 如果查找的值，大于中间值
			return BinarySearch(data, middenIndex+1, endIndex, search)
		}
	}
	return -1
}

// BinarySearch2 二分查找
func BinarySearch2(data []int, search int) int {
	beginIndex := 0
	endIndex := len(data) - 1
	for beginIndex <= endIndex {
		midIndex := (endIndex + beginIndex) / 2
		if data[midIndex] == search {
			return midIndex
		} else if search < data[midIndex] {
			endIndex = midIndex - 1
		} else {
			beginIndex = midIndex + 1
		}
	}

	return -1
}
