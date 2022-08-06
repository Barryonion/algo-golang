package hash_table_18

import "sort"

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	for _, str := range strs {
		array := []byte(str)
		sort.Slice(array, func(i, j int) bool {
			return array[i] < array[j]
		})
		key := string(array)
		list := make([]string, 0)
		if _, ok := groupMap[key]; ok {
			list = groupMap[key]
		}
		list = append(list, str)
		groupMap[key] = list
	}
	//没有现成的取map里values的方法，这里一个一个获取
	result := make([][]string, 0)
	for _, value := range groupMap {
		result = append(result, value)
	}
	return result
}
