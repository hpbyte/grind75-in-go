/**
 * classic Dynamic Programming problem
 *
 * the base cases (0 and 1 steps) are initialized to 1 since there's only one way to climb
 * so the iteration starts from 2
 * from step 2, track all the sums of the previous two steps until step n
 */
func climbStairs(n int) int {
  if n == 0 || n == 1 {
    return 1
  }

  dp := map[int]int{ 0: 1, 1: 1 }

  for i := 2; i <= n; i++ {
    dp[i] = dp[i - 1] + dp[i - 2]
  }

  return dp[n]
}
