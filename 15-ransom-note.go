/**
 * Have an array of 26 characters default to 0
 * Loop thru the magazine and increment the count of the frequency of each character
 * Loop thru the ransomNote and decrement the counter
 * If there's a character whose count is 0,
 * that means, there's no occurence of that character in the magazine
 * So we return false
 * After the loop ends without returning false
 * We can say that it's constructible
 */
func canConstruct(ransomNote string, magazine string) bool {
  if len(ransomNote) > len(magazine) {
    return false
  }

  var freq [26]int

  for i := range magazine {
    freq[magazine[i] - 97] += 1
  }

  for j := range ransomNote {
    if (freq[ransomNote[j] - 97] == 0) {
      return false
    }

    freq[ransomNote[j] - 97] -= 1
  }

  return true

