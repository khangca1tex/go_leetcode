func twoSum(nums []int, target int) []int {

	flag := 0

	for flag < len(nums)-1 {

		numNeedFind := target - nums[flag]

		sliceArray := nums[(flag + 1):len(nums)]

		indexMatched := findEleInArray(sliceArray, numNeedFind)

		if indexMatched >= 0 {
			return []int{flag, indexMatched + flag + 1}
		}

		flag++
	}

	return nil
}

// return the index of slice array
func findEleInArray(nums []int, target int) int {

	flag := 0
	for flag < len(nums) {

		if nums[flag] == target {
			return flag
		}

		flag++
	}

	return -1
}