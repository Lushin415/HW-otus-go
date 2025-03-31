package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/client"
	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/db"
	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Подключаемся к ОСНОВНОЙ базе данных.
	connectionString := "postgres://postgres:qwerty123@postgres-chat:5432/postgres?sslmode=disable"
	mainPool, err := db.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Проверяем существование базы DB_Alex
	var exists bool
	err = mainPool.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = 'db_alex')").Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем базу, если она не существует
	if !exists {
		_, err = mainPool.Exec(context.Background(), "CREATE DATABASE db_alex")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("База данных db_alex создана")
	}
	mainPool.Close()

	// Подключаемся к базе db_alex
	dbConnection := "postgres://postgres:qwerty123@postgres-chat:5432/db_alex?sslmode=disable"
	pool, err := db.Connect(context.Background(), dbConnection)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем менеджер.
	manager := db.NewManager(
		pool,                   // теперь pool определен
		"db_alex",              // имя базы данных
		"sqlc/hw14_db_new.sql", // путь к SQL файлу
	)

	// Затем выполнить скрипт (hw14_db_new.sql).
	err = manager.ExecuteSQL(context.Background())
	if err != nil {
		log.Printf("Ошибка: %v", err)
		os.Exit(1)
	}
	defer pool.Close()
	// Создание и запуск сервера.
	srv := server.NewServer(pool)
	go srv.Start("0.0.0.0", "8080")

	// Даем серверу время запуститься перед запуском клиента.
	time.Sleep(2 * time.Second)

	// Запуск клиента
	client.RunClient()

	return nil
}
