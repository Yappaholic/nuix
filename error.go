package main

import (
	"fmt"
	"os"
)

func notSpecifiedGlobal() {
	err := fmt.Errorf("No command specified")
	fmt.Fprintf(os.Stderr, "%s\nUse \"nuix -h\" or \"nuix --help\" to see available commands", err.Error())
	os.Exit(1)
}
func notSpecified(command string) {
	err := fmt.Errorf("No command specified")
	fmt.Fprintf(os.Stderr, "%s\nUse \"nuix %s -h\" or \"nuix %s --help\" to see available commands", err.Error(), command, command)
	os.Exit(1)
}

func doesntExist(command string, argument string) {
	err := fmt.Errorf("Command does not exist: %s", argument)
	fmt.Fprintf(os.Stderr, "%s\nUse \"nuix %s -h\" or \"nuix %s --help\" to see available commands", err.Error(), command, command)
	os.Exit(1)
}

func doesntExistGlobal(argument string) {
	err := fmt.Errorf("Command does not exist: %s", argument)
	fmt.Fprintf(os.Stderr, "%s\nUse \"nuix -h\" or \"nuix --help\" to see available commands", err.Error())
	os.Exit(1)
}
