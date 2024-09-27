package main

import "testing"

func TestCurl(t *testing.T) {
	tests := []TestCase{
		{[]string{"-u", "http://example.com", "-o", "output.txt"}, []string{}, []string{"output.txt"}, true},
		{[]string{"-u", "http://example.com"}, []string{}, []string{}, true},
		{[]string{}, nil, nil, false},
	}
	TestParse(&Curl{}, tests, t)
}

func TestCurlValidateTestCase(t *testing.T) {
	tests := []TestCase{
		{[]string{"-u", "http://example.com", "-o", "output.txt"}, []string{}, []string{"output.txt"}, true},
		{[]string{"-u", "http://example.com"}, []string{}, []string{}, true},
	}
	TestValidateTestCase(&Curl{}, tests, t)
}
