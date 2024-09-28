package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"github.com/mattn/go-shellwords"
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

type HumanReadableTestCase struct {
	cmd           string // "cp src dst"
	expectedSrcs  string // "src"
	expectedTgts  string // "dst"
	expectedMatch bool
}
type MachineReadableTestCase struct {
	args []string
	expectedSrcs []string
	expectedTgts []string
	expectedMatch bool
}

func (tc *HumanReadableTestCase) Parse() (MachineReadableTestCase, error) {
	args, err := shellwords.Parse(tc.cmd)
	if err != nil {
		return MachineReadableTestCase{}, err
	}
	expectedSrcs, err := shellwords.Parse(tc.expectedSrcs)
	if err != nil {
		return MachineReadableTestCase{}, err
	}
	expectedTgts, err := shellwords.Parse(tc.expectedTgts)
	if err != nil {
		return MachineReadableTestCase{}, err
	}

	return MachineReadableTestCase{args: args, expectedSrcs: expectedSrcs, expectedTgts: expectedTgts, expectedMatch: tc.expectedMatch}, nil
}

func TestParse(c Command, cases []HumanReadableTestCase, t *testing.T) {
	for _, test := range cases {
		parsed, err := test.Parse()
		if err != nil {
			t.Errorf("error parsing command %v: %v", test.cmd, err)
			continue
		}
		result, match := c.Parse(parsed.args[1:])
		failed := false
		if match != parsed.expectedMatch {
			t.Logf("with args %v:\nexpected match %v, got %v", parsed.args, parsed.expectedMatch, match)
			failed = true
		}
		
		if !equal(result.SourceFiles, parsed.expectedSrcs) {
			t.Logf("with args %v:\nexpected source files %v, got %v", parsed.args, parsed.expectedSrcs, result.SourceFiles)
			failed = true
		}
		if !equal(result.TargetFiles, parsed.expectedTgts) {
			t.Logf("with args %v:\nexpected target files %v, got %v", parsed.args, parsed.expectedTgts, result.TargetFiles)
			failed = true
		}
		if failed {
			t.Errorf("with args %v:\nexpected match %v, got %v\nexpected source files %v, got %v\nexpected target files %v, got %v", 
				parsed.args, parsed.expectedMatch, match, parsed.expectedSrcs, result.SourceFiles, parsed.expectedTgts, result.TargetFiles)
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


func validateTestCase(c Command, test MachineReadableTestCase, t *testing.T) bool {
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

	for _, src := range test.expectedSrcs {
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
	cmd := exec.Command(test.args[0], test.args[1:]...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("PATH=%s", os.Getenv("PATH")))
	if err := cmd.Run(); err != nil {
		t.Logf("command %v failed: %v", cmd.Args, err)
		return false
	}

	// Check that the expected target files were created
	for _, tgt := range test.expectedTgts {
		tgtPath := filepath.Join(tempDir, tgt)
		if _, err := os.Stat(tgtPath); os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func TestValidateTestCase(c Command, tests []HumanReadableTestCase, t *testing.T) {
	for _, tt := range tests {
		parsed, err := tt.Parse()
		if err != nil {
			t.Errorf("error parsing command %v: %v", tt.cmd, err)
			continue
		}
		if valid := validateTestCase(c, parsed, t); valid != tt.expectedMatch {
			t.Errorf("validateTestCase(%v, %v) = %v, want %v", tt.cmd, tt.expectedMatch, valid, tt.expectedMatch)
		}
	}
}

