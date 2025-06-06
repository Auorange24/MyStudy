package monotonicstack

func CalculateGreaterElement(nums []int) []int {
	stack := []int{}
	result := make([]int, len(nums))
	n := len(nums)

	for i := n - 1 ; i >= 0 ; i -- {
		// 对于nums[i]，要保证单调栈的递增顺序，应该保证栈顶元素大于nums[i]
		// 如果nums[i]小于栈顶元素，则直接弹出栈顶元素
		/*
		对于下一个最大的元素
		*/
		for len(stack) > 0 && stack[len(stack) - 1] <= nums[i] {
			stack = stack[ : len(stack) - 1]
		}
	}
}