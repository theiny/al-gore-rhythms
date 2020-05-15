package duplicates

func removeDuplicates(nums []int) int {
	var current int
	var pointer int

	for i, num := range nums {

		if i == 0 {
			current = num
			pointer = 1
			continue
		}

		if num == current {
			continue
		}

		nums[pointer] = num
		pointer++
		current = num
	}

	return pointer
}
