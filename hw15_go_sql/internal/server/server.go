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
	// Users.
	mux.HandleFunc("/v1/user/create", s.CreateUserHandler)
	mux.HandleFunc("/v1/user/delete", s.DeleteUserHandler)
	mux.HandleFunc("/v1/user/spending_stats", s.GetUserSpendingStatsHandler)
	mux.HandleFunc("/v1/user/get_by_password", s.GetUsersByPasswordHandler)
	mux.HandleFunc("/v1/user/update_name", s.UpdateUserNameHandler)
	// Products.
	mux.HandleFunc("/v1/product/create", s.CreateProductHandler)
	mux.HandleFunc("/v1/product/delete_cheap", s.DeleteCheapProductsHandler)
	mux.HandleFunc("/v1/product/price_range", s.GetProductsByPriceRangeHandler)
	mux.HandleFunc("/v1/product/update_price", s.UpdateProductPriceHandler)
	// Orders.
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
