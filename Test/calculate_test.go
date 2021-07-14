package Test

import (
	"Calculator/service"
	"testing"
)

// 单元测试函数
func Test_Calculate(t *testing.T) {
	cases := []struct {
		a      string
		expect int
	}{
		{"5+2", 7},
		{"5-2", 3},
		{"5*2", 10},
		{"5/2", 2},
	}
	for _, c := range cases {
		if r := service.Calculate(service.Convert(c.a)); r != c.expect {
			t.Errorf("test failed")
		}
	}
}
