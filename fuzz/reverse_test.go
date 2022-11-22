package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, World", "dlroW ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testcases {
		reverse := Reverse(tc.in)
		if reverse != tc.want {
			t.Errorf("Reverse:%q, want %q", reverse, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", "", "!12345"}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)

		if orig != doubleRev {
			t.Errorf("Before : %q, After: %q", orig, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
