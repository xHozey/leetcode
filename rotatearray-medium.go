package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
