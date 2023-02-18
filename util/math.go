package util

func GetIntersection(slice1 []int64, slice2 []int64) []int64 {
	m := make(map[int64]bool, len(slice1))
	for _, value := range slice1 {
		m[value] = true
	}
	ans := make([]int64, 0)
	for _, value := range slice2 {
		if m[value] {
			ans = append(ans, value)
		}
	}
	return ans
}
