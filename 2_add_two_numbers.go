/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * Runtime: 3 ms, faster than 99.67% of Go online submissions for Add Two Numbers.
 * Memory Usage: 4.6 MB, less than 46.23% of Go online submissions for Add Two Numbers.
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}
	result := &ListNode{0, nil}

	val1, val2 := 0, 0
	if l1 != nil {
		val1 = l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		val2 = l2.Val
		l2 = l2.Next
	}

	result.Val = (val1 + val2) % 10
	if (val1 + val2) >= 10 {

		if l1 != nil {
			l1.Val += 1
		} else if l2 != nil {
			l2.Val += 1
		} else {
			result.Next = &ListNode{1, nil}
			return result
		}

	}

	result.Next = addTwoNumbers(l1, l2)
	return result

}