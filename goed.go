package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	version     = "0.2.1"
	versionFlag = flag.Bool("version", false, "Display version and exit.")
	build       = flag.Bool("build", false, "Build script.")
	install     = flag.Bool("install", false, "Install script.")
)

func executeCommand(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func pausePrompt() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to continue...")
	_, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		os.Exit(1)
	}

	filename := flag.Arg(0)

	cmds := []struct {
		label   string
		command *exec.Cmd
		pause   bool
	}{
		{"### FORMAT ###################", exec.Command("gofmt", "-w", filename), true},
		{"### EDIT   ###################", exec.Command("editor", filename), false},
		{"### FORMAT ###################", exec.Command("gofmt", "-w", filename), false},
		{"### LINT   ###################", exec.Command("golint", filename), false},
		{"### VET    ###################", exec.Command("go", "vet", filename), false},
	}

	for _, cmd := range cmds {

		fmt.Println(cmd.label)
		executeCommand(cmd.command)
		if cmd.pause {
			pausePrompt()
		}
	}

	switch {
	case *build:
		fmt.Println("### BUILD ####################")
		executeCommand(exec.Command("go", "build", filename))
	case *install:
		fmt.Println("### INSTALL ##################")
		executeCommand(exec.Command("go", "install", filename))
	}
}
