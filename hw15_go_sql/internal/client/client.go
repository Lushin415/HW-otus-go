package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// RunClient запускает клиентские тесты API
func RunClient() {
	fmt.Println("Запуск клиентских тестов API...")

	// Создаем пользователя
	createUser()

	// Получаем пользователей с паролем "123"
	getUsersByPassword("123")

	// Получаем продукты в ценовом диапазоне
	getProductsByPriceRange(50.0, 100.0)

	// Обновляем цену продукта
	updateProductPrice(5, 95.0)

	// Получаем статистику пользователей
	getUsersSpendingStats()
}

// Создаем нового пользователя
func createUser() {
	fmt.Println("\n--- Создание нового пользователя ---")

	// Данные для создания пользователя
	userData := map[string]string{
		"name_user": "Новый Пользователь",
		"email":     fmt.Sprintf("user_%d@example.com", time.Now().Unix()),
		"password":  "simple123",
	}

	// Кодируем данные в JSON
	jsonData, _ := json.Marshal(userData)

	// Отправляем POST-запрос
	resp, err := http.Post("http://localhost:8080/v1/user/create",
		"application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Ответ сервера:", string(body))
}

// Получаем пользователей по паролю
func getUsersByPassword(password string) {
	fmt.Println("\n--- Получение пользователей с паролем ---")

	// Отправляем GET-запрос
	resp, err := http.Get("http://localhost:8080/v1/user/get_by_password?password=" + password)

	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Результат:", string(body))
}

// Получаем продукты в ценовом диапазоне
func getProductsByPriceRange(min, max float64) {
	fmt.Println("\n--- Получение продуктов по ценовому диапазону ---")

	// Формируем URL с параметрами
	url := fmt.Sprintf("http://localhost:8080/v1/product/price_range?min=%.2f&max=%.2f", min, max)

	// Отправляем GET-запрос
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Продукты в диапазоне:", string(body))
}

// Обновляем цену продукта
func updateProductPrice(productID int, newPrice float64) {
	fmt.Println("\n--- Обновление цены продукта ---")

	// Данные для обновления цены
	updateData := map[string]interface{}{
		"id_product_main": productID,
		"price":           fmt.Sprintf("%.2f", newPrice),
	}

	// Кодируем данные в JSON
	jsonData, _ := json.Marshal(updateData)

	// Создаем PUT-запрос
	req, _ := http.NewRequest(http.MethodPut,
		"http://localhost:8080/v1/product/update_price",
		bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Результат обновления:", string(body))
}

// Получаем статистику пользователей
func getUsersSpendingStats() {
	fmt.Println("\n--- Получение статистики пользователей ---")

	// Отправляем GET-запрос
	resp, err := http.Get("http://localhost:8080/v1/user/spending_stats")

	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Статистика пользователей:", string(body))
}
