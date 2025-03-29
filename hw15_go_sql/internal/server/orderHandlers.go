package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// CreateOrderHandler – обработчик создания заказа.
// Ожидается POST-запрос с JSON, содержащим id_user_f, order_date и total_amount.
// order_date и total_amount приходят в виде строк, которые мы преобразуем в pgtype.
func (s *Server) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	// Данные из запроса
	var input struct {
		IDUserF     int32  `json:"idUserF"`
		OrderDate   string `json:"orderDate"`   // Ожидается формат "YYYY-MM-DD"
		TotalAmount string `json:"totalAmount"` // Строка, которую можно установить в pgtype.Numeric
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
		return
	}

	// Преобразуем order_date в pgtype.Date
	var orderDate pgtype.Date
	if err := orderDate.Scan(input.OrderDate); err != nil {
		respondError(w, http.StatusBadRequest, "Неверный формат даты")
		return
	}

	// Преобразуем total_amount в pgtype.Numeric
	var totalAmount pgtype.Numeric
	if err := totalAmount.Scan(input.TotalAmount); err != nil {
		respondError(w, http.StatusBadRequest, "Неверный формат суммы")
		return
	}

	// Формируем параметры для sqlc-функции CreateOrder
	params := db.CreateOrderParams{
		IDUserF:     input.IDUserF,
		OrderDate:   orderDate,
		TotalAmount: totalAmount,
	}

	orderID, err := s.queries.CreateOrder(context.Background(), params)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка создания заказа")
		return
	}

	respondJSON(w, http.StatusCreated, map[string]interface{}{
		"idOrderMain": orderID,
		"message":     "Заказ создан",
	})
}

// CreateOrderProductHandler – обработчик для вставки записи о продукте в заказ.
// Ожидается POST-запрос с JSON, содержащим id_order_f, id_product_f и quantity.
func (s *Server) CreateOrderProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	var input db.CreateOrderProductParams
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
		return
	}

	if err := s.queries.CreateOrderProduct(context.Background(), input); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка добавления продукта в заказ")
		return
	}

	respondJSON(w, http.StatusCreated, map[string]string{"message": "Продукт добавлен в заказ"})
}

// DeleteOrderHandler – обработчик для удаления заказа.
// Ожидается DELETE-запрос с параметром id (id_order_main) в URL.
func (s *Server) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	// Получаем id заказа из параметров URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		respondError(w, http.StatusBadRequest, "ID заказа не указан")
		return
	}
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Некорректный ID заказа")
		return
	}

	if err := s.queries.DeleteOrder(context.Background(), int32(orderID)); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка удаления заказа")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Заказ удалён"})
}

// GetOrderByUserIDHandler – обработчик для получения заказов по ID пользователя.
// Ожидается GET-запрос с параметром id_user_f в URL.
func (s *Server) GetOrderByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	idStr := r.URL.Query().Get("idUserF")
	if idStr == "" {
		respondError(w, http.StatusBadRequest, "ID пользователя не указан")
		return
	}
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Некорректный ID пользователя")
		return
	}

	orders, err := s.queries.GetOrderByUserID(context.Background(), int32(userID))
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка получения заказов")
		return
	}

	respondJSON(w, http.StatusOK, orders)
}

// UpdateOrderTotalHandler – обработчик для обновления итоговой суммы заказа.
// Ожидается PUT-запрос. Обычно вызов этой функции происходит после добавления или изменения продуктов в заказе.
func (s *Server) UpdateOrderTotalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	if err := s.queries.UpdateOrderTotal(context.Background()); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка обновления итоговой суммы заказа")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Итоговая сумма заказа обновлена"})
}
