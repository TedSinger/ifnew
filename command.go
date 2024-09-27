package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
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
