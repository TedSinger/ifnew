package main

import "testing"

func TestWget(t *testing.T) {
	tests := []TestCase{
		{[]string{"-u", "http://example.com", "-o", "output.txt"}, []string{}, []string{"output.txt", "index.html"}, true},
	}
	TestParse(&Wget{}, tests, t)
}
