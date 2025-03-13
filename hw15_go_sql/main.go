package main

import (
	"context"
	"github.com/fixme_my_friend/hw15_go_sql/internal/client"
	"log"
	"os"
	"time"

	"github.com/fixme_my_friend/hw15_go_sql/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Получение параметров подключения к БД из переменных окружения или конфига
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:qwerty123@localhost:5433/postgres"
	}

	// Установка соединения с базой данных
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Ошибка парсинга URL базы данных: %v", err)
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer dbPool.Close()

	// Проверка соединения
	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatalf("Ошибка проверки соединения с базой данных: %v", err)
	}

	// Создание и запуск сервера
	srv := server.NewServer(dbPool)
	go srv.Start("0.0.0.0", "8080")

	// Даем серверу время запуститься перед запуском клиента
	time.Sleep(2 * time.Second)

	// Запуск клиента
	client.Klient()
}
