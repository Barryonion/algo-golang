package hash_table_18

import "testing"

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(groupAnagrams(strs))
}
