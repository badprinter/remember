package models

// Task
type Task struct {
	ID       uint
	USER_ID  uint
	Document string
}

func NewTask(RawDocument string) *Task {
	return &Task{
		Document: RawDocument,
	}
}
