package bs

func Bs(key int, arr []int, lo int, hi int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2
	if key < arr[mid] {
		return Bs(key, arr, lo, mid-1)
	} else if key > arr[mid] {
		return Bs(key, arr, mid+1, hi)
	}
	return mid
}
