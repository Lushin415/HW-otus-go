package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Klient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/v1/get_user", nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	fmt.Println(string(body))

	user := User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Ошибка десериализации", err)
		return
	}

	fmt.Printf("Пользователь: %+v\n", user)

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8080/v1/create_user",
		strings.NewReader(user.String()))
	if err != nil {
		fmt.Println("Ошибка создания запроса", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return
	}

	fmt.Println("Пользователь создан успешно")

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	var newUser User

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		fmt.Println("Ошибка десериализации", err)
		return
	}

	fmt.Printf("Пользователь: %+v\n", newUser)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) String() string {
	body, _ := json.Marshal(u)

	return string(body)
}
