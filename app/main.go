package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			parts = []string{""}
		}

		commands := []string{
			"exit",
			"echo",
			"type",
			"pwd",
		}

		command := strings.TrimSpace(parts[0])

		if command == "exit" {
			break
		}

		if command == "echo" {
			result := strings.Join(parts[1:], " ")
			fmt.Println(result)
		} else if command == "type" {
			if slices.Contains(commands, parts[1]) {
				fmt.Printf("%s is a shell builtin\n", parts[1])
			} else if path, err := exec.LookPath(parts[1]); err == nil && path != "" {
				fmt.Printf("%s is %s\n", parts[1], path)
			} else {
				fmt.Printf("%s: not found\n", parts[1])
			}
		} else if command == "pwd" {
			wd, _ := os.Getwd()
			fmt.Println(wd)
		} else if path, err := exec.LookPath(parts[0]); err == nil && path != "" {
			args := parts[1:]
			c := exec.Command(command, args...)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Run()
		} else {
			fmt.Println(command + ": command not found")
		}
	}
}
