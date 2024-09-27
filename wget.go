package main

import (
	"github.com/spf13/cobra"
	"path"
	"net/url"
)

type Wget struct{}

func (w *Wget) Parse(args []string) (MatchResult, bool) {
	rootCmd := &cobra.Command{
		Use: "wget",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.Flags().StringP("output-document", "O", "", "write documents to <file>")
	rootCmd.Flags().StringP("url", "u", "", "specify a URL to download")
	rootCmd.Flags().StringP("output-file", "o", "", "log messages to <file>")

	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		return MatchResult{}, false
	}

	if len(args) == 0 {
		return MatchResult{}, false

	}
	urlStr, _ := rootCmd.Flags().GetString("url")
	if urlStr == "" {
		urlStr = args[0]
	}

	outputDocument, _ := rootCmd.Flags().GetString("output-document")
	outputFile, _ := rootCmd.Flags().GetString("output-file")

	target_files := []string{}
	if outputDocument == "" {
		parsedUrl, err := url.Parse(urlStr)
		if err != nil {
			return MatchResult{}, false
		}
		defaultOutputDocument := path.Base(parsedUrl.Path)
		if defaultOutputDocument == "" || defaultOutputDocument == "/" || defaultOutputDocument == "." {
			defaultOutputDocument = "index.html"
		}
		target_files = append(target_files, defaultOutputDocument)
	} else {
		target_files = append(target_files, outputDocument)
	}
	if outputFile != "" {
		target_files = append(target_files, outputFile)
	}

	return MatchResult{SourceFiles: []string{}, TargetFiles: target_files}, true
}

var _ Command = &Wget{}
