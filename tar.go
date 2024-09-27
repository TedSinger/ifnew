package main

import (
	"github.com/spf13/cobra"
	"regexp"
)

type Tar struct{}

func (t *Tar) Parse(args []string) (MatchResult, bool) {
	rootCmd := &cobra.Command{
		Use: "tar",
		Run: func(cmd *cobra.Command, args []string) {},
	}
	newArgs := []string{}
	for _, arg := range args {
		if len(arg) > 1 && arg[0] != '-' {
			if regexp.MustCompile(`^[cvfzx]+$`).MatchString(arg) {
				newArgs = append(newArgs, "-"+string(arg))
			} else {
				newArgs = append(newArgs, arg)
			}
		} else if len(arg) == 1 {
			newArgs = append(newArgs, "-" + arg)
		} else if arg != "" {
			newArgs = append(newArgs, arg)
		}
	}
	args = newArgs
	rootCmd.Flags().BoolP("create", "c", false, "create a new archive")
	rootCmd.Flags().BoolP("extract", "x", false, "extract files from archive")
	rootCmd.Flags().StringP("file", "f", "", "use archive file or device <file>")
	rootCmd.Flags().StringP("directory", "C", "", "extract to <directory>")

	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		return MatchResult{}, false
	}

	create, _ := rootCmd.Flags().GetBool("create")
	extract, _ := rootCmd.Flags().GetBool("extract")
	file, _ := rootCmd.Flags().GetString("file")
	directory, _ := rootCmd.Flags().GetString("directory")
	remainingArgs := rootCmd.Flags().Args()

	if create && file == "" {
		return MatchResult{}, false
	}
	if extract && file == "" {
		return MatchResult{}, false
	}

	source_files := []string{}
	target_files := []string{}
	if create {
		target_files = append(target_files, file)
		source_files = append(source_files, remainingArgs...)
	} else if extract {
		source_files = append(source_files, file)
		if directory != "" {
			for _, file := range remainingArgs {
				target_files = append(target_files, directory+"/"+file)
			}
		} else {
			target_files = append(target_files, remainingArgs...)
		}
	} else {
		return MatchResult{}, false
	}

	return MatchResult{SourceFiles: source_files, TargetFiles: target_files}, true
}

var _ Command = &Tar{}