package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing
//moving one unit north, south, east, or west, respectively. You start at the
//origin (0, 0) on a 2D plane and walk on the path specified by path.
//
// Return true if the path crosses itself at any point, that is, if at any time
//you are on a location you have previously visited. Return false otherwise.
//
//
// Example 1:
//
//
//Input: path = "NES"
//Output: false
//Explanation: Notice that the path doesn't cross any point more than once.
//
//
// Example 2:
//
//
//Input: path = "NESWW"
//Output: true
//Explanation: Notice that the path visits the origin twice.
//
//
// Constraints:
//
//
// 1 <= path.length <= 10⁴
// path[i] is either 'N', 'S', 'E', or 'W'.
//
//
// Related Topics Hash Table String 👍 1369 👎 42

// leetcode submit region begin(Prohibit modification and deletion)
func isPathCrossing(path string) bool {
	hTable := make(map[string]bool)

	curX := 0
	curY := 0
	hTable["0/0"] = true
	for _, p := range path {
		switch p {
		case 'N':
			curX += 1
		case 'S':
			curX -= 1
		case 'E':
			curY += 1
		case 'W':
			curY -= 1
		}
		pos := fmt.Sprintf("%d/%d", curX, curY)
		if _, exists := hTable[pos]; exists {
			return true
		}
		hTable[pos] = true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_PathCrossing(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			name: "1",
			args: args{
				path: "NES",
			},
			expected: false,
		},
		{
			name: "2",
			args: args{
				path: "NESWW",
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, isPathCrossing(test.args.path))
		})
	}

}
