package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	version     = "0.3.0"
	versionFlag = flag.Bool("version", false, "Display version and exit.")
	build       = flag.Bool("build", false, "Build script.")
	install     = flag.Bool("install", false, "Install script.")
)

func runInteractively(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func enterToContinue(msg string) {
	r := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	_, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
}

func runQuietly(cmd *exec.Cmd) {
	var cmdout bytes.Buffer
	var cmderr bytes.Buffer
	cmd.Stdout = &cmdout
	cmd.Stderr = &cmderr
	err := cmd.Run()
	if err != nil {
		log.Print(err)
	}
	if cmdout.String() != "" {
		fmt.Print(cmdout.String())
		enterToContinue("Press enter to continue...")
	}
	if cmderr.String() != "" {
		fmt.Print(cmderr.String())
		enterToContinue("Errors returned. Press enter to continue...")
	}
}

func main() {
	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		os.Exit(1)
	}

	filename := flag.Arg(0)

	runQuietly(exec.Command("gofmt", "-w", filename))
	runInteractively(exec.Command("editor", filename))
	runQuietly(exec.Command("gofmt", "-w", filename))
	runQuietly(exec.Command("golint", filename))
	runQuietly(exec.Command("go", "vet", filename))

	switch {
	case *build:
		runQuietly(exec.Command("go", "build", filename))
	case *install:
		runQuietly(exec.Command("go", "install", filename))
	}
}
