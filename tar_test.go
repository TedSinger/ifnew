package main

import "testing"

func TestTar(t *testing.T) {
	tests := []TestCase{
		{[]string{"-c", "-f", "archive.tar", "file1.txt", "file2.txt"}, []string{"file1.txt", "file2.txt"}, []string{"archive.tar"}, true},
		{[]string{"cf", "archive.tar", "file1.txt", "file2.txt"}, []string{"file1.txt", "file2.txt"}, []string{"archive.tar"}, true},
		{[]string{"x", "f", "archive.tar", "file1.txt", "file2.txt"}, []string{"archive.tar"}, []string{"file1.txt", "file2.txt"}, true},
		// not testing `tar x archive.tar` yet because of the implicit wildcard
		// maybe we can `tar t archive.tar`
		{[]string{}, nil, nil, false},
	}
	TestParse(&Tar{}, tests, t)
}

func TestTarValidateTestCase(t *testing.T) {
	tests := []TestCase{
		{[]string{"-c", "-f", "archive.tar", "file1.txt", "file2.txt"}, []string{"file1.txt", "file2.txt"}, []string{"archive.tar"}, true},
		{[]string{"cf", "archive.tar", "file1.txt", "file2.txt"}, []string{"file1.txt", "file2.txt"}, []string{"archive.tar"}, true},
		{[]string{"x", "f", "archive.tar", "file1.txt", "file2.txt"}, []string{"archive.tar"}, []string{"file1.txt", "file2.txt"}, true},
	}
	TestValidateTestCase(&Tar{}, tests, t)
}
