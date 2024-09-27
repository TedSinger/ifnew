package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"os"
	"os/exec"
	"path/filepath"
)

type MatchResult struct {
	SourceFiles []string
	TargetFiles []string
}

func (m MatchResult) String() string {
	return fmt.Sprintf("SourceFiles: %v, TargetFiles: %v", m.SourceFiles, m.TargetFiles)
}

type Command interface {
	Parse(args []string) (MatchResult, bool)
	Name() string
}

type TestCase struct {
	args          []string
	expectedSrc   []string
	expectedTgt   []string
	expectedMatch bool
}
func TestParse(c Command, cases []TestCase, t *testing.T) {
	for _, test := range cases {
		result, match := c.Parse(test.args)
		failed := false
		if match != test.expectedMatch {
			t.Logf("with args %v:\nexpected match %v, got %v", test.args, test.expectedMatch, match)
			failed = true
		}
		if !equal(result.SourceFiles, test.expectedSrc) {
			t.Logf("with args %v:\nexpected source files %v, got %v", test.args, test.expectedSrc, result.SourceFiles)
			failed = true
		}
		if !equal(result.TargetFiles, test.expectedTgt) {
			t.Logf("with args %v:\nexpected target files %v, got %v", test.args, test.expectedTgt, result.TargetFiles)
			failed = true
		}
		if failed {
			t.Errorf("with args %v:\nexpected match %v, got %v\nexpected source files %v, got %v\nexpected target files %v, got %v", 
				test.args, test.expectedMatch, match, test.expectedSrc, result.SourceFiles, test.expectedTgt, result.TargetFiles)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}


func validateTestCase(c Command, test TestCase, t *testing.T) bool {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testcase")
	if err != nil {
		return false
	}
	defer os.RemoveAll(tempDir)

	// Identify the project root and copy source files from test_data to the temporary directory
	projectRoot, err := os.Getwd()
	if err != nil {
		return false
	}
	testDataDir := filepath.Join(projectRoot, "test_data")

	for _, src := range test.expectedSrc {
		srcPath := filepath.Join(testDataDir, src)
		destPath := filepath.Join(tempDir, src)
		t.Logf("copying %v to %v", srcPath, destPath)
		data, err := os.ReadFile(srcPath)
		if err != nil {
			return false
		}
		if err := os.WriteFile(destPath, data, 0644); err != nil {
			return false
		}
	}

	// Change to the temporary directory
	originalDir, err := os.Getwd()
	if err != nil {
		return false
	}
	defer os.Chdir(originalDir)
	if err := os.Chdir(tempDir); err != nil {
		return false
	}

	// Run the command, including the command name
	cmd := exec.Command(c.Name(), test.args...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("PATH=%s", os.Getenv("PATH")))
	if err := cmd.Run(); err != nil {
		t.Logf("command %v failed: %v", cmd.Args, err)
		return false
	}

	// Check that the expected target files were created
	for _, tgt := range test.expectedTgt {
		tgtPath := filepath.Join(tempDir, tgt)
		if _, err := os.Stat(tgtPath); os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func TestValidateTestCase(c Command, tests []TestCase, t *testing.T) {
	for _, tt := range tests {
		if valid := validateTestCase(c, tt, t); valid != tt.expectedMatch {
			t.Errorf("validateTestCase(%v, %v) = %v, want %v", tt.args, tt.expectedMatch, valid, tt.expectedMatch)
		}
	}
}

