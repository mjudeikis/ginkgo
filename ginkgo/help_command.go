package main

import (
	"flag"
	"fmt"
)

func BuildHelpCommand() *Command {
	return &Command{
		Name:         "help",
		FlagSet:      flag.NewFlagSet("help", flag.ExitOnError),
		UsageCommand: "ginkgo help <COMAND>",
		Usage: []string{
			"Print usage information.  If a command is passed in, print usage information just for that command.",
		},
		Command: printHelp,
	}
}

func printHelp(args []string) {
	if len(args) == 0 {
		usage()
	} else {
		command, found := commandMatching(args[0])
		if !found {
			complainAndQuit(fmt.Sprintf("Unkown command: %s", args[0]))
		}

		usageForCommand(command, true)
	}
}
