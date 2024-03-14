package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tasks(w, r)
		case http.MethodPost:
			createTasks(w, r)
		case http.MethodDelete:
			deleteTasks(w, r)
		default:
			http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("start http server :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
