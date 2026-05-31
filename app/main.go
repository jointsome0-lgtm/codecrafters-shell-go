package main

import (
	"fmt"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")

	var command string 
	_, err := fmt.Scanln(&command)
	if err != nil {
		fmt.Println("Reding input error:", err)
		return
	}

	fmt.Printf("%v: command not found", command)
}
