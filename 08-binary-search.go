func search(nums []int, target int) int {
  start := 0
  end := len(nums) - 1

  for start <= end {
    mid := (start + end) >> 1

    if target < nums[mid] {
      end = mid - 1
    } else if target > nums[mid] {
      start = mid + 1
    } else {
      return mid
    }
  }

  return -1
}
