package tasks

import (
	"encoding/json" // Like JSON in JS
	"os"
)

// Store tasks in a file (similar to localStorage but more manual)
const storageFile = "tasks.json"

func (tm *TaskManager) loadTasks() error {
	// JS analogy: fs.readFileSync but with explicit error handling
	data, err := os.ReadFile(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet - return empty slice
			tm.tasks = []Task{}
			return nil
		}
		return err
	}

	// Like JSON.parse() but with type safety
	return json.Unmarshal(data, &tm.tasks)
}

func (tm *TaskManager) saveTasks() error {
	// Like JSON.stringify()
	data, err := json.Marshal(tm.tasks)
	if err != nil {
		return err
	}

	// Write file with permissions (0644 = owner read/write, others read)
	return os.WriteFile(storageFile, data, 0644)
}