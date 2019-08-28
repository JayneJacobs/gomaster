package sliceexamples

import "fmt"

// SliceExamples cuts a slice that containes values and pointers
func SliceExamples() {
	s1 := []int{1, 2, 3}
	s1 = append(s1, 4, 5, 6)
	s2 := []int{7, 8, 9}
	s1 = append(s1, s2...)
	fmt.Println(s1)
	a := s1
	i := 0
	j := cap(a) - 2
	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k++ {
		a[k] = 0 // or zero value of T
	}
	a = a[:len(a)-j+i]
}
