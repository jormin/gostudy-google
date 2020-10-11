package main

import (
	"testing"
)

func Test_lengthOfNoneRepeatingSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcdgfhydhklssaa", 8},
		{"abcdea", 5},
		{"", 0},
		{"a", 1},
		{"这里是慕课网", 6},
		{"这里是这里", 3},
		{"老龙恼怒闹老农,老农恼怒闹老龙", 7},
	}
	for _, tt := range tests {
		actual := lengthOfNoneRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

// 性能测试
func Benchmark_lengthOfNoneRepeatingSubStr(b *testing.B) {
	s := "老龙恼怒闹老农,老农恼怒闹老龙"
	ans := 7

	for i := 0; i < b.N; i++ {
		actual := lengthOfNoneRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
