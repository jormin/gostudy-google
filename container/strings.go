package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "yes我爱慕课网！"
	fmt.Printf("%s\n", []byte(s))
	fmt.Printf("%X\n", []byte(s))
	for _, v := range []byte(s) {
		fmt.Printf("%X ", v)
	}

	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d,%X)", i, ch)
	}

	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(utf8.DecodeRune([]byte(s)))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("(%c,%d)", ch, size)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
	}
	fmt.Println()

	s1 := "老龙恼怒    闹老 农,老 农恼怒  闹老龙"
	fmt.Println(strings.Fields(s1))
	fmt.Println(strings.Split(s1, " "))
	fmt.Println(strings.Join(strings.Split(s1, ""), ""))
}
