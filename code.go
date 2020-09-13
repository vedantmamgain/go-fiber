func findMin(nums []int) int {
	low := 0 high := len(nums) - 1
	for low <= high{
		median := low + (high-low)/2
		if median > 0  && nums[median] < nums[median-1]{
			return nums[median]
		}else if nums[median] > nums[high] {
			low = median + 1
		}else{
			high = median - 1
		}
	}
	
		return nums[low]}