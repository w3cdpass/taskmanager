package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TaskManager is similar to a JS class
type TaskManager struct {
	tasks []Task
}

// Constructor function (like `new Class()` in JS)
func NewTaskManager() *TaskManager {
	tm := &TaskManager{}
	if err := tm.loadTasks(); err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
	}
	return tm
}

func (tm *TaskManager) AddTask(description string) {
	task := NewTask(description)
	tm.tasks = append(tm.tasks, task) // Similar to array.push()
	if err := tm.saveTasks(); err != nil {
		fmt.Printf("Error saving task: %v\n", err)
	}
	fmt.Printf("Added task: %s\n", description)
}

func (tm *TaskManager) ListTasks() {
	if len(tm.tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for i, task := range tm.tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}
		// Similar to template literals but with fmt.Printf
		fmt.Printf("%d. [%s] %s (ID: %s)\n", i+1, status, task.Description, task.ID)
	}
}

func (tm *TaskManager) CompleteTask(id string) {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks[i].Complete()
			if err := tm.saveTasks(); err != nil {
				fmt.Printf("Error completing task: %v\n", err)
			}
			fmt.Println("Task marked as complete")
			return
		}
	}
	fmt.Println("Task not found")
}

func (tm *TaskManager) DeleteTask(id string) {
	for i, task := range tm.tasks {
		if task.ID == id {
			// Similar to array.splice(i, 1)
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			if err := tm.saveTasks(); err != nil {
				fmt.Printf("Error deleting task: %v\n", err)
			}
			fmt.Println("Task deleted")
			return
		}
	}
	fmt.Println("Task not found")
}

func (tm *TaskManager) InteractiveMode() {
	// JS analogy: Similar to readline.createInterface()
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		
		if input == "exit" {
			break
		}

		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]
		
		switch cmd {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Please provide a task description")
				continue
			}
			tm.AddTask(parts[1])
		case "list":
			tm.ListTasks()
		case "complete":
			if len(parts) < 2 {
				fmt.Println("Please provide a task ID")
				continue
			}
			tm.CompleteTask(parts[1])
		case "delete":
			if len(parts) < 2 {
				fmt.Println("Please provide a task ID")
				continue
			}
			tm.DeleteTask(parts[1])
		default:
			fmt.Println("Unknown command. Try: add, list, complete, delete, exit")
		}
	}
}