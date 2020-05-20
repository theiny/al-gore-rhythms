package remove

func removeElement(nums []int, val int) int {
	var cursor int

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}

		nums[cursor] = nums[i]
		cursor++
	}

	return cursor
}
