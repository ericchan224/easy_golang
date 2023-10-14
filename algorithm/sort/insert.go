package sort

func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i - 1] {
			for j := i; j > 0; j-- {
				if arr[j - 1]  < arr[j] {
					break
				}
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}
		}
	}
	return arr
}

func InsertSortOfficial(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		deal := arr[i]
		j := i - 1
		if deal < arr[j] {
			for ; j >= 0 && deal < arr[j]; j-- {
				arr[j + 1] = arr[j]
			}
			arr[j + 1] = deal
		}
	}
	return arr
}

func InsertionSort(arr []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		deal := arr[i]
		j := i - 1
		if deal < arr[j] {
			for ; j >= left && deal < arr[j]; j-- {
				arr[j + 1] = arr[j]
			}
			arr[j + 1] = deal
		}
	}
}