package db

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Manager представляет собой объект для работы с базой данных
type Manager struct {
	pool        *pgxpool.Pool
	dbName      string
	sqlFilePath string
}

// NewManager создает новый экземпляр менеджера базы данных
func NewManager(pool *pgxpool.Pool, dbName, sqlFilePath string) *Manager {
	return &Manager{
		pool:        pool,
		dbName:      dbName,
		sqlFilePath: sqlFilePath,
	}
}

// Connect устанавливает соединение с базой данных
func Connect(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	// Проверка соединения
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ошибка проверки соединения с базой данных: %w", err)
	}

	return pool, nil
}

// ExecuteSQL выполняет SQL-скрипт из файла
func (m *Manager) ExecuteSQL(ctx context.Context) error {
	// Читаем SQL-скрипт
	sqlBytes, err := os.ReadFile(m.sqlFilePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения SQL файла %s: %w", m.sqlFilePath, err)
	}

	sqlScript := string(sqlBytes)
	commands := strings.Split(sqlScript, ";")

	for _, cmd := range commands {
		trimmedCmd := strings.TrimSpace(cmd)
		if trimmedCmd == "" {
			continue
		}

		_, err = m.pool.Exec(ctx, trimmedCmd)
		if err != nil {
			return fmt.Errorf("ошибка выполнения SQL-команды: %w", err)
		}
	}

	return nil
}

// Close закрывает соединение с базой данных
func (m *Manager) Close() {
	if m.pool != nil {
		m.pool.Close()
	}
}
