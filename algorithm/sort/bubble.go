package sort

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		flag := true
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				flag = false
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
		if flag {
			break
		}
	}
	return arr
}