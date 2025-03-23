package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RunClient запускает клиентские тесты API
func RunClient() {
	fmt.Println("Запуск клиентских тестов API...")

	// Создаем пользователей
	createUser()

	// Создаем продукты
	createProducts()

	//Создаем заказы
	createOrders()

	//Создать сводную таблицу
	createOrdersProduct()

	//Удалить заказ
	deleteOrder()

	//Удалить пользователя
	deleteUser("jbanov@example.com")

	// Обновляем цену продукта
	updateProductPrice(5, 95.0)

	//Обновить пользователя, установить новое имя, по email.
	updateUserName("jukova@example.com", "Екатерина Муравьева")

	// Получаем пользователей с паролем "123"
	getUsersByPassword("123")

	//Удаляем продукты с ценой менее 50
	deleteCheapProducts(50.0)

	// Получаем продукты в ценовом диапазоне
	getProductsByPriceRange(50.0, 100.0)

	// Получаем статистику пользователей
	getUsersSpendingStats()

	//Обновить таблицу заказов
	updateOrderTotal()

	//Получаем заказы пользователя по ID
	getOrdersByUserID(2)
}

//Создаем таблицу пользователей

func createUser() {
	fmt.Println("\n--- Создание таблицы пользователей ---")

	// Данные для создания пользователя
	userData := []map[string]interface{}{
		{
			"name_user": "Иван Иванов",
			"email":     "ivan@example.com",
			"password":  "qwerty",
		},
		{
			"name_user": "Петя Смирнов",
			"email":     "petr@example.com",
			"password":  "123",
		},
		{
			"name_user": "Женя Жбанов",
			"email":     "jbanov@example.com",
			"password":  "123",
		},
		{
			"name_user": "Катя Жукова",
			"email":     "jukova@example.com",
			"password":  "456",
		},
		{
			"name_user": "Аня Гавриленко",
			"email":     "gavrilenko@example.com",
			"password":  "456",
		},
	}

	// Отправляем пользователей по одному
	for _, user := range userData {
		jsonData, err := json.Marshal(user)
		if err != nil {
			fmt.Println("Ошибка кодирования JSON:", err)
			continue
		}
		// Отправляем POST запрос
		resp, err := http.Post("http://localhost:8080/v1/user/create",
			"application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Println("Ошибка при отправке запроса:", err)
			continue
		}
		defer resp.Body.Close()

		// Читаем ответ сервера
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			continue
		}

		fmt.Println("Ответ сервера:", string(body))
	}
}

// Создаем список товаров

func createProducts() {
	fmt.Println("\n--- Заполнение таблицы товаров ---")

	// Данные для создания товаров
	productData := []map[string]interface{}{
		{
			"name_product": "Носки",
			"price":        "10.00",
		},
		{
			"name_product": "Ботинки",
			"price":        "120.00",
		},
		{
			"name_product": "Калоши",
			"price":        "50.00",
		},
		{
			"name_product": "Сандали",
			"price":        "70.00",
		},
		{
			"name_product": "Туфли",
			"price":        "90.00",
		},
	}

	// Отправляем каждый товар по отдельности
	for _, product := range productData {
		jsonData, err := json.Marshal(product)
		if err != nil {
			fmt.Println("Ошибка кодирования JSON:", err)
			continue
		}
		// Отправляем POST запрос
		resp, err := http.Post("http://localhost:8080/v1/product/create",
			"application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Println("Ошибка при отправке запроса:", err)
			continue
		}
		defer resp.Body.Close()

		// Читаем ответ
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			continue
		}

		fmt.Println("Ответ сервера:", string(body))
	}
}

func createOrders() {
	fmt.Println("\n--- Заполнение таблицы с ордерами заказов ---")

	// Данные для создания ордеров
	orderData := []map[string]interface{}{
		{
			"id_user_f":    1,
			"order_date":   "2025-02-25",
			"total_amount": "0.00",
		},
		{
			"id_user_f":    2,
			"order_date":   "2025-02-25",
			"total_amount": "0.00",
		},
		{
			"id_user_f":    3,
			"order_date":   "2025-02-25",
			"total_amount": "0.00",
		},
		{
			"id_user_f":    4,
			"order_date":   "2025-02-25",
			"total_amount": "0.00",
		},
		{
			"id_user_f":    5,
			"order_date":   "2025-02-25",
			"total_amount": "0.00",
		},
	}

	// Отправляем каждый заказ отдельно
	for _, order := range orderData {
		jsonData, err := json.Marshal(order)
		if err != nil {
			fmt.Println("Ошибка кодирования JSON:", err)
			continue
		}

		// Отправляем POST запрос
		resp, err := http.Post("http://localhost:8080/v1/order/create",
			"application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Println("Ошибка при отправке запроса:", err)
			continue
		}
		defer resp.Body.Close()

		// Читаем ответ
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			continue
		}

		fmt.Println("Ответ сервера:", string(body))
	}
}

func createOrdersProduct() {
	fmt.Println("\n--- Заполнение сводной таблицы ---")

	// Данные для создания ордеров
	orderProductData := []map[string]interface{}{
		{
			"id_order_f":   1,
			"id_product_f": 1,
			"quantity":     1,
		},
		{
			"id_order_f":   1,
			"id_product_f": 2,
			"quantity":     2,
		},
		{
			"id_order_f":   2,
			"id_product_f": 4,
			"quantity":     3,
		},
		{
			"id_order_f":   2,
			"id_product_f": 5,
			"quantity":     1,
		},
		{
			"id_order_f":   3,
			"id_product_f": 1,
			"quantity":     5,
		},
		{
			"id_order_f":   4,
			"id_product_f": 5,
			"quantity":     1,
		},
		{
			"id_order_f":   5,
			"id_product_f": 5,
			"quantity":     3,
		},
		{
			"id_order_f":   5,
			"id_product_f": 1,
			"quantity":     1,
		},
	}

	// Отправляем каждую запись отдельно
	for _, item := range orderProductData {
		jsonData, err := json.Marshal(item)
		if err != nil {
			fmt.Println("Ошибка кодирования JSON:", err)
			continue
		}

		// Отправляем POST запрос
		resp, err := http.Post("http://localhost:8080/v1/order/add_product",
			"application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Println("Ошибка при отправке запроса:", err)
			continue
		}
		defer resp.Body.Close()

		// Читаем ответ
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			continue
		}

		fmt.Println("Ответ сервера:", string(body))
	}
}

func deleteOrder() {
	fmt.Println("\n--- Удаление заказа ---")

	// Формируем URL с параметром id=1
	url := "http://localhost:8080/v1/order/delete?id=1"

	// Создаем DELETE-запрос
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Выводим ответ сервера
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

// Обновляем имя пользователя по email
func updateUserName(email string, newName string) {
	fmt.Println("\n--- Обновление имени пользователя ---")

	// Данные для обновления имени
	updateData := map[string]interface{}{
		"email":     email,
		"name_user": newName,
	}

	// Кодируем данные в JSON
	jsonData, err := json.Marshal(updateData)
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	// Создаем PUT-запрос
	req, err := http.NewRequest(http.MethodPut,
		"http://localhost:8080/v1/user/update_name",
		bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Printf("Имя пользователя с email '%s' изменено на '%s'\n", email, newName)
	fmt.Println("Ответ сервера:", string(body))
}

// Удаляем пользователя по email
func deleteUser(email string) {
	fmt.Println("\n--- Удаление пользователя ---")

	// Формируем URL с параметром email
	url := fmt.Sprintf("http://localhost:8080/v1/user/delete?email=%s", url.QueryEscape(email))

	// Создаем DELETE-запрос
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	fmt.Printf("Пользователь с email '%s' удален\n", email)
	fmt.Println("Ответ сервера:", string(body))
}

// Удаляем дешевые продукты с ценой ниже порогового значения
func deleteCheapProducts(thresholdPrice float64) {
	fmt.Println("\n--- Удаление дешевых продуктов ---")

	// Формируем URL с параметром price
	url := fmt.Sprintf("http://localhost:8080/v1/product/delete_cheap?price=%.2f", thresholdPrice)

	// Создаем DELETE-запрос
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Парсим ответ
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Ошибка при разборе ответа:", err)
		return
	}

	// Выводим сообщение о результате
	fmt.Println(response["message"])

	// Если есть информация о количестве, выводим ее
	if count, ok := response["count"].(float64); ok && count > 0 {
		fmt.Printf("Удалено продуктов: %.0f\n", count)
	}
}

// Получаем заказы пользователя по ID
func getOrdersByUserID(userID int) {
	fmt.Println("\n--- Получение заказов пользователя ---")

	// Формируем URL с параметром id_user_f
	url := fmt.Sprintf("http://localhost:8080/v1/order/get_by_user?id_user_f=%d", userID)

	// Отправляем GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Если код ответа не OK, выводим ошибку
	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]string
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			fmt.Println("Ошибка:", errorResponse["error"])
		} else {
			fmt.Println("Ошибка при получении заказов. Код ответа:", resp.StatusCode)
		}
		return
	}

	// Пытаемся распарсить JSON с заказами
	var orders []map[string]interface{}
	if err := json.Unmarshal(body, &orders); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	// Если заказов нет, выводим соответствующее сообщение
	if len(orders) == 0 {
		fmt.Printf("У пользователя с ID=%d нет заказов\n", userID)
		return
	}

	// Выводим заказы в формате таблицы
	fmt.Printf("Заказы пользователя с ID=%d:\n", userID)
	fmt.Printf("%-15s | %-15s | %s\n", "ID заказа", "Дата", "Сумма")
	fmt.Println(strings.Repeat("-", 50))

	for _, order := range orders {
		orderID := fmt.Sprintf("%.0f", order["id_order_main"].(float64))
		orderDate := order["order_date"]
		totalAmount := order["total_amount"]

		fmt.Printf("%-15s | %-15v | %v\n", orderID, orderDate, totalAmount)
	}
}

// Обновляем итоговую сумму заказа
func updateOrderTotal() {
	fmt.Println("\n--- Обновление итоговых сумм заказов ---")

	// Создаем PUT-запрос без тела, так как обработчик не ожидает параметров
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/v1/order/update_total", nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	// Устанавливаем заголовок Content-Type, хотя тело запроса пустое
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка при обновлении итоговых сумм. Код ответа:", resp.StatusCode)
		fmt.Println("Сообщение:", string(body))
		return
	}

	fmt.Println("Итоговые суммы заказов успешно обновлены")
	fmt.Println("Ответ сервера:", string(body))
}
