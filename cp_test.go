package main

import "testing"

func TestCp(t *testing.T) {
	tests := []HumanReadableTestCase{
		{"cp src.txt dst.txt", "src.txt", "dst.txt", true},
		{"cp -t targetDir src1.txt src2.txt", "src1.txt src2.txt", "targetDir", true},
		{"cp", "", "", false},
	}
	TestParse(&Cp{}, tests, t)
}