package main

import "testing"

func TestCp(t *testing.T) {
	tests := []TestCase{
		{[]string{"src.txt", "dst.txt"}, []string{"src.txt"}, []string{"dst.txt"}, true},
		{[]string{"-t", "targetDir", "src1.txt", "src2.txt"}, []string{"src1.txt", "src2.txt"}, []string{"targetDir"}, true},
		{[]string{}, nil, nil, false},
	}
	TestParse(&Cp{}, tests, t)
}