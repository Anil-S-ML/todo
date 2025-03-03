package todo

// Todo struct represents a single task.
type Todo struct {
    ID        int
    Title     string
    Completed bool
}

// MarkComplete marks the task as completed.
func (t *Todo) MarkComplete() {
    t.Completed = true
}
