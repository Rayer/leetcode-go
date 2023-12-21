package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//Given two strings needle and haystack, return the index of the first
//occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//
// Example 1:
//
//
//Input: haystack = "sadbutsad", needle = "sad"
//Output: 0
//Explanation: "sad" occurs at index 0 and 6.
//The first occurrence is at index 0, so we return 0.
//
//
// Example 2:
//
//
//Input: haystack = "leetcode", needle = "leeto"
//Output: -1
//Explanation: "leeto" did not occur in "leetcode", so we return -1.
//
//
//
// Constraints:
//
//
// 1 <= haystack.length, needle.length <= 10⁴
// haystack and needle consist of only lowercase English characters.
//
//
// Related Topics Two Pointers String String Matching 👍 5211 👎 318

// leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_FindTheIndexOfTheFirstOccurrenceInAString(t *testing.T) {
	type args struct {
		haystack string
		neddle   string
	}

	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "1",
			args: args{
				haystack: "leetcode",
				neddle:   "leeto",
			},
			expected: -1,
		},
		{
			name: "2",
			args: args{
				haystack: "sadbutsad",
				neddle:   "sad",
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ret := strStr(test.args.haystack, test.args.neddle); ret != test.expected {
				assert.Equal(t, test.expected, ret)
			}
		})
	}

}
