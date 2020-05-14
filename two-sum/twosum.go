package twosum

func twoSum(nums []int, target int) []int {

	ref := make(map[int]int)

	for i, num := range nums {
		if v, ok := ref[target-num]; ok {
			return []int{i, v}
		}

		ref[num] = i
	}

	return nil
}
