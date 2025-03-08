package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var tasks []string
	reader := bufio.NewReader(os.Stderr)

	fmt.Println("CLI Base to app task")
	fmt.Println("1. Add a Task")
	fmt.Println("2. List a Task")
	fmt.Println("3. Mark a task as Completed")
	fmt.Println("4. Deleted a Task")
	fmt.Println("5. Exit")

	input , _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

}