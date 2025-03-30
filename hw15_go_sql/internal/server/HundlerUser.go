package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/db"
)

// CreateUserHandler - обработчик для создания пользователя.
// Использует sqlc-функцию CreateUser.
func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}
	// Декодируем входящий JSON в структуру CreateUserParams
	var input db.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
		return
	}
	// Вызываем сгенерированную функцию CreateUser
	userID, err := s.queries.CreateUser(context.Background(), input)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка создания пользователя")
		return
	}
	respondJSON(w, http.StatusCreated, map[string]interface{}{
		"idUserMain": userID,
		"message":    "Пользователь создан",
	})
}

// DeleteUserHandler - обработчик для удаления пользователя по email.
// Использует sqlc-функцию DeleteUser.
func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}
	// Получаем email из query-параметров
	email := r.URL.Query().Get("email")
	if email == "" {
		respondError(w, http.StatusBadRequest, "Email не указан")
		return
	}
	// Вызываем sqlc-функцию DeleteUser
	if err := s.queries.DeleteUser(context.Background(), email); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка удаления пользователя")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Пользователь удалён"})
}

// GetUserSpendingStatsHandler - обработчик для получения статистики расходов пользователей.
// Использует sqlc-функцию GetUserSpendingStats.
func (s *Server) GetUserSpendingStatsHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}
	// Вызываем sqlc-функцию GetUserSpendingStats
	stats, err := s.queries.GetUserSpendingStats(context.Background())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка получения статистики")
		return
	}
	respondJSON(w, http.StatusOK, stats)
}

// GetUsersByPasswordHandler - обработчик для получения пользователей с заданным паролем.
// Использует sqlc-функцию GetUsersByPassword.
func (s *Server) GetUsersByPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}
	// Получаем пароль из query-параметров
	password := r.URL.Query().Get("password")
	if password == "" {
		respondError(w, http.StatusBadRequest, "Пароль не указан")
		return
	}
	// Вызываем sqlc-функцию GetUsersByPassword
	users, err := s.queries.GetUsersByPassword(context.Background(), password)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка получения пользователей")
		return
	}
	respondJSON(w, http.StatusOK, users)
}

// UpdateUserNameHandler - обработчик для обновления имени пользователя по email.
// Использует sqlc-функцию UpdateUserName.
func (s *Server) UpdateUserNameHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodPut {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}
	// Декодируем входящий JSON в структуру UpdateUserNameParams
	var input db.UpdateUserNameParams
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
		return
	}
	// Вызываем sqlc-функцию UpdateUserName
	if err := s.queries.UpdateUserName(context.Background(), input); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка обновления имени пользователя")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Имя пользователя обновлено"})
}
