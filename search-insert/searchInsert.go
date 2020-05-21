package searchinsert

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target || nums[i] > target {
			return i
		}

		if nums[i] == target-1 {
			return i + 1
		}
	}

	return len(nums)
}