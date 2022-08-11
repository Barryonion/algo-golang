package hash_table_18

func hasCycle(head *ListNode) bool {
	//go 里没有hashset，用 map 代替
	hashTable := make(map[*ListNode]bool)
	p := head
	for p != nil {
		if hashTable[p] {
			return true
		} else {
			hashTable[p] = true
		}
		p = p.Next
	}
	return false
}
