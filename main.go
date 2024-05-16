package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <query>")
		os.Exit(1)
	}

	query := strings.Join(os.Args[2:], "+")
	searchURL := fmt.Sprintf("https://duckduckgo.com/?q=w3m+%s&ia=answer", query)

	cmd := exec.Command("w3m", searchURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing w3m: %v\n", err)
		os.Exit(1)
	}
}
