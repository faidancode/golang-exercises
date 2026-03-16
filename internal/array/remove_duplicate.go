package array

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	insertIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
	}

	return insertIndex
}
