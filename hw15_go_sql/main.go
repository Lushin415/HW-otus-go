package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/client"
	"github.com/Lushin415/HW-otus-go/hw15_go_sql/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Определение имени базы данных (в нижнем регистре!)
	dbName := "db_alex"

	// Подключение к базе postgres для создания нашей базы
	mainURL := "postgres://postgres:qwerty123@localhost:5433/postgres?sslmode=disable"
	mainPool, err := pgxpool.New(context.Background(), mainURL)
	if err != nil {
		return fmt.Errorf("ошибка подключения к основной базе данных: %w", err)
	}
	defer mainPool.Close()

	log.Printf("Подключение к основной БД успешно")

	// Проверяем существование нашей базы данных (запрос в нижнем регистре)
	var exists bool
	err = mainPool.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE lower(datname) = lower($1))", dbName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("ошибка проверки существования БД: %w", err)
	}

	// Если база не существует, создаем её
	if !exists {
		log.Printf("База данных %s не существует, создаем...", dbName)
		_, err = mainPool.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("ошибка создания базы данных: %w", err)
		}
		log.Printf("База данных %s успешно создана", dbName)
	} else {
		log.Printf("База данных %s уже существует", dbName)
	}

	// Подключение к созданной базе данных
	dbURL := fmt.Sprintf("postgres://postgres:qwerty123@localhost:5433/%s?sslmode=disable", dbName)
	log.Printf("Подключение к БД: %s", dbURL)

	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных %s: %w", dbName, err)
	}
	defer dbPool.Close()

	// Проверка соединения
	err = dbPool.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("ошибка проверки соединения с базой данных: %w", err)
	}
	log.Println("Соединение с БД установлено успешно")

	// Инициализация базы данных
	if err = initializeDatabase(dbPool); err != nil {
		return fmt.Errorf("ошибка инициализации базы данных: %w", err)
	}

	// Создание и запуск сервера
	srv := server.NewServer(dbPool)
	go srv.Start("0.0.0.0", "8080")

	// Даем серверу время запуститься перед запуском клиента
	time.Sleep(2 * time.Second)

	// Запуск клиента
	client.Klient()

	return nil
}

// Проверка существования схемы и таблиц.
func checkSchemaAndTables(dbPool *pgxpool.Pool) (tablesExist bool, err error) {
	// Проверяем существование необходимых таблиц
	err = dbPool.QueryRow(context.Background(), `
		SELECT EXISTS(
			SELECT 1 
			FROM information_schema.tables 
			WHERE table_schema = 'schema' AND table_name = 'users'
		)`).Scan(&tablesExist)
	if err != nil {
		// Проверяем, существует ли схема
		log.Println("Проверяем существование схемы 'schema'...")
		var schemaExists bool
		dbPool.QueryRow(context.Background(),
			"SELECT EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = 'schema')").Scan(&schemaExists)

		if !schemaExists {
			log.Println("Схема 'schema' не существует, создаем...")
			_, err = dbPool.Exec(context.Background(), "CREATE SCHEMA IF NOT EXISTS schema")
			if err != nil {
				return false, err
			}
		}
	}

	return tablesExist, err
}

// Выполнение SQL-скрипта.
func executeSQL(dbPool *pgxpool.Pool) error {
	// Читаем SQL-скрипт
	var sqlBytes []byte
	var err error
	sqlBytes, err = os.ReadFile("sqlc/hw14_DB.sql")
	if err != nil {
		return err
	}

	sqlScript := string(sqlBytes)

	// Разделяем скрипт на отдельные команды
	commands := strings.Split(sqlScript, ";")

	// Выполняем каждую команду отдельно
	for _, cmd := range commands {
		trimmedCmd := strings.TrimSpace(cmd)
		if trimmedCmd == "" {
			continue
		}

		_, err = dbPool.Exec(context.Background(), trimmedCmd)
		if err != nil {
			// Пропускаем ошибки создания базы данных, так как мы её уже создали
			if strings.Contains(err.Error(), "CREATE DATABASE") {
				continue
			}
			log.Printf("Предупреждение при выполнении SQL: %v\nКоманда: %s", err, trimmedCmd)
		}
	}

	return nil
}

// Инициализация базы данных с помощью SQL-скрипта.
func initializeDatabase(dbPool *pgxpool.Pool) error {
	tablesExist, err := checkSchemaAndTables(dbPool)

	// Инициализируем таблицы, если их нет или была ошибка при проверке
	if err != nil || !tablesExist {
		log.Println("Таблицы не существуют, выполняем инициализацию базы данных...")
		err = executeSQL(dbPool)
		if err != nil {
			return err
		}
		log.Println("Инициализация базы данных завершена")
	} else {
		log.Println("База данных уже инициализирована")
	}

	return nil
}
