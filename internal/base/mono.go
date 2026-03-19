package base

func Mono(nums []int) bool {
	sliceSize := len(nums)

	if sliceSize <= 1 {
		return true
	}

	isIncreasing := true
	isDecreasing := true

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			isIncreasing = false
		}
		if nums[i] < nums[i+1] {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
