package monotonicstack

func CalculateGreaterElement(nums []int) []int {
	stack := []int{}
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			result[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
	}

	for len(stack) > 0 {
		result[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	return result
}