package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 0 {
		return -1
	}
	if len(list) == 1 {
		if list[0] == key {
			return 0
		}
		return -1
	}
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := low + (high-low)/2
		if list[mid] == key {
			return mid
		}
		if list[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
