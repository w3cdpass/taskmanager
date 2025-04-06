package main

import (
	"fmt"
	"os"
	"taskmanager/tasks"
)

func main() {
	manager := tasks.NewTaskManager()
	
	// JS analogy: Similar to process.argv in Node.js
	args := os.Args[1:]
	
	if len(args) == 0 {
		fmt.Println("Welcome to Go Task Manager!")
		fmt.Println("Commands: add, list, complete, delete, exit")
		manager.InteractiveMode()
		return
	}

	// Handle CLI arguments
	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Please provide a task description")
			return
		}
		manager.AddTask(args[1])
	case "list":
		manager.ListTasks()
	case "complete":
		if len(args) < 2 {
			fmt.Println("Please provide a task ID")
			return
		}
		manager.CompleteTask(args[1])
	case "delete":
		if len(args) < 2 {
			fmt.Println("Please provide a task ID")
			return
		}
		manager.DeleteTask(args[1])
	default:
		fmt.Println("Unknown command")
	}
}