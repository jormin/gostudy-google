package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//fmt.Println("arr[2:6]: ", arr[2:6])
	//fmt.Println("arr[:6]: ", arr[:6])
	//fmt.Println("arr[2:]: ", arr[2:])
	//fmt.Println("arr[:]: ", arr[:])
	//
	//s := arr[2:6]
	//s1 := arr[:6]
	//fmt.Println(s)
	//fmt.Println(s1)
	//fmt.Println(arr)
	//updateSlice(s)
	//fmt.Println(s)
	//fmt.Println(s1)
	//fmt.Println(arr)
	//
	//s1 = s1[2:5]
	//fmt.Println(s1)
	//s1 = s1[2:]
	//fmt.Println(s1)
	//s1 = s1[:4]
	//fmt.Println(s1)

	//arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s1 = arr[2:6:6] // [2,3,4,5]
	//s2 = s1[3:5] // 报错
	//fmt.Println(s1, s2)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("[before] arr: %v，len: %d, cap: %d, ptr: %p\n", arr, len(arr), cap(arr), &arr)
	s1 := arr[2:6] // [2,3,4,5]
	s2 := s1[3:5]  // [5,6]
	s3 := append(s2, 10)
	s4 := append(s3, 11, 12, 13)
	s5 := append(s4, 14, 15, 16)
	fmt.Printf("[after] arr: %v，len: %d, cap: %d, ptr: %p\n", arr, len(arr), cap(arr), &arr)
	fmt.Printf("[after] s1: %v，len: %d, cap: %d, ptr: %p\n", s1, len(s1), cap(s1), &s1)
	fmt.Printf("[after] s2: %v，len: %d, cap: %d, ptr: %p\n", s2, len(s2), cap(s2), &s2)
	fmt.Printf("[after] s3: %v，len: %d, cap: %d, ptr: %p\n", s3, len(s3), cap(s3), &s3)
	fmt.Printf("[after] s4: %v，len: %d, cap: %d, ptr: %p\n", s4, len(s4), cap(s4), &s4)
	fmt.Printf("[after] s5: %v，len: %d, cap: %d, ptr: %p\n", s5, len(s5), cap(s5), &s5)

	fmt.Println(&arr[5], &s1[3], &s2[0], &s3[0], &s4[0], &s5[0])
}
