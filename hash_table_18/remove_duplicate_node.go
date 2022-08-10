package hash_table_18

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	set := make(map[int]bool)
	newHead := &ListNode{}
	tail := newHead
	p := head
	for p != nil {
		tmp := p.Next
		if !set[p.Val] {
			set[p.Val] = true
			tail.Next = p
			tail = p
			tail.Next = nil
		}
		p = tmp
	}
	return newHead.Next
}
