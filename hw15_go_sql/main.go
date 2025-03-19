package main

import (
	"context"
	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/db"
	"log"
	"time"

	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/client"
	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Подключаемся к ОСНОВНОЙ базе данных.
	connectionString := "postgres://postgres:qwerty123@localhost:5433/postgres?sslmode=disable"
	pool, err := db.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// Создаем менеджер.
	manager := db.NewManager(
		pool,                   // теперь pool определен
		"db_alex",              // имя базы данных
		"sqlc/hw14_db_new.sql", // путь к вашему SQL файлу
	)

	// Затем выполнить скрипт (hw14_db.sql).
	err = manager.ExecuteSQL(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Создание и запуск сервера.
	srv := server.NewServer(pool)
	go srv.Start("0.0.0.0", "8080")

	// Даем серверу время запуститься перед запуском клиента.
	time.Sleep(2 * time.Second)

	// Запуск клиента
	client.Klient()

	return nil
}
