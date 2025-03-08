package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var tasks []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("CLI Base to app task")
	fmt.Println("1. Add a Task")
	fmt.Println("2. List a Task")
	fmt.Println("3. Mark a task as Completed")
	fmt.Println("4. Deleted a Task")
	fmt.Println("5. Exit")

	for {

		fmt.Print("Enter Your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, error := strconv.Atoi(input)
		if error != nil {
			fmt.Println("invalid input, please input an number")
		}
		switch choice {
		case 1:
			// Add a Task
			fmt.Print("Enter Task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			// add this task to the task list
			tasks = append(tasks, task)
			fmt.Println("Task added")
		case 2:
			// List a Task
			if len(tasks) == 0 {
				fmt.Println("no task is available")
			} else {
				fmt.Println("Your Tasks")
				for index, task := range tasks {
					fmt.Printf("%d - %s \n", index+1, task)
				}
			}
		case 3:
			// Mark a task as Completed
			if len(tasks) == 0 {
				fmt.Println("no task is available")
			}

			fmt.Print("Enter task number to complete")

			taskIdString, _ := reader.ReadString('\n')
			taskIdString = strings.TrimSpace(taskIdString)

			taskId, error := strconv.Atoi(taskIdString)

			if error == nil && taskId > 0 && taskId <= len(tasks) {
				// taskId = 1 => 0
				tasks[taskId-1] = "[+] " + tasks[taskId-1]
				fmt.Println("Task marked as completed: ")
			} else {
				fmt.Println("invalid input, please input an number")
			}
		case 4:
			// Deleted a Task
		case 5:
			// Exit
		}
	}
}
