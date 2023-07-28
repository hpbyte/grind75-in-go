/**
 * make a stack to keep the open brackets
 * loop over the string
 * if the char is one of the open brackets i.e { [ (
 * append its corresponding closed bracket to the list
 * if not, it's a closed bracket
 * since it must be in order,
 * check the top from the stack if it's equal to the current char
 * if it doesn't, then it's invalid
 * if after the loop is over and the stack is empty
 * it's a valid one
 */
func isValid(str string) bool {
  var stack []string
  m := map[string]string{"{": "}", "(": ")", "[": "]"}

  for _, s := range str {
    char := string(s)
    _, isOpen := m[char]

    if isOpen {
      stack = append(stack, m[char])
    } else {
      top := len(stack) - 1

      if top < 0 || stack[top] != char {
        return false
      }

      // popping the top
      stack = stack[:top]
    }
  }

  return len(stack) == 0
}
