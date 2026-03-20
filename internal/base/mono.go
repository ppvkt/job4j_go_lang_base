package base

func Mono(nums []int) bool {
	sliceSize := len(nums)

	if sliceSize < 3 {
		return true
	}

	isIncreasing, isDecreasing := true, true

	for i := 0; i < len(nums)-1; i++ {
		isIncreasing = isIncreasing && nums[i] >= nums[i+1]
		isDecreasing = isDecreasing && nums[i] <= nums[i+1]
	}

	return isIncreasing || isDecreasing
}
