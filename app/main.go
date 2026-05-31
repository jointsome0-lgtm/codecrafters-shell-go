package main

import (
	"bufio"
	"fmt"
	"os"
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
			} else {
				fmt.Printf("%s: not found\n", parts[1])
			}
		} else {
			fmt.Println(command + ": command not found")
		}
	}
}
