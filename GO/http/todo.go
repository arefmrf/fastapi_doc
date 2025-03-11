package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Todo  struct  with json tags
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// just work as a DB for us!
var (
	todos []Todo
	mu    sync.Mutex // To ensure thread safety when modifying the slice
)

func sendJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//err := json.NewEncoder(w).Encode(todos)
	//if err != nil {
	//	log.Fatal("Error encoding json: ", todos)
	//}
	sendJSONResponse(w, http.StatusOK, todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//var todo Todo
	//json.NewDecoder(r.Body).Decode(&todo)
	//todo.ID = len(todos) + 1
	//todos = append(todos, todo)
	//json.NewEncoder(w).Encode(todo)

	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Printf("Invalid request body: %v", err)
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	todo.ID = len(todos) + 1 // Note: This is not ideal for production.
	todos = append(todos, todo)
	sendJSONResponse(w, http.StatusCreated, todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//var todo Todo
	//json.NewDecoder(r.Body).Decode(&todo)
	//for i, t := range todos {
	//	if t.ID == todo.ID {
	//		todos[i].Status = todo.Status
	//		json.NewEncoder(w).Encode(todos[i])
	//		return
	//	}
	//}
	//w.WriteHeader(http.StatusNotFound)
	//json.NewEncoder(w).Encode(map[string]string{"message": "TODO not found"})

	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		log.Printf("Invalid request body: %v", err)
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, t := range todos {
		if t.ID == updatedTodo.ID {
			todos[i].Status = updatedTodo.Status
			sendJSONResponse(w, http.StatusOK, todos[i])
			return
		}
	}

	sendJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Todo not found"})
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//var todo Todo
	//json.NewDecoder(r.Body).Decode(&todo)
	//for i, t := range todos {
	//	if t.ID == todo.ID {
	//		todos = append(todos[:i], todos[i+1:]...)
	//		json.NewEncoder(w).Encode(map[string]string{"message": "TODO deleted"})
	//		return
	//	}
	//}
	//w.WriteHeader(http.StatusNotFound)
	//json.NewEncoder(w).Encode(map[string]string{"message": "TODO not found"})

	var todoToDelete Todo
	if err := json.NewDecoder(r.Body).Decode(&todoToDelete); err != nil {
		log.Printf("Invalid request body: %v", err)
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for i, t := range todos {
		if t.ID == todoToDelete.ID {
			todos = append(todos[:i], todos[i+1:]...)
			sendJSONResponse(w, http.StatusOK, map[string]string{"message": "Todo deleted"})
			return
		}
	}
	sendJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Todo not found"})
}

func main() {
	//http.HandleFunc("/todos", getTodos)
	//http.HandleFunc("/todos/add", addTodo)
	//http.HandleFunc("/todos/update", updateTodo)
	//http.HandleFunc("/todos/delete", deleteTodo)

	//fmt.Println("Server starting at port 8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))

	// Create a new router
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", getTodos)
	mux.HandleFunc("/todos/add", addTodo)
	mux.HandleFunc("/todos/update", updateTodo)
	mux.HandleFunc("/todos/delete", deleteTodo)

	// Start the server
	serverAddr := ":8080"
	fmt.Printf("Server starting at %s...\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, mux))
}
