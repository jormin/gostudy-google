package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func normal(a, b int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
	return result, nil
}

func div(a int, b int) (q, r int) {
	return a / b, a % b
}

// 函数体长的话用这个方法不方便
func div2(a int, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	ptr := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(ptr).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for _, v := range numbers {
		//s += numbers[i]
		s += v
	}
	return s
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(normal(3, 4, "*"))
	if result, err := normal(3, 4, "&"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(div(13, 2))
	fmt.Println(div2(13, 2))
	q, _ := div(13, 2)
	fmt.Println(q)

	fmt.Println(apply(pow, 3, 2))

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(
			float64(a),
			float64(b),
		))
	}, 3, 2))

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	a, b := 3, 4
	//swap(&a, &b)
	//fmt.Println(a, b)
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
