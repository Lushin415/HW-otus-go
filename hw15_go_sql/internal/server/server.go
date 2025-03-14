package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func NewServer(dbPool *pgxpool.Pool) *Server {
	return &Server{
		queries: db.New(dbPool),
		pool:    dbPool,
	}
}

func (s *Server) Start(ip, port string) {
	log.Printf("Сервер запущен на %s:%s\n", ip, port)

	mux := http.NewServeMux()

	// Эндпоинты
	mux.HandleFunc("/v1/get_user", s.getUser)
	mux.HandleFunc("/v1/create_user", s.createUser)
	mux.HandleFunc("/v1/products/price_range", s.getProductsByPriceRange)
	mux.HandleFunc("/v1/update_product_price", s.updateProductPrice)

	server := &http.Server{
		Addr:         ip + ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}

// Вспомогательные функции
func respondJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
	}
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

// Обработчик получения пользователя
func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		respondError(w, http.StatusBadRequest, "ID пользователя не указан")
		return
	}

	userID, err := strconv.Atoi(idParam)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Некорректный ID пользователя")
		return
	}

	// Используем транзакцию
	tx, err := s.pool.Begin(context.Background())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
		return
	}
	defer tx.Rollback(context.Background())

	// Создаем новый экземпляр Queries внутри транзакции
	qtx := s.queries.WithTx(tx)
	user, err := qtx.GetUserByID(context.Background(), int32(userID))
	if err != nil {
		if err.Error() == "no rows in result set" {
			respondError(w, http.StatusNotFound, "Пользователь не найден")
		} else {
			log.Printf("Ошибка БД при получении пользователя: %v", err)
			respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
		}
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
		return
	}

	respondJSON(w, http.StatusOK, user)
}

// Обработчик создания пользователя
func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	var input struct {
		Name     string `json:"nameUser"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
		return
	}

	// Используем транзакцию
	tx, err := s.pool.Begin(context.Background())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
		return
	}
	defer tx.Rollback(context.Background())

	// Создаем параметры для запроса
	params := db.CreateUserParams{
		NameUser: input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	qtx := s.queries.WithTx(tx)
	userID, err := qtx.CreateUser(context.Background(), params)
	if err != nil {
		log.Printf("Ошибка создания пользователя: %v", err)
		respondError(w, http.StatusInternalServerError, "Ошибка создания пользователя")
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
		return
	}

	respondJSON(w, http.StatusCreated, map[string]interface{}{
		"idUserMain": userID,
		"message":    "Пользователь создан",
	})
}

// Обработчик получения продуктов по ценовому диапазону
func (s *Server) getProductsByPriceRange(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	minStr := r.URL.Query().Get("min")
	maxStr := r.URL.Query().Get("max")
	if minStr == "" || maxStr == "" {
		respondError(w, http.StatusBadRequest, "Укажите диапазон цен min и max")
		return
	}

	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Некорректная минимальная цена")
		return
	}

	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Некорректная максимальная цена")
		return
	}

	// Используем транзакцию
	tx, err := s.pool.Begin(context.Background())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
		return
	}
	defer tx.Rollback(context.Background())

	// Используем прямой SQL
	rows, err := tx.Query(context.Background(),
		"SELECT id_product_main, name_product, price FROM schema.Products WHERE price BETWEEN $1 AND $2",
		min, max)

	if err != nil {
		log.Printf("Ошибка получения продуктов: %v", err)
		respondError(w, http.StatusInternalServerError, "Ошибка получения продуктов")
		return
	}
	defer rows.Close()

	var products []db.SchemaProduct
	for rows.Next() {
		var product db.SchemaProduct
		if err := rows.Scan(&product.IDProductMain, &product.NameProduct, &product.Price); err != nil {
			log.Printf("Ошибка сканирования продукта: %v", err)
			respondError(w, http.StatusInternalServerError, "Ошибка чтения данных продукта")
			return
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Ошибка в курсоре: %v", err)
		respondError(w, http.StatusInternalServerError, "Ошибка обработки результатов")
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
		return
	}

	respondJSON(w, http.StatusOK, products)
}

// Обработчик обновления цены продукта
func (s *Server) updateProductPrice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	var input struct {
		ID    int32   `json:"id_product_main"`
		Price float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
		return
	}

	// Используем транзакцию
	tx, err := s.pool.Begin(context.Background())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
		return
	}
	defer tx.Rollback(context.Background())

	// Используем прямой SQL запрос
	_, err = tx.Exec(context.Background(),
		"UPDATE schema.Products SET price = $1 WHERE id_product_main = $2",
		input.Price, input.ID)

	if err != nil {
		log.Printf("Ошибка обновления цены продукта: %v", err)
		respondError(w, http.StatusInternalServerError, "Ошибка обновления цены продукта")
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Цена продукта успешно обновлена"})
}
