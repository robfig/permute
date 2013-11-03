package main

import "testing"

// START TEST OMIT
func TestPermute(t *testing.T) {
	tests := map[string][]string{
		"":    {""},
		"a":   {"a"},
		"ab":  {"ab", "ba"},
		"abc": {"abc", "acb", "bac", "bca", "cab", "cba"},
	}
	for k, expected := range tests {
		actual := permute(k)
		if !eq(expected, actual) {
			t.Errorf("(expected) %#v != %#v (actual)", expected, actual)
		}
	}
}

// END TEST OMIT

// START BENCH OMIT
func BenchmarkPermute5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permute("12345")
	}
}

// END BENCH OMIT

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if ai != b[i] {
			return false
		}
	}
	return true
}
