package tasks

import (
	"fmt"
	"time"
)

// Task struct is similar to a JS object but with strict typing
type Task struct {
	ID          string    `json:"id"` // Like { id: string } in TS
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"` // Go has built-in time package
}

// NewTask is a constructor function - similar to factory functions in JS
func NewTask(description string) Task {
	return Task{
		ID:          generateID(),
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
}

// Method on Task type (similar to JS class methods)
func (t *Task) Complete() {
	t.Completed = true
}

// Private function (starts with lowercase)
func generateID() string {
	// JS analogy: Similar to Math.random().toString(36).substring(2)
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
