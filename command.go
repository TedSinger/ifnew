package main

import "fmt"

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
