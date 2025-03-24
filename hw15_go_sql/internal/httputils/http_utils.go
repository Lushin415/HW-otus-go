package httputils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SendData отправляет данные на указанный URL.
func SendData(url string, data map[string]interface{}) error {
	body, err := ReadResponse(http.MethodPost, url, data)
	if err != nil {
		return err
	}

	fmt.Println("Ответ сервера:", string(body))
	return nil
}

// SendJSONRequest отправляет JSON-запрос с контекстом и возвращает ответ.
func SendJSONRequest(method, url string, data interface{}) (*http.Response, error) {
	// Создаем контекст с тайм-аутом
	// #nosec G107 - URL задается только внутри кода и не получается от пользователя.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("ошибка кодирования JSON: %w", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	// Создаем запрос с контекстом
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при отправке запроса: %w", err)
	}

	return resp, nil
}

// ReadResponse отправляет JSON-запрос и возвращает тело ответа.
func ReadResponse(method, url string, data interface{}) ([]byte, error) {
	resp, err := SendJSONRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	return body, nil
}
