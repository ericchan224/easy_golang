package sort

/**
	计数排序，了解即可
	只适用于基础数据类型的使用，且所传的数据在一定范围内
	这里假设数据范围是-5000 ~ 5000
**/
const size = 100000
const offset = 50000

func CountSort(nums []int) {
	count := make([]int, size)

	// 为了保证数据都大于0，这里加上偏移量
	for _, num := range nums {
		count[num + offset]++
	}

	// 将count数组转换为前缀和数组，定位每个元素
	for i := 1; i < size; i++ {
		count[i] += count[i - 1]
	}

	// copy出一个临时数组，用于备份
	temp := make([]int, len(nums))
	copy(temp, nums)

	// 从后往前放置元素位置到nums里，完成排序
	for i := len(nums) - 1; i >= 0; i-- {
		index := count[temp[i] + offset] - 1
		nums[index] = temp[i]
		count[temp[i] + offset]--
	}
}