package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Klient() {
	resp, err := http.Get(
		"http://localhost:8080/v1/get_user",
	)
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

	fmt.Printf("User: %+v\n", user)

	resp, err = http.Post(
		"http://localhost:8080/v1/create_user",
		"application/json",
		strings.NewReader(user.String()),
	)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return
	}

	fmt.Println("User created successfully")

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

	fmt.Printf("User: %+v\n", newUser)
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
