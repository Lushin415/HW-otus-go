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
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/v1/get_user?id=2", nil)
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

	// Создаём нового пользователя
	newUser := struct {
		Name     string `json:"name_user"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Name:     "Новый Пользователь",
		Email:    "new@example.com",
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
}
