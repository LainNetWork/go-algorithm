package algorithms

//时间复杂度
//一个有序数组A，长度为M，另一个无序数组B,长度为N，打印B中所有不在A中的数，A数组长度为N，B数组长度为M。

// Complexity1 方法一：遍历所有B数组中的数，将B数组中的数与A数组一一比对
// 时间复杂度为
func Complexity1(A []int, B []int) []int {
	results := make([]int, 0)
	for _, item := range B {
		var flag = true
		for _, itemA := range A {
			if itemA == item {
				flag = false
			}
		}
		if flag {
			results = append(results, item)
		}
	}
	return results
}

// Complexity2 方法2：循环遍历B，在A中使用二分法进行查找,没有找到则加入结果数组
func Complexity2(A []int, B []int) []int {
	results := make([]int, 0)
	for _, item := range B {
		if BinarySearch(A, item) == -1 {
			results = append(results, item)
		}
	}
	return results
}
