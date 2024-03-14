package main

import (
	"context"
	"net/http"
	"time"
)

func tasks(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := db.SelectContext(ctx, &tasks, `SELECT * FROM tasks`); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := serialize(w, tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTasks(w http.ResponseWriter, r *http.Request) {
	task := Task{}
	if err := deserialize(r.Body, &task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if _, err := db.ExecContext(ctx, `
		INSERT INTO tasks (name, title, status) VALUES (?, ?, ?)
	`, task.Name, task.Title, task.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func deleteTasks(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if _, err := db.ExecContext(ctx, `DELETE FROM tasks WHERE id = ?`, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
