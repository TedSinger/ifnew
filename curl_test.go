package main

import "testing"

func TestCurl(t *testing.T) {
	tests := []HumanReadableTestCase{
		{"curl http://example.com -o output.txt", "", "output.txt", true},
		{"curl http://example.com", "", "", true},
		{"curl", "", "", false},
	}
	TestParse(&Curl{}, tests, t)
}

func TestCurlValidateTestCase(t *testing.T) {
	tests := []HumanReadableTestCase{
		{"curl http://example.com -o output.txt", "", "output.txt", true},
		{"curl http://example.com", "", "", true},
		{"curl", "", "", false},
	}
	TestValidateTestCase(&Curl{}, tests, t)
}
