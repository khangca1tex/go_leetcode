func twoSum(nums []int, target int) []int {

	for i, num := range nums {

		numNeedFind := target - num

		sliceArray := nums[(i + 1):len(nums)]

		indexMatched := findEleInArray(sliceArray, numNeedFind)

		if indexMatched >= 0 {
			return []int{i, indexMatched + i + 1}
		}

	}

	return nil
}

// return the index of slice array
func findEleInArray(nums []int, target int) int {

	for i, _ := range nums {

		if nums[i] == target {
			return i
		}

		i++
	}

	return -1
}