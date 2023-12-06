package leetcode

import (
	"fmt"
	"testing"
)

//You are given two non-empty linked lists representing two non-negative
//integers. The digits are stored in reverse order, and each of their nodes contains a
//single digit. Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the
//number 0 itself.
//
//
// Example 1:
//
//
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.
//
//
// Example 2:
//
//
//Input: l1 = [0], l2 = [0]
//Output: [0]
//
//
// Example 3:
//
//
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]
//
//
//
// Constraints:
//
//
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have
//leading zeros.
//
//
// Related Topics Linked List Math Recursion 👍 29188 👎 5668

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var indicator *ListNode
	head := indicator
	for {
		if l1 == nil && l2 == nil {
			break
		}
		val := 0
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		next := &ListNode{
			val, nil,
		}

		if indicator == nil {
			indicator = next
			head = indicator
			continue
		}

		indicator.Next = next
		indicator = next
	}

	indicator = head

	for {
		if indicator == nil {
			break
		}
		step := indicator.Val / 10
		indicator.Val = indicator.Val % 10
		if step > 0 && indicator.Next == nil {
			indicator.Next = &ListNode{Val: step}
		} else if step > 0 {
			indicator.Next.Val += step
		}
		indicator = indicator.Next
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func GenerateData(data ...int) *ListNode {
	var ret *ListNode
	var ind *ListNode
	for _, d := range data {
		node := &ListNode{d, nil}
		if ind == nil {
			ind = node
			ret = ind
		} else {
			ind.Next = node
		}
		ind = node
	}

	return ret
}

func Test_AddTwoNumbers(t *testing.T) {
	ret := addTwoNumbers(GenerateData(2, 4, 3), GenerateData(5, 6, 4))
	fmt.Println(ret)
}

func Test_AddTwoNumbers_table(t *testing.T) {
	type args struct {
	}

	tests := []struct {
		name     string
		args     args
		expected interface{}
	}{
		{},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}

}
