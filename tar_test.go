package main

import "testing"

func TestTar(t *testing.T) {
	tests := []HumanReadableTestCase{
		{"tar -c -f archive.tar file1.txt file2.txt", "file1.txt file2.txt", "archive.tar", true},
		{"tar cf archive.tar file1.txt file2.txt", "file1.txt file2.txt", "archive.tar", true},
		{"tar xf archive.tar file1.txt file2.txt", "archive.tar", "file1.txt file2.txt", true},
		// not testing `tar x archive.tar` yet because of the implicit wildcard
		// maybe we can `tar t archive.tar`
		{"tar", "", "", false},
	}
	TestParse(&Tar{}, tests, t)
}

func TestTarValidateTestCase(t *testing.T) {
	tests := []HumanReadableTestCase{
		{"tar -c -f archive.tar file1.txt file2.txt", "file1.txt file2.txt", "archive.tar", true},
		{"tar cf archive.tar file1.txt file2.txt", "file1.txt file2.txt", "archive.tar", true},
		{"tar xf archive.tar file1.txt file2.txt", "archive.tar", "file1.txt file2.txt", true},
	}
	TestValidateTestCase(&Tar{}, tests, t)
}
