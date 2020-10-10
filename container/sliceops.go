package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("val=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)

	s3 := make([]int, 16, 20)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	s2 = s2[1:]
	printSlice(s2)

	fmt.Println("Popping from back")
	s2 = s2[:len(s2)-1]
	printSlice(s2)
}
