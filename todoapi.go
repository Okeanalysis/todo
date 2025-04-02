package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var (
	tasks  = []Todo{}
	nextID = 1
	mu     sync.Mutex
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	todo.ID = nextID
	nextID++
	tasks = append(tasks, todo)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func removeTask(w http.ResponseWriter, r *http.Request) {
	var id int
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("/tasks/create", createTask)
	http.HandleFunc("/tasks/remove", removeTask)

	fmt.Println("Starting server on : 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
