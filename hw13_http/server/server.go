package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Server(ip, port string) {
	println("Сервер запущен", ip+":"+port)

	http.HandleFunc("/v1/hello", hello)
	http.HandleFunc("/v1/get_user", getUser)
	http.HandleFunc("/v1/create_user", createUser)

	server := &http.Server{
		Addr:         ip + ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/xml")

	user := User{
		ID:   1,
		Name: "Александр Лушин",
		Age:  32,
	}

	json.NewEncoder(w).Encode(user)
	// xml.NewEncoder(w).Encode(user)
}

type User struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	fmt.Printf("Новый пользователь: %+v\n", newUser)

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}
