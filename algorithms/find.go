package algorithms

//查找

func BinarySearch(origin []int, target int) int {
	length := len(origin)
	if length == 0 {
		return -1
	}
	start := 0
	end := length - 1
	for {
		mid := (start + end) / 2
		midValue := origin[mid]
		if midValue > target {
			end = mid - 1
			if start > end {
				return -1
			}
		} else if midValue < target {
			start = mid + 1
			if start > end {
				return -1
			}
		} else {
			return mid
		}
	}
}
