package server

import (
	"encoding/json"
	"log"
	"net/http"
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

	// Эндпоинты.
	//Users.
	mux.HandleFunc("/v1/user/create", s.CreateUserHandler)
	mux.HandleFunc("/v1/user/delete", s.DeleteUserHandler)
	mux.HandleFunc("/v1/user/spending_stats", s.GetUserSpendingStatsHandler)
	mux.HandleFunc("/v1/user/get_by_password", s.GetUsersByPasswordHandler)
	mux.HandleFunc("/v1/user/update_name", s.UpdateUserNameHandler)
	//Products.
	mux.HandleFunc("/v1/product/create", s.CreateProductHandler)
	mux.HandleFunc("/v1/product/delete_cheap", s.DeleteCheapProductsHandler)
	mux.HandleFunc("/v1/product/price_range", s.GetProductsByPriceRangeHandler)
	mux.HandleFunc("/v1/product/update_price", s.UpdateProductPriceHandler)
	//Orders.
	mux.HandleFunc("/v1/order/create", s.CreateOrderHandler)
	mux.HandleFunc("/v1/order/add_product", s.CreateOrderProductHandler)
	mux.HandleFunc("/v1/order/delete", s.DeleteOrderHandler)
	mux.HandleFunc("/v1/order/get_by_user", s.GetOrderByUserIDHandler)
	mux.HandleFunc("/v1/order/update_total", s.UpdateOrderTotalHandler)

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

// Вспомогательные функции.
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

//// Обработчик получения пользователя.
//func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodGet {
//		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
//		return
//	}
//
//	passwordParam := r.URL.Query().Get("password")
//	if passwordParam == "" {
//		respondError(w, http.StatusBadRequest, "Password пользователя не указан")
//		return
//	}
//
//	userPassword, err := strconv.Atoi(passwordParam)
//	if err != nil {
//		respondError(w, http.StatusBadRequest, "Некорректный ID пользователя")
//		return
//	}
//
//	// Используем транзакцию.
//	tx, err := s.pool.Begin(context.Background())
//	if err != nil {
//		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
//		return
//	}
//	defer tx.Rollback(context.Background())
//
//	// Создаем новый экземпляр Queries внутри транзакции.
//	qtx := s.queries.WithTx(tx)
//	user, err := qtx.GetUsersByPassword(context.Background(), string(rune(userPassword)))
//	if err != nil {
//		if err.Error() == "no rows in result set" {
//			respondError(w, http.StatusNotFound, "Пользователь не найден")
//		} else {
//			log.Printf("Ошибка БД при получении пользователя: %v", err)
//			respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
//		}
//		return
//	}
//
//	if err = tx.Commit(context.Background()); err != nil {
//		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
//		return
//	}
//
//	respondJSON(w, http.StatusOK, user)
//}
//
//// Обработчик создания пользователя.
//func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
//		return
//	}
//
//	var input struct {
//		Name     string `json:"nameUser"`
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
//
//	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
//		return
//	}
//
//	// Используем транзакцию.
//	tx, err := s.pool.Begin(context.Background())
//	if err != nil {
//		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
//		return
//	}
//	defer tx.Rollback(context.Background())
//
//	// Создаем параметры для запроса.
//	params := db.CreateUserParams{
//		NameUser: input.Name,
//		Email:    input.Email,
//		Password: input.Password,
//	}
//
//	qtx := s.queries.WithTx(tx)
//	userID, err := qtx.CreateUser(context.Background(), params)
//	if err != nil {
//		log.Printf("Ошибка создания пользователя: %v", err)
//		respondError(w, http.StatusInternalServerError, "Ошибка создания пользователя")
//		return
//	}
//
//	if err = tx.Commit(context.Background()); err != nil {
//		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
//		return
//	}
//
//	respondJSON(w, http.StatusCreated, map[string]interface{}{
//		"idUserMain": userID,
//		"message":    "Пользователь создан",
//	})
//}
//
//// Обработчик получения продуктов по ценовому диапазону.
//func (s *Server) getProductsByPriceRange(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodGet {
//		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
//		return
//	}
//
//	minStr := r.URL.Query().Get("min")
//	maxStr := r.URL.Query().Get("max")
//	if minStr == "" || maxStr == "" {
//		respondError(w, http.StatusBadRequest, "Укажите диапазон цен min и max")
//		return
//	}
//
//	minV, err := strconv.ParseFloat(minStr, 64)
//	if err != nil {
//		respondError(w, http.StatusBadRequest, "Некорректная минимальная цена")
//		return
//	}
//
//	maxV, err := strconv.ParseFloat(maxStr, 64)
//	if err != nil {
//		respondError(w, http.StatusBadRequest, "Некорректная максимальная цена")
//		return
//	}
//
//	// Используем транзакцию.
//	tx, err := s.pool.Begin(context.Background())
//	if err != nil {
//		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
//		return
//	}
//	defer tx.Rollback(context.Background())
//
//	// Используем прямой SQL.
//	rows, err := tx.Query(context.Background(),
//		"SELECT id_product_main, name_product, price FROM schema.Products WHERE price BETWEEN $1 AND $2",
//		minV, maxV)
//	if err != nil {
//		log.Printf("Ошибка получения продуктов: %v", err)
//		respondError(w, http.StatusInternalServerError, "Ошибка получения продуктов")
//		return
//	}
//	defer rows.Close()
//
//	var products []db.SchemaProduct
//	for rows.Next() {
//		var product db.SchemaProduct
//		if scanErr := rows.Scan(&product.IDProductMain, &product.NameProduct, &product.Price); scanErr != nil {
//			log.Printf("Ошибка сканирования продукта: %v", scanErr)
//			respondError(w, http.StatusInternalServerError, "Ошибка чтения данных продукта")
//			return
//		}
//		products = append(products, product)
//	}
//
//	if rowsErr := rows.Err(); rowsErr != nil {
//		log.Printf("Ошибка в курсоре: %v", rowsErr)
//		respondError(w, http.StatusInternalServerError, "Ошибка обработки результатов")
//		return
//	}
//
//	if err = tx.Commit(context.Background()); err != nil {
//		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
//		return
//	}
//
//	respondJSON(w, http.StatusOK, products)
//}
//
//// Обработчик обновления цены продукта.
//func (s *Server) updateProductPrice(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPut {
//		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
//		return
//	}
//
//	var input struct {
//		ID    int32   `json:"idProductMain"`
//		Price float64 `json:"price"`
//	}
//
//	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
//		return
//	}
//
//	// Используем транзакцию.
//	tx, err := s.pool.Begin(context.Background())
//	if err != nil {
//		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
//		return
//	}
//	defer tx.Rollback(context.Background())
//
//	// Используем прямой SQL запрос.
//	_, err = tx.Exec(context.Background(),
//		"UPDATE schema.Products SET price = $1 WHERE id_product_main = $2",
//		input.Price, input.ID)
//	if err != nil {
//		log.Printf("Ошибка обновления цены продукта: %v", err)
//		respondError(w, http.StatusInternalServerError, "Ошибка обновления цены продукта")
//		return
//	}
//
//	if err = tx.Commit(context.Background()); err != nil {
//		respondError(w, http.StatusInternalServerError, "Ошибка базы данных")
//		return
//	}
//
//	respondJSON(w, http.StatusOK, map[string]string{"message": "Цена продукта успешно обновлена"})
//}
