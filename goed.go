package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	version     = "0.1.3"
	versionFlag = flag.Bool("version", false, "Display version and exit.")
	build       = flag.Bool("build", false, "Build script.")
	install     = flag.Bool("install", false, "Install script.")
)

func executeCommand(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	pausePrompt()
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

	cmds := []*exec.Cmd{
		exec.Command("gofmt", "-w", filename),
		exec.Command("editor", filename),
		exec.Command("gofmt", "-w", filename),
		exec.Command("golint", filename),
		exec.Command("go", "vet", filename),
	}

	for _, cmd := range cmds {
		executeCommand(cmd)
	}

	switch {
	case *build:
		fmt.Println("building...")
		executeCommand(exec.Command("go", "build", filename))
	case *install:
		fmt.Println("installing...")
		executeCommand(exec.Command("go", "install", filename))
	}
}
