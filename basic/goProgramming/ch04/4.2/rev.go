package main

import "fmt"

func main()  {
	s1 := []int{1,2,3,4,5,6,7,8,9}
	//s2 := []string{"a","b","c","d","e"}

	s2 := s1[4:]
	s3 := s1[:4]
	reverse(s2)
	reverse(s3)
	fmt.Println(s1)
}

// reverse the slice
func reverse(s []int)  {
	if len(s) == 0 {
		return
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
