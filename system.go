package main

import (
	"fmt"
)

var commandsSystem = []Pair[string, string]{
	CPair("name", "nuix system - NixOS related commands"),
	CPair("usage", "nuix system [COMMAND] [OPTIONS] - NixOS related commands"),
	CPair("switch", "Add new configuration to the system"),
	CPair("boot", "Add new configuration, but don't switch to it"),
	CPair("build", "Build new configuration, but don't add or switch to it"),
	CPair("test", "Try out new configuration"),
	CPair("dry-build", "Show what would be built and downloaded"),
	CPair("dry-activate", "Build new configuration, but only show changes"),
	CPair("edit", "Open configuration in $EDITOR"),
	CPair("repl", "Open configuration in nix repl"),
	CPair("build-image", "Build a disk image, specify variant with `--image-variant`"),
}

var optionsSystem = []Pair[string, string]{}

func system(args []string) {
	if len(args) == 0 {
		notSpecified("system")
	}
	fmt.Println(args)
}
