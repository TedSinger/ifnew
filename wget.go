package main

import (
	"github.com/spf13/cobra"
	"strings"
)

type Wget struct{}

func (w *Wget) Parse(args []string) (MatchResult, bool) {
	rootCmd := &cobra.Command{
		Use: "wget",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.Flags().StringP("output-document", "O", "", "write documents to <file>")

	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		return MatchResult{}, false
	}

	if len(args) == 0 {
		return MatchResult{}, false

	}
	url := args[0]

	outputDocument, _ := rootCmd.Flags().GetString("output-document")

	var source_files, target_files []string
	source_files = []string{}
	if outputDocument == "" {
		parts := strings.Split(url, "/")
		defaultOutputDocument := parts[len(parts)-1]
		target_files = []string{defaultOutputDocument}
	} else {
		target_files = []string{outputDocument}
	}

	return MatchResult{SourceFiles: source_files, TargetFiles: target_files}, true
}

var _ Command = &Wget{}
