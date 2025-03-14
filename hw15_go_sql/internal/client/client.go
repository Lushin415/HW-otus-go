package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Структуры для работы с данными
type User struct {
	ID       int    `json:"id_user_main"`
	Name     string `json:"name_user"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type NewUser struct {
	Name     string `json:"name_user"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Product struct {
	ID    int     `json:"id_product_main"`
	Name  string  `json:"name_product"`
	Price float64 `json:"price"`
}

type ProductPriceUpdate struct {
	ID    int     `json:"id_product_main"`
	Price float64 `json:"price"`
}

func Klient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Просто вызываем вспомогательные функции по порядку
	RunGetUser(ctx)
	RunCreateUser(ctx)
	RunGetAndDisplayProducts(ctx)
}

// RunGetUser - получение и отображение пользователя
func RunGetUser(ctx context.Context) {
	user, err := GetUser(ctx, 1)
	if err != nil {
		fmt.Println("Ошибка при получении пользователя:", err)
		return
	}
	fmt.Printf("Полученный пользователь: %+v\n", user)
}

// RunCreateUser - создание нового пользователя
func RunCreateUser(ctx context.Context) {
	newUserEmail := fmt.Sprintf("new_%d@example.com", time.Now().Unix())
	err := CreateUser(ctx, "Новичок_Пользователь", newUserEmail, "123456")
	if err != nil {
		fmt.Println("Ошибка при создании пользователя:", err)
		return
	}
}

// RunGetAndDisplayProducts - получение, отображение продуктов и обновление цены
func RunGetAndDisplayProducts(ctx context.Context) {
	products, err := GetProductsByPriceRange(ctx, 50, 100)
	if err != nil {
		fmt.Println("Ошибка при получении продуктов:", err)
		return
	}

	fmt.Println("\nПродукты в ценовом диапазоне 50-100:")
	for _, p := range products {
		fmt.Printf("  ID: %d, Название: %s, Цена: %.2f\n", p.ID, p.Name, p.Price)
	}

	// Обновление цены продукта, если в списке есть хотя бы один продукт
	if len(products) > 0 {
		product := products[0]
		newPrice := product.Price + 10.0

		err = UpdateProductPrice(ctx, product.ID, newPrice)
		if err != nil {
			fmt.Println("Ошибка при обновлении цены продукта:", err)
			return
		}

		fmt.Printf("\nЦена продукта '%s' (ID: %d) обновлена с %.2f на %.2f\n",
			product.Name, product.ID, product.Price, newPrice)
	}
}

// GetUser - получение информации о пользователе по ID
func GetUser(ctx context.Context, userID int) (*User, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/get_user?id=%d", userID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка получения пользователя: код %d, ответ: %s",
			resp.StatusCode, string(body))
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	return &user, nil
}

// CreateUser - создание нового пользователя
func CreateUser(ctx context.Context, name, email, password string) error {
	newUser := NewUser{
		Name:     name,
		Email:    email,
		Password: password,
	}

	userData, err := json.Marshal(newUser)
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"http://localhost:8080/v1/create_user", bytes.NewBuffer(userData))
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("ошибка создания пользователя: код %d, ответ: %s",
			resp.StatusCode, string(body))
	}

	fmt.Println("Ответ сервера:", string(body))
	return nil
}

// GetProductsByPriceRange - получение продуктов в заданном ценовом диапазоне
func GetProductsByPriceRange(ctx context.Context, minPrice, maxPrice float64) ([]Product, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/products/price_range?min=%.2f&max=%.2f",
		minPrice, maxPrice)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка получения продуктов: код %d, ответ: %s",
			resp.StatusCode, string(body))
	}

	var products []Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	return products, nil
}

// UpdateProductPrice - обновление цены продукта
func UpdateProductPrice(ctx context.Context, productID int, newPrice float64) error {
	updateData := ProductPriceUpdate{
		ID:    productID,
		Price: newPrice,
	}

	updateJSON, err := json.Marshal(updateData)
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut,
		"http://localhost:8080/v1/update_product_price", bytes.NewBuffer(updateJSON))
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка обновления цены: код %d, ответ: %s",
			resp.StatusCode, string(body))
	}

	fmt.Println("Ответ сервера:", string(body))
	return nil
}
