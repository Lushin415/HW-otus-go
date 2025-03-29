package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/httputils"
)

// RunClient запускает клиентские тесты API.
func RunClient() {
	fmt.Println("Запуск клиентских тестов API...")

	// Создаем пользователей
	createUser()

	// Создаем продукты
	createProducts()

	// Создаем заказы
	createOrders()

	// Создать сводную таблицу
	createOrdersProduct()

	// Удалить заказ
	deleteOrder()

	// Удалить пользователя
	deleteUser("jbanov@example.com")

	// Обновляем цену продукта
	updateProductPrice(5, 95.0)

	// Обновить пользователя, установить новое имя, по email.
	updateUserName("jukova@example.com", "Екатерина Муравьева")

	// Получаем пользователей с паролем "123"
	getUsersByPassword("123")

	// Удаляем продукты с ценой менее 50
	deleteCheapProducts(50.0)

	// Получаем продукты в ценовом диапазоне
	getProductsByPriceRange(50.0, 100.0)

	// Получаем статистику пользователей
	getUsersSpendingStats()

	// Обновить таблицу заказов
	updateOrderTotal()

	// Получаем заказы пользователя по ID
	getOrdersByUserID(2)
}

// Создаем таблицу пользователей

func createUser() {
	fmt.Println("\n--- Создание таблицы пользователей ---")

	// Данные для создания пользователя
	userData := []map[string]interface{}{
		{
			"nameUser": "Иван Иванов",
			"email":    "ivan@example.com",
			"password": "qwerty",
		},
		{
			"nameUser": "Петя Смирнов",
			"email":    "petr@example.com",
			"password": "123",
		},
		{
			"nameUser": "Женя Жбанов",
			"email":    "jbanov@example.com",
			"password": "123",
		},
		{
			"nameUser": "Катя Жукова",
			"email":    "jukova@example.com",
			"password": "456",
		},
		{
			"nameUser": "Аня Гавриленко",
			"email":    "gavrilenko@example.com",
			"password": "456",
		},
	}

	// Отправляем пользователей по одному
	for _, user := range userData {
		if err := httputils.SendData("http://localhost:8080/v1/user/create", user); err != nil {
			fmt.Println("Ошибка при отправке данных пользователя:", err)
		}
	}
}

// Создаем список товаров

func createProducts() {
	fmt.Println("\n--- Заполнение таблицы товаров ---")

	// Данные для создания товаров
	productData := []map[string]interface{}{
		{
			"nameProduct": "Носки",
			"price":       "10.00",
		},
		{
			"nameProduct": "Ботинки",
			"price":       "120.00",
		},
		{
			"nameProduct": "Калоши",
			"price":       "50.00",
		},
		{
			"nameProduct": "Сандали",
			"price":       "70.00",
		},
		{
			"nameProduct": "Туфли",
			"price":       "90.00",
		},
	}

	// Отправляем каждый товар по отдельности
	for _, product := range productData {
		if err := httputils.SendData("http://localhost:8080/v1/product/create", product); err != nil {
			fmt.Println("Ошибка при отправке данных заказа:", err)
		}
	}
}

func createOrders() {
	fmt.Println("\n--- Заполнение таблицы с ордерами заказов ---")

	// Данные для создания ордеров
	orderData := []map[string]interface{}{
		{
			"idUserF":     1,
			"orderDate":   "2025-02-25",
			"totalAmount": "0.00",
		},
		{
			"idUserF":     2,
			"orderDate":   "2025-02-25",
			"totalAmount": "0.00",
		},
		{
			"idUserF":     3,
			"orderDate":   "2025-02-25",
			"totalAmount": "0.00",
		},
		{
			"idUserF":     4,
			"orderDate":   "2025-02-25",
			"totalAmount": "0.00",
		},
		{
			"idUserF":     5,
			"orderDate":   "2025-02-25",
			"totalAmount": "0.00",
		},
	}

	for _, order := range orderData {
		if err := httputils.SendData("http://localhost:8080/v1/order/create", order); err != nil {
			fmt.Println("Ошибка при отправке данных заказа:", err)
		}
	}
}

func createOrdersProduct() {
	fmt.Println("\n--- Заполнение сводной таблицы ---")

	// Данные для создания ордеров
	orderProductData := []map[string]interface{}{
		{
			"idOrderF":   1,
			"idProductF": 1,
			"quantity":   1,
		},
		{
			"idOrderF":   1,
			"idProductF": 2,
			"quantity":   2,
		},
		{
			"idOrderF":   2,
			"idProductF": 4,
			"quantity":   3,
		},
		{
			"idOrderF":   2,
			"idProductF": 5,
			"quantity":   1,
		},
		{
			"idOrderF":   3,
			"idProductF": 1,
			"quantity":   5,
		},
		{
			"idOrderF":   4,
			"idProductF": 5,
			"quantity":   1,
		},
		{
			"idOrderF":   5,
			"idProductF": 5,
			"quantity":   3,
		},
		{
			"idOrderF":   5,
			"idProductF": 1,
			"quantity":   1,
		},
	}

	// Отправляем каждую запись отдельно
	for _, item := range orderProductData {
		if err := httputils.SendData("http://localhost:8080/v1/order/add_product", item); err != nil {
			fmt.Println("Ошибка при отправке данных товара в заказ:", err)
		}
	}
}

func deleteOrder() {
	fmt.Println("\n--- Удаление заказа ---")

	body, err := httputils.ReadResponse(http.MethodDelete, "http://localhost:8080/v1/order/delete?id=1", nil)
	if err != nil {
		fmt.Println("Ошибка при удалении заказа:", err)
		return
	}

	// Выводим ответ сервера
	fmt.Println("Ответ сервера:", string(body))
}

// Получаем пользователей по паролю.
func getUsersByPassword(password string) {
	fmt.Println("\n--- Получение пользователей с паролем ---")

	// Формируем URL с параметром password
	url := fmt.Sprintf("http://localhost:8080/v1/user/get_by_password?password=%s", password)

	body, err := httputils.ReadResponse(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Ошибка при получении пользователей:", err)
		return
	}

	// Выводим ответ сервера
	fmt.Println("Результат:", string(body))
}

// Получаем продукты в ценовом диапазоне.
func getProductsByPriceRange(minPrice, maxPrice float64) {
	fmt.Println("\n--- Получение продуктов по ценовому диапазону ---")

	// Формируем URL с параметрами
	url := fmt.Sprintf("http://localhost:8080/v1/product/price_range?minPrice=%.2f&maxPrice=%.2f",
		minPrice, maxPrice)
	// #nosec G107 - URL строится только из доверенных компонентов.

	// Используем readResponse вместо прямого вызова http_GET
	body, err := httputils.ReadResponse(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Ошибка при получении продуктов:", err)
		return
	}

	// Выводим результат
	fmt.Println("Продукты в диапазоне:", string(body))
}

// Обновляем цену продукта.
func updateProductPrice(productID int, newPrice float64) {
	fmt.Println("\n--- Обновление цены продукта ---")

	// Данные для обновления цены
	updateData := map[string]interface{}{
		"idProductMain": productID,
		"price":         fmt.Sprintf("%.2f", newPrice),
	}

	// Отправляем PUT-запрос
	body, err := httputils.ReadResponse(http.MethodPut,
		"http://localhost:8080/v1/product/update_price",
		updateData)
	if err != nil {
		fmt.Println("Ошибка при обновлении цены продукта:", err)
		return
	}

	// Выводим результат
	fmt.Println("Результат обновления:", string(body))
}

// Получаем статистику пользователей.
func getUsersSpendingStats() {
	fmt.Println("\n--- Получение статистики пользователей ---")

	// GET-запрос
	body, err := httputils.ReadResponse(http.MethodGet, "http://localhost:8080/v1/user/spending_stats", nil)
	if err != nil {
		fmt.Println("Ошибка при получении статистики пользователей:", err)
		return
	}

	// Выводим результат
	fmt.Println("Статистика пользователей:", string(body))
}

// Обновляем имя пользователя по email.
func updateUserName(email string, newName string) {
	fmt.Println("\n--- Обновление имени пользователя ---")

	// Данные для обновления имени
	updateData := map[string]interface{}{
		"email":    email,
		"nameUser": newName,
	}

	// Используем readResponse для отправки PUT-запроса с данными
	body, err := httputils.ReadResponse(http.MethodPut, "http://localhost:8080/v1/user/update_name", updateData)
	if err != nil {
		fmt.Println("Ошибка при обновлении имени пользователя:", err)
		return
	}

	fmt.Printf("Имя пользователя с email '%s' изменено на '%s'\n", email, newName)
	fmt.Println("Ответ сервера:", string(body))
}

// Удаляем пользователя по email.
func deleteUser(email string) {
	fmt.Println("\n--- Удаление пользователя ---")

	// Формируем URL с параметром email
	url := fmt.Sprintf("http://localhost:8080/v1/user/delete?email=%s", url.QueryEscape(email))

	// Используем readResponse для отправки DELETE-запроса
	body, err := httputils.ReadResponse(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Ошибка при удалении пользователя:", err)
		return
	}

	fmt.Printf("Пользователь с email '%s' удален\n", email)
	fmt.Println("Ответ сервера:", string(body))
}

// Удаляем дешевые продукты с ценой ниже порогового значения.
func deleteCheapProducts(thresholdPrice float64) {
	fmt.Println("\n--- Удаление дешевых продуктов ---")

	// Формируем URL с параметром price
	url := fmt.Sprintf("http://localhost:8080/v1/product/delete_cheap?price=%.2f", thresholdPrice)

	// Используем readResponse для отправки DELETE-запроса
	body, err := httputils.ReadResponse(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Ошибка при удалении дешевых продуктов:", err)
		return
	}

	// Парсинг ответа
	var response map[string]interface{}
	if err = json.Unmarshal(body, &response); err != nil {
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

// Получаем заказы пользователя по ID.
func getOrdersByUserID(userID int) {
	fmt.Println("\n--- Получение заказов пользователя ---")

	// Формируем URL с параметром idUserF
	url := fmt.Sprintf("http://localhost:8080/v1/order/get_by_user?idUserF=%d", userID)
	// #nosec G107 - URL строится только из доверенных компонентов.

	// Получаем ответ через sendJSONRequest, чтобы иметь доступ к статус-коду
	resp, err := httputils.SendJSONRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Ошибка при получении заказов пользователя:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Если код ответа не OK, выводим ошибку
	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]string
		if err = json.Unmarshal(body, &errorResponse); err == nil {
			fmt.Println("Ошибка:", errorResponse["error"])
		} else {
			fmt.Println("Ошибка при получении заказов. Код ответа:", resp.StatusCode)
		}
		return
	}

	// Парсинг JSON с заказами
	var orders []map[string]interface{}
	if err = json.Unmarshal(body, &orders); err != nil {
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
		orderID := fmt.Sprintf("%.0f", order["idOrderMain"].(float64))
		orderDate := order["orderDate"]
		totalAmount := order["totalAmount"]

		fmt.Printf("%-15s | %-15v | %v\n", orderID, orderDate, totalAmount)
	}
}

// Обновляем итоговую сумму заказа.
func updateOrderTotal() {
	fmt.Println("\n--- Обновление итоговых сумм заказов ---")

	// Используем sendJSONRequest для отправки PUT-запроса без тела
	resp, err := httputils.SendJSONRequest(http.MethodPut, "http://localhost:8080/v1/order/update_total", nil)
	if err != nil {
		fmt.Println("Ошибка при обновлении итоговых сумм заказов:", err)
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
