package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// CreateProductHandler – обработчик для создания продукта.
// Ожидается POST-запрос с JSON, содержащим поля name_product и price (в виде строки, например, "120.00").
func (s *Server) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса.
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	// Декодируем JSON-тело запроса.
	var input struct {
		NameProduct string `json:"nameProduct"`
		Price       string `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
		return
	}

	// Преобразуем цену из строки в pgtype.Numeric.
	var price pgtype.Numeric
	if err := price.Scan(input.Price); err != nil {
		respondError(w, http.StatusBadRequest, "Неверный формат цены")
		return
	}

	// Формируем параметры для sqlc-функции CreateProduct.
	params := db.CreateProductParams{
		NameProduct: input.NameProduct,
		Price:       price,
	}

	// Вызываем функцию для создания продукта.
	productID, err := s.queries.CreateProduct(context.Background(), params)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка создания продукта")
		return
	}

	respondJSON(w, http.StatusCreated, map[string]interface{}{
		"idProductMain": productID,
		"message":       "Продукт создан",
	})
}

// DeleteCheapProductsHandler – обработчик для удаления продуктов с ценой меньше заданной.
// Ожидается DELETE-запрос с query-параметром price (например, ?price=50.00).
func (s *Server) DeleteCheapProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	priceStr := r.URL.Query().Get("price")
	if priceStr == "" {
		respondError(w, http.StatusBadRequest, "Цена не указана")
		return
	}

	// Преобразуем цену из строки в pgtype.Numeric.
	var threshold pgtype.Numeric
	if err := threshold.Scan(priceStr); err != nil {
		respondError(w, http.StatusBadRequest, "Неверный формат цены")
		return
	}

	// Предполагаем, что функция DeleteCheapProducts теперь возвращает количество удаленных строк
	count, err := s.queries.DeleteCheapProducts(context.Background(), threshold)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка удаления дешевых продуктов")
		return
	}

	if count == 0 {
		// Если ничего не удалено, отправляем сообщение "не найдено"
		respondJSON(w, http.StatusOK, map[string]interface{}{
			"message": "Продукты с ценой ниже указанного порога не найдены",
			"count":   0,
		})
		return
	}

	// Если удаление прошло успешно, отправляем информацию о количестве удаленных продуктов
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Дешевые продукты удалены",
		"count":   count,
	})
}

// GetProductsByPriceRangeHandler – обработчик для получения продуктов в заданном ценовом диапазоне.
// Ожидается GET-запрос с query-параметрами min и max (например, ?min=50.00&max=150.00).
func (s *Server) GetProductsByPriceRangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	minStr := r.URL.Query().Get("minPrice")
	maxStr := r.URL.Query().Get("maxPrice")
	if minStr == "" || maxStr == "" {
		respondError(w, http.StatusBadRequest, "Укажите min и max")
		return
	}

	// Преобразуем min и max в pgtype.Numeric.
	var minPrice, maxPrice pgtype.Numeric
	if err := minPrice.Scan(minStr); err != nil {
		respondError(w, http.StatusBadRequest, "Неверный формат min цены")
		return
	}
	if err := maxPrice.Scan(maxStr); err != nil {
		respondError(w, http.StatusBadRequest, "Неверный формат max цены")
		return
	}

	// Формируем параметры для sqlc-функции GetProductsByPriceRange.
	params := db.GetProductsByPriceRangeParams{
		Price:   minPrice,
		Price_2: maxPrice,
	}

	// Вызываем функцию для получения продуктов.
	products, err := s.queries.GetProductsByPriceRange(context.Background(), params)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка получения продуктов")
		return
	}

	respondJSON(w, http.StatusOK, products)
}

// UpdateProductPriceHandler – обработчик для обновления цены продукта.
// Ожидается PUT-запрос с JSON, содержащим id_product_main и новое значение price (в виде строки).
func (s *Server) UpdateProductPriceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	var input struct {
		IDProductMain int32  `json:"idProductMain"`
		Price         string `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "Ошибка обработки JSON")
		return
	}

	// Преобразуем новое значение цены из строки в pgtype.Numeric.
	var newPrice pgtype.Numeric
	if err := newPrice.Scan(input.Price); err != nil {
		respondError(w, http.StatusBadRequest, "Неверный формат цены")
		return
	}

	// Формируем параметры для sqlc-функции UpdateProductPrice.
	params := db.UpdateProductPriceParams{
		IDProductMain: input.IDProductMain,
		Price:         newPrice,
	}

	// Вызываем функцию обновления цены.
	if err := s.queries.UpdateProductPrice(context.Background(), params); err != nil {
		respondError(w, http.StatusInternalServerError, "Ошибка обновления цены продукта")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Цена продукта обновлена"})
}
