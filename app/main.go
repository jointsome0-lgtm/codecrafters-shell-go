package main

import (
	"bufio"
	"fmt"
	"os"
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
		command := strings.TrimSpace(parts[0])
		if command == "exit" {
			break
		}

		if command == "echo" {
			result := strings.Join(parts[1:], " ")
			fmt.Println(result)
			continue
		}
		fmt.Println(command + ": command not found")
	}
}
