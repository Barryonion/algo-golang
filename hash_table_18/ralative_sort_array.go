package hash_table_18

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// arr2中每个数字在arr1中出现的次数
	counts := make(map[int]int, 0)
	// 先用arr2构建hash表
	for i := 0; i < len(arr2); i++ {
		counts[arr2[i]] = 0
	}
	// 扫描arr1统计arr2中每个数字在arr1中出现的次数
	for i := 0; i < len(arr1); i++ {
		if oldCount, ok := counts[arr1[i]]; ok {
			counts[arr1[i]] = oldCount + 1
		}
	}
	result := make([]int, len(arr1))
	k := 0
	// 将counts的数据按照arr2的顺序输出
	for i := 0; i < len(arr2); i++ {
		count := counts[arr2[i]]
		for j := 0; j < count; j++ {
			result[k+j] = arr2[i]
		}
		k += count
	}
	// 将arr1中未出现在arr2中的数据有序输出到result
	sort.Ints(arr1)
	for i := 0; i < len(arr1); i++ {
		if _, ok := counts[arr1[i]]; !ok {
			result[k] = arr1[i]
			k++
		}
	}
	return result
}
