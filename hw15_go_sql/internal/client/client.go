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

func Klient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Получаем пользователя с ID=2 (т.к. пользователь с ID=1 удален в скрипте)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/v1/get_user?id=1", nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса на получение пользователя:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса на получение пользователя:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка получения пользователя: код состояния %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Ответ сервера:", string(body))
		return
	}

	// Структура соответствует полям в таблице Users
	var user struct {
		ID       int    `json:"id_user_main"`
		Name     string `json:"name_user"`
		Email    string `json:"email"`
		Password string `json:"password,omitempty"` // omitempty, чтобы не отображать, если сервер не возвращает пароль
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return
	}
	fmt.Printf("Полученный пользователь: %+v\n", user)

	// Создаём нового пользователя с уникальным email, используя текущее время
	currentTime := time.Now().Unix()
	newUser := struct {
		Name     string `json:"name_user"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Name:     "Новичок_Пользователь",
		Email:    fmt.Sprintf("new_%d@example.com", currentTime),
		Password: "123456",
	}

	userData, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println("Ошибка сериализации JSON:", err)
		return
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8080/v1/create_user", bytes.NewBuffer(userData))
	if err != nil {
		fmt.Println("Ошибка создания запроса на создание пользователя:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса на создание пользователя:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Ошибка создания пользователя: код состояния %d\n", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}
	fmt.Println("Ответ сервера:", string(body))

	// Добавим запрос на получение продуктов в ценовом диапазоне
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/v1/products/price_range?min=50&max=100", nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса на получение продуктов:", err)
		return
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса на получение продуктов:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка получения продуктов: код состояния %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Ответ сервера:", string(body))
		return
	}

	var products []struct {
		ID    int     `json:"id_product_main"`
		Name  string  `json:"name_product"`
		Price float64 `json:"price"`
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return
	}

	fmt.Println("\nПродукты в ценовом диапазоне 50-100:")
	for _, p := range products {
		fmt.Printf("  ID: %d, Название: %s, Цена: %.2f\n", p.ID, p.Name, p.Price)
	}

	// Обновление цены продукта
	if len(products) > 0 {
		productToUpdate := products[0]
		newPrice := productToUpdate.Price + 10.0

		updateData := struct {
			ID    int     `json:"id_product_main"`
			Price float64 `json:"price"`
		}{
			ID:    productToUpdate.ID,
			Price: newPrice,
		}

		updateJSON, err := json.Marshal(updateData)
		if err != nil {
			fmt.Println("Ошибка сериализации JSON:", err)
			return
		}

		req, err = http.NewRequestWithContext(ctx, http.MethodPut, "http://localhost:8080/v1/update_product_price", bytes.NewBuffer(updateJSON))
		if err != nil {
			fmt.Println("Ошибка создания запроса на обновление цены:", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Ошибка запроса на обновление цены:", err)
			return
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Ошибка обновления цены: код состояния %d\n", resp.StatusCode)
			fmt.Println("Ответ сервера:", string(body))
		} else {
			fmt.Printf("\nЦена продукта '%s' (ID: %d) обновлена с %.2f на %.2f\n",
				productToUpdate.Name, productToUpdate.ID, productToUpdate.Price, newPrice)
			fmt.Println("Ответ сервера:", string(body))
		}
	}
}
