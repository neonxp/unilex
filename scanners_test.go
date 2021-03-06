package unilex

import "testing"

func TestScanNumber(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected bool
		Pos      int
	}{
		{"asd", false, 0},
		{"asd123", false, 0},
		{"123", true, 3},
		{"123asd", true, 3},
		{"123.321", true, 7},
	}
	for _, tc := range testCases {
		l := New(tc.Input)
		actual := ScanNumber(l)
		if actual != tc.Expected {
			t.Errorf("Input: %s expected scan result: %v actual: %v", tc.Input, tc.Expected, actual)
		}
		if l.Pos != tc.Pos {
			t.Errorf("Input: %s expected scan position: %d actual: %d", tc.Input, tc.Pos, l.Pos)
		}
	}
}

func TestScanAlphaNum(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected bool
		Pos      int
	}{
		{"asd", true, 3},
		{"asd123", true, 6},
		{"123", false, 0},
		{"123asd", false, 0},
		{"123.321", false, 0},
		{"asd!dsa", true, 3},
		{"asd dsa", true, 3},
	}
	for _, tc := range testCases {
		l := New(tc.Input)
		actual := ScanAlphaNum(l)
		if actual != tc.Expected {
			t.Errorf("Input: %s expected scan result: %v actual: %v", tc.Input, tc.Expected, actual)
		}
		if l.Pos != tc.Pos {
			t.Errorf("Input: %s expected scan position: %d actual: %d", tc.Input, tc.Pos, l.Pos)
		}
	}
}

func TestScanQuotedString(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected bool
		Pos      int
	}{
		{`asd`, false, 0},
		{`"asd`, false, 0},
		{`"asd"qwe`, true, 5},
	}
	for _, tc := range testCases {
		l := New(tc.Input)
		actual := ScanQuotedString(l, '"')
		if actual != tc.Expected {
			t.Errorf("Input: %s expected scan result: %v actual: %v", tc.Input, tc.Expected, actual)
		}
		if l.Pos != tc.Pos {
			t.Errorf("Input: %s expected scan position: %d actual: %d", tc.Input, tc.Pos, l.Pos)
		}
	}
}
