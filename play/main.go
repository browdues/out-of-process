// Times repeated execution of a do-nothing command.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	useStartProcess bool
	cmdWords        []string
	repetitions     = 200
)

func options() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Times repeated execution of a do-nothing command.

	args: [options] [repetitions]
  	repetitions
    	number of repetitions (default %d)
	`, repetitions)
		flag.PrintDefaults()
	}
	defaultDoNothingCommand := "true"
	if runtime.GOOS == "windows" {
		defaultDoNothingCommand = "cmd /c rem"
	}
	var cmd string
	flag.StringVar(&cmd, "c", defaultDoNothingCommand,
		"a no-op `command` that runs quickly")
	flag.BoolVar(&useStartProcess, "s", false,
		"use os.StartProcess to run process instead of exec.Command")
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "expected 0 or 1 arguments, got %d (%q)\n",
			len(args), args)
		flag.Usage()
		os.Exit(2)
	}
	if len(args) > 0 {
		var err error
		repetitions, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, `invalid "repetitions":`, err)
			os.Exit(2)
		}
	}
	cmdWords = strings.Fields(cmd)
	if len(cmdWords) == 0 {
		fmt.Fprintln(os.Stderr, "cannot execute empty command (-c option)")
		os.Exit(2)
	}
}

func main() {
	fmt.Println(`HI from "Go"`)
	options()
	resolved, err := exec.LookPath(cmdWords[0])
	if err != nil {
		resolved = cmdWords[0]
	} else {
		cmdWords[0] = resolved
	}

	// Run either via exec.Command or os.StartProcess.
	// (Tests show that both yield about the same timing.)
	var t time.Time
	if !useStartProcess {
		fmt.Println(" (using exec.Command).")
		t = time.Now()
		for i := 0; i < repetitions; i++ {
			var cmd *exec.Cmd
			if len(cmdWords) > 0 {
				cmd = exec.Command(resolved, cmdWords[1:]...)
			}
			cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
			err := cmd.Run()
			if err != nil {
				if _, ok := err.(*exec.ExitError); !ok {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(2)
				}
			}
		}
	} else {
		fmt.Println(" (using os.StartProcess).")
		procFiles := []*os.File{os.Stdin, os.Stdout, os.Stderr}
		t = time.Now()
		for i := 0; i < repetitions; i++ {
			p, err := os.StartProcess(resolved, cmdWords,
				&os.ProcAttr{Files: procFiles})
			if err != nil {
				if _, ok := err.(*exec.ExitError); !ok {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(2)
				}
			}
			_, err = p.Wait()
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println(time.Now().Sub(t))
}
