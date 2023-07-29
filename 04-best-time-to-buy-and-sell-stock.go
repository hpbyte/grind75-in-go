/**
 * use two pointers approach
 * start with left 0, right 1 and maxProfit 0
 * loop until right pointer exceeds prices length
 * increment the right pointer in each iteration
 *
 * if the left price is less than the right,
 * that means, we get profit
 * so calculate the profit and compare with the current maxProfit
 * if it's greater than current maxProfit, update maxProfit
 * if the left price is not less than the right
 * we update the left pointer pos with the right
 */
func maxProfit(prices []int) int {
  left := 0
  right := 1
  maxProfit := 0

  for right < len(prices) {
    if prices[left] < prices[right] {
      diff := prices[right] - prices[left]

      if diff > maxProfit {
        maxProfit = diff
      }
    } else {
      left = right
    }

    right += 1
  }

  return maxProfit
}
