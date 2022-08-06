package hash_table_18

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	nums1 := make([]int, 26)
	for i := 0; i < len(s); i++ {
		c := s[i]
		nums1[c-'a']++
	}
	nums2 := make([]int, 26)
	for i := 0; i < len(t); i++ {
		c := t[i]
		nums2[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
