package main

import "testing"

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		want string
	}{
		// Checking signed operations.
		{"10", "101", "1010"},
		{"+10", "101", "1010"},
		{"10", "+101", "1010"},
		{"+10", "+101", "1010"},
		{"-10", "101", "-1010"},
		{"10", "-101", "-1010"},
		{"-10", "-101", "1010"},
		{"+10", "-101", "-1010"},
		{"-10", "+101", "-1010"},
		// Resulting zero must never contain sign literals.
		{"0", "0", "0"},
		{"-0", "0", "0"},
		{"0", "-0", "0"},
		{"1", "0", "0"},
		{"1", "+0", "0"},
		{"1", "-0", "0"},
		{"+1", "-0", "0"},
		{"-1", "+0", "0"},

		// Other examples.
		{"1101100100000011", "110111", "1011101001111110100101"},
	} {
		got := LongMul(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("ERR: LongMul(%s, %s): got: %s, want: %s", tc.a, tc.b, got, tc.want)
		}
	}
}

func BenchmarkLongMul(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LongMul("1101100100000011", "110111")
	}
}
