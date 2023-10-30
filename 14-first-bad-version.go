/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
  start := 0
  end := n

  for start <= end {
    mid := (start + end) >> 1

    if (isBadVersion(mid)) {
      end = mid - 1
    } else {
      start = mid + 1
    }
  }

  return start
}
