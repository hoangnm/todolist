package main

import (
	"fmt"
	"net/http"

	"hoangnm/todolist/db"
	"hoangnm/todolist/services"

	_ "github.com/mattn/go-sqlite3"
)

type HttpHeader struct {
	name  string
	value string
}

type HttpMiddleware struct {
	handler http.Handler
	headers []HttpHeader
}

func NewHttpMiddleware(handler http.Handler, headers []HttpHeader) *HttpMiddleware {
	return &HttpMiddleware{
		handler,
		headers,
	}
}

func (m *HttpMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, header := range m.headers {
		w.Header().Add(header.name, header.value)
	}
	m.handler.ServeHTTP(w, r)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		services.GetTasks(w, r)
		return
	}
	if r.Method == http.MethodPost {
		services.CreateTask(w, r)
		return
	}
	if r.Method == http.MethodPut {
		services.UpdateTask(w, r)
	}
}

func main() {
	db.InitDB()

	router := http.NewServeMux()
	router.HandleFunc("/tasks", handleTasks)
	headers := []HttpHeader{
		{"Access-Control-Allow-Origin", "*"},
	}
	wrappedRouter := NewHttpMiddleware(router, headers)
	err := http.ListenAndServe(":8080", wrappedRouter)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
