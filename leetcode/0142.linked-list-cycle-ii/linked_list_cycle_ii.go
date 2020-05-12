package leetcode

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycle(head *ListNode) *ListNode {
	p1, p2 := head, head

	for {
		if p2 == nil {
			return nil
		}

		p1, p2 = p1.Next, p2.Next
		if p2 == nil {
			return nil
		}

		p2 = p2.Next
		if p1 == p2 {
			p1 = head
			for {
				if p1 == p2 {
					return p1
				}
				p1, p2 = p1.Next, p2.Next
			}
		}
	}
}
