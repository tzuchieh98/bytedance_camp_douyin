package util

import (
	"fmt"
	"testing"
)

func TestGetIntersection(t *testing.T) {
	s1 := []int64{1, 2, 3, 4, 5}
	s2 := []int64{2, 3, 4, 5, 6}
	ans := GetIntersection(s1, s2)
	fmt.Println(ans)
}
