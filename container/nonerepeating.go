package main

import "fmt"

func lengthOfNoneRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	for k, v := range lastOccurred {
		fmt.Printf("%s:%d\n", string(k), v)
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNoneRepeatingSubStr("abcdgfhydhklssaa"))
	fmt.Println(lengthOfNoneRepeatingSubStr("abcdea"))
	fmt.Println(lengthOfNoneRepeatingSubStr(""))
	fmt.Println(lengthOfNoneRepeatingSubStr("a"))
	fmt.Println(lengthOfNoneRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNoneRepeatingSubStr("这里是这里"))
	fmt.Println(lengthOfNoneRepeatingSubStr("老龙恼怒闹老农,老农恼怒闹老龙"))
}
