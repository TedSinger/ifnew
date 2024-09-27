package main

import (
	"github.com/spf13/cobra"
)

type Cp struct{}

func (c *Cp) Parse(args []string) (MatchResult, bool) {
	rootCmd := &cobra.Command{
		Use: "cp",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.Flags().BoolP("recursive", "r", false, "copy directories recursively")
	rootCmd.Flags().BoolP("verbose", "v", false, "explain what is being done")
	rootCmd.Flags().BoolP("force", "f", false, "if an existing destination file cannot be opened, remove it and try again")
	rootCmd.Flags().BoolP("interactive", "i", false, "prompt before overwrite")
	rootCmd.Flags().StringP("target", "t", "", "specify target directory")

	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		return MatchResult{}, false
	}

	remainingArgs := rootCmd.Flags().Args()
	if len(remainingArgs) < 2 {
		return MatchResult{}, false
	}
	targetDir, _ := rootCmd.Flags().GetString("target")

	var source_files, target_files []string
	if targetDir != "" {
		target_files = []string{targetDir}
		source_files = remainingArgs
	} else {
		source_files = []string{remainingArgs[0]}
		target_files = []string{remainingArgs[1]}
	}

	return MatchResult{SourceFiles: source_files, TargetFiles: target_files}, true
}

var _ Command = &Cp{}
