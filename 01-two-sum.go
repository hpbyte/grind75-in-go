/**
 * Find the difference between the target and the current number
 * Find the current number in the sumMap i.e the difference of previous number
 * If exists, the sum of those numbers equal to the target
 * If not, add the difference to the sumMap with the curent index as the value
 */
func twoSum(nums []int, target int) []int {
  sumMap := make(map[int]int)

  for i, num := range nums {
    _, e := sumMap[num]
    if e {
      return []int{sumMap[num], i}
    }

    diff := target - num
    sumMap[diff] = i
  }

  return []int{}
}
