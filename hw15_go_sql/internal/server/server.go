package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fixme_my_friend/hw15_go_sql/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	queries *db.Queries
}

func NewServer(dbPool *pgxpool.Pool) *Server {
	return &Server{
		queries: db.New(dbPool),
	}
}

func (s *Server) Start(ip, port string) {
	println("Сервер запущен", ip+":"+port)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/get_user", s.getUser)
	mux.HandleFunc("/v1/create_user", s.createUser)

	server := &http.Server{
		Addr:         ip + ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID пользователя из параметров запроса
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "ID пользователя не указан", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Некорректный ID пользователя", http.StatusBadRequest)
		return
	}

	user, err := s.queries.GetUserByID(context.Background(), int32(userID))
	if err != nil {
		http.Error(w, "Ошибка получения пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	var newUser db.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Ошибка обработки JSON", http.StatusBadRequest)
		return
	}

	createdUser, err := s.queries.CreateUser(context.Background(), newUser)
	if err != nil {
		http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
