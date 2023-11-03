func longestPalindrome(s string) int {
    if len(s) == 1 {
        return 1
    }

    charCount := make(map[rune]int)
    oddCount := 0

    for _, char := range s {
        charCount[char]++
        if charCount[char]%2 == 1 {
            oddCount++
        } else {
            oddCount--
        }
    }

    if oddCount > 1 {
        return len(s) - oddCount + 1
    }

    return len(s)
}                                                                                                                      }

