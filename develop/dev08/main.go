package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		executeCommand(strings.Fields(scanner.Text()))
	}
}
func executeCommand(commands []string) {
	args := commands[1:]
	command := commands[0]

	switch command {
	case `\quit`:
		fmt.Println("Shutting down")
		os.Exit(1)
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(dir)
	case "ps":
		res, err := ps.Processes()
		if err != nil {
			fmt.Println(err)
		}
		for _, process := range res {
			fmt.Printf("Name p: %v Pid: %v\n", process.Executable(), process.Pid())
		}
	case "cd":
		dir := strings.Join(args, "")
		os.Chdir(dir)
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "kill":
		pid, err := strconv.Atoi(strings.Join(args, ""))
		if err != nil {
			fmt.Println(err)
			break
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println(err)
			break
		}
		err = proc.Kill()
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("command not recognized")
	}

}
