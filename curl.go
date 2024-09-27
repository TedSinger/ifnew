package main

import (
	"github.com/spf13/cobra"
)

type Curl struct{}

func (c *Curl) Name() string {
	return "curl"
}

func (c *Curl) Parse(args []string) (MatchResult, bool) {
	rootCmd := &cobra.Command{
		Use: "curl",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.Flags().StringP("url", "u", "", "specify the URL to fetch")
	rootCmd.Flags().BoolP("verbose", "v", false, "make the operation more talkative")
	rootCmd.Flags().BoolP("silent", "s", false, "silent mode")
	rootCmd.Flags().StringP("output", "o", "", "write output to <file> instead of stdout")

	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		return MatchResult{}, false
	}

	url, _ := rootCmd.Flags().GetString("url")
	output, _ := rootCmd.Flags().GetString("output")

	if url == "" {
		if len(args) > 0 {
			url = args[0]
		} else {
			return MatchResult{}, false
		}
	}

	var source_files, target_files []string
	source_files = []string{}
	if output != "" {
		target_files = []string{output}
	}

	return MatchResult{SourceFiles: source_files, TargetFiles: target_files}, true
}

var _ Command = &Curl{}
