package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		println(i, v)
	}
}

func printArray2(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		println(i, v)
	}
}

func main() {

	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]int
	fmt.Println(grid)

	for i, v := range arr3 {
		fmt.Printf("下标 %d: %d\n", i, v)
	}

	printArray(arr1)
	printArray(arr3)
	fmt.Println(arr1, arr3)

	fmt.Println("=======")

	printArray2(&arr1)
	printArray2(&arr3)
	fmt.Println(arr1, arr3)

}
