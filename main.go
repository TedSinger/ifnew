package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var COMMANDS = map[string]Command{
	"cp":   &Cp{},
	"curl": &Curl{},
	"wget": &Wget{},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mm <command> [args...]")
		os.Exit(1)
	}
	commandName := os.Args[1]
	command, exists := COMMANDS[commandName]
	if !exists {
		fmt.Printf("Unknown command: %s\n", commandName)
		os.Exit(1)
	}

	matchResult, success := command.Parse(os.Args[2:])
	if !success {
		fmt.Println("Failed to parse command arguments")
		os.Exit(1)
	}

	var maxSourceModTime = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	var minTargetModTime time.Time
	allTargetsNewer := true

	for _, source := range matchResult.SourceFiles {
		sourceInfo, err := os.Stat(source)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if sourceInfo.ModTime().After(maxSourceModTime) {
			maxSourceModTime = sourceInfo.ModTime()
		}
	}

	for _, target := range matchResult.TargetFiles {
		targetInfo, err := os.Stat(target)
		if err != nil {
			allTargetsNewer = false
			break
		}
		if minTargetModTime.IsZero() || targetInfo.ModTime().Before(minTargetModTime) {
			minTargetModTime = targetInfo.ModTime()
		}
	}

	if maxSourceModTime.After(minTargetModTime) {
		allTargetsNewer = false
	}

	if allTargetsNewer {
		fmt.Println("All target files are newer than source files. No action taken.")
	} else {
		cmd := exec.Command(commandName, os.Args[2:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
}
