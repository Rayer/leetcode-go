package leetcode

import (
	"testing"
)

//Hercy wants to save money for his first car. He puts money in the Leetcode
//bank every day.
//
// He starts by putting in $1 on Monday, the first day. Every day from Tuesday
//to Sunday, he will put in $1 more than the day before. On every subsequent
//Monday, he will put in $1 more than the previous Monday.
//
// Given n, return the total amount of money he will have in the Leetcode bank
//at the end of the nᵗʰ day.
//
//
// Example 1:
//
//
//Input: n = 4
//Output: 10
//Explanation: After the 4ᵗʰ day, the total is 1 + 2 + 3 + 4 = 10.
//
//
// Example 2:
//
//
//Input: n = 10
//Output: 37
//Explanation: After the 10ᵗʰ day, the total is (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2
// + 3 + 4) = 37. Notice that on the 2ⁿᵈ Monday, Hercy only puts in $2.
//
//
// Example 3:
//
//
//Input: n = 20
//Output: 96
//Explanation: After the 20ᵗʰ day, the total is (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2
// + 3 + 4 + 5 + 6 + 7 + 8) + (3 + 4 + 5 + 6 + 7 + 8) = 96.
//
//
//
// Constraints:
//
//
// 1 <= n <= 1000
//
//
// Related Topics Math 👍 1307 👎 40

// leetcode submit region begin(Prohibit modification and deletion)
func totalMoney(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += (i % 7) + 1
		sum += i / 7
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_CalculateMoneyInLeetcodeBank(t *testing.T) {

}

func Test_CalculateMoneyInLeetcodeBank_table(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "1",
			args: args{
				n: 4,
			},
			expected: 10,
		},
		{
			name: "2",
			args: args{
				n: 10,
			},
			expected: 37,
		},
		{
			name: "3",
			args: args{
				n: 20,
			},
			expected: 96,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := totalMoney(test.args.n); res != test.expected {
				t.Errorf("expected %v but get %v", test.expected, res)
			}
		})
	}

}
