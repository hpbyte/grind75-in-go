/**
 * Naive approach in JS:
 * - sort the two strings and check if they're the same
 *
 * Approach 2:
 * - check if the two strings length are the same, if not, they're not valid
 * - have an array to keep track of frequencies of the 26 characters
 * - loop thru the string
 * - increment the character occurences from `s` string
 * - decrement the character occurences from `t` string
 * - eventually, if it's valid, the frequenceis should cancel out i.e 0s in freq array
 */
func isAnagram(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }

  var freq [26]int

  for i := range s {
    freq[s[i] - 97] += 1
    freq[t[i] - 97] -= 1
  }

  for _, f := range freq {
    if f != 0 {
      return false
    }
  }

  return true
}
