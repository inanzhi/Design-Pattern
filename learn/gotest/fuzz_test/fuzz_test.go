package fuzz_test

import (
	"github.com/inanzhi/factualcase/learn/gotest/fuzz"
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{in: "Hello, world", want: "dlrow ,olleH"},
		{in: " ", want: " "},
		{in: "!12345", want: "54321!"},
	}

	for _, tc := range testcases {
		rev, err := fuzz.Reverse(tc.in)
		if err != nil {
			t.Errorf("Unexpected error:%v", err)
		}
		if rev != tc.want {
			t.Errorf("Reverse:%q,want %q", rev, tc.want)
		}
	}
}

// go test . -fuzz=Fuzz
// go test . -fuzz=Fuzz -fuzztime=5s 限制测试运行时间
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) //添加测试种子
	}
	f.Fuzz(func(t *testing.T, orig string) {
		if !utf8.ValidString(orig) { //忽略构造的无效字符（非UTF-8编码字符串
			return
		}
		rev, err := fuzz.Reverse(orig)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !utf8.ValidString(rev) { //忽略构造的无效字符（非UTF-8编码字符串
			t.Fatalf("Reverse produced invalid UTF-8 string %q", rev)
		}

		doubleRev, err := fuzz.Reverse(rev)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if orig != doubleRev {
			t.Errorf("Before:%q, after: %q", orig, doubleRev)
		}

	})
}
