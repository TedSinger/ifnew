package main

import "testing"

func TestWget(t *testing.T) {
	tests := []HumanReadableTestCase{
		{"wget -u http://example.com -o output.txt", "", "index.html output.txt", true},
		{"wget -u http://example.com", "", "index.html", true},
		{"wget", "", "", false},
	}
	TestParse(&Wget{}, tests, t)
}
