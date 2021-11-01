/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
package leetcode

func isBadVersion(i int) bool {
	if i == 1 {
		return true
	} else if i == 4 {
		return true
	} else {
		return false
	}
}

func firstBadVersion(n int) int {
	first, latest := 1, n

	for first < latest {
		mid := first + ((latest - first) / 2)

		if isBadVersion(mid) {
			// update latest
			latest = mid
		} else {
			// move first
			first = mid + 1
		}
	}
	return first
}
