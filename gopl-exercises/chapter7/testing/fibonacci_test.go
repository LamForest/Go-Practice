package fibonacci

import "testing"

// func TestFib(t *testing.T) {
// 	var (
// 		in       = 7
// 		expected = 13
// 	)
// 	actual := Fib(in)
// 	if actual != expected {
// 		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
// 	}
// }

func TestFib(t *testing.T) { //t用于管理测试状态并支持格式化测试日志
	var fibTests = []struct {
		in       int // input
		expected int // expected result
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{11, 89},
		{14, 377},
		{20, 9999},
	}

	for _, tt := range fibTests {
		actual := Fib(tt.in)
		if actual != tt.expected {
			t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}
