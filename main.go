package main

import (
	"fmt"
	"os"
	"strings"
)

type Application struct {
	args []string
}

var app = Application{}

var commandsGlobal = []Pair[string, string]{
	CPair("name", "Comfiness of guix cli for nix/NixOS users"),
	CPair("usage", "nuix [COMMAND] [OPTIONS]"),
	CPair("system", "NixOS related commands"),
	CPair("home", "Configure your home"),
	CPair("gc", "Clean your nix store"),
	CPair("search", "Search for packages"),
	CPair("package", "Manage installed nix packages(not recommended)"),
	CPair("channels", "Manage nix channels"),
	CPair("shell", "Create a new shell environment"),
	CPair("download", "Fetch source and get hash output"),
}

func formatOptions(commands []Pair[string, string]) {
	help_len := 15
	for i := range commands {
		spacer := strings.Repeat(" ", help_len-len(Fst(commands[i], nil).(string)))
		fmt.Println(" ", Fst(commands[i], nil), spacer, Snd(commands[i], nil))
	}
}

func showHelp(args []string) {
	var commands []Pair[string, string]
	var name any
	var usage any
	if len(args) > 1 && args[0] != "-h" && args[1] != "--help" {
		switch args[0] {
		case "system":
			commands = commandsSystem
		}
	} else {
		commands = commandsGlobal
	}
	name = Snd(FindPair(commands, "name"))
	usage = Snd(FindPair(commands, "usage"))
	if err := DeletePair(commands, "name"); err != nil {
		fmt.Fprintf(os.Stderr, "Encountered error: %s", err)
		os.Exit(1)
	}
	if err := DeletePair(commands, "usage"); err != nil {
		fmt.Fprintf(os.Stderr, "Encountered error: %s", err)
		os.Exit(1)
	}
	fmt.Println(name)
	fmt.Printf("Usage: %s\n", usage)
	fmt.Println("Available commands:")
	formatOptions(commands)
	os.Exit(0)
}

func isHelp(args []string) {
	for _, v := range args {
		if v == "-h" || v == "--help" {
			showHelp(args)
		}
	}
}

func parseCommands(args []string) {
	isHelp(args)
	if len(args) == 0 {
		notSpecifiedGlobal()
	}
	if _, err := FindPair(commandsGlobal, args[0]); err != nil {
		doesntExistGlobal(args[0])
	} else {
		switch args[0] {
		case "system":
			system(args[1:])

		}
	}
}

func main() {
	app.args = os.Args[1:]
	parseCommands(app.args)
}
