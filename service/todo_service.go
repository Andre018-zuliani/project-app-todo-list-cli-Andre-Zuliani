package service

import (
	"encoding/json"
	"errors"
	"os"
	"project-app-todo-list-cli-nama/model"
	"strings"
)

const dataFile = "data/todo.json"

// Load all tasks from JSON file
func LoadTasks() ([]model.Task, error) {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return []model.Task{}, nil // jika file belum ada, kembalikan array kosong
	}

	var tasks []model.Task
	json.Unmarshal(file, &tasks)
	return tasks, nil
}

// Save tasks back to JSON
func SaveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

// Add new task with default Status & Priority
func AddTask(title string, description string, priority string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("judul tidak boleh kosong")
	}

	tasks, _ := LoadTasks()

	// Validasi duplikat nama tugas
	for _, t := range tasks {
		if strings.EqualFold(t.Title, title) {
			return errors.New("judul tugas sudah ada, gunakan nama lain")
		}
	}

	// Validasi priority
	if priority == "" {
		priority = "low"
	}

	newID := len(tasks) + 1

	newTask := model.Task{
		ID:          newID,
		Title:       title,
		Description: description,
		Status:      "pending",
		Priority:    priority,
	}

	tasks = append(tasks, newTask)

	return SaveTasks(tasks)
}

// Mark task as completed
func MarkDone(id int) error {
	tasks, _ := LoadTasks()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "completed"
			return SaveTasks(tasks)
		}
	}

	return errors.New("ID tidak ditemukan")
}

// Delete task by ID
func DeleteTask(id int) error {
	tasks, _ := LoadTasks()

	newTasks := []model.Task{}
	found := false

	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("ID tidak ditemukan")
	}

	return SaveTasks(newTasks)
}

// Search tasks by title or description (case insensitive)
func SearchTask(keyword string) ([]model.Task, error) {
	tasks, _ := LoadTasks()

	keyword = strings.ToLower(strings.TrimSpace(keyword))
	if keyword == "" {
		return tasks, nil
	}

	var results []model.Task
	for _, t := range tasks {
		if strings.Contains(strings.ToLower(t.Title), keyword) ||
			strings.Contains(strings.ToLower(t.Description), keyword) {
			results = append(results, t)
		}
	}

	return results, nil
}
