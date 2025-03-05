package createlog_test

import (
	"os"
	"strings"
	"testing"

	"github.com/HW-otus-go/Lushin415/hw12_log_util/createlog"
)

func TestLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"This is an error message", "ERROR"},
		{"This is a warning", "WARN"},
		{"Debugging info", "DEBUG"},
		{"Just an info message", "INFO"},
	}

	for _, tt := range tests {
		result := createlog.LogLevel(tt.input)
		if result != tt.expected {
			t.Errorf("LogLevel(\"%s\") = \"%s\"; expected \"%s\"", tt.input, result, tt.expected)
		}
	}
}

func TestReadLastLines(t *testing.T) {
	testData := `
	line 1
	line 2
	line 3
	line 4
	`
	tmpFile, err := os.CreateTemp("", "test.log")
	if err != nil {
		t.Fatalf("Не удалось создать временный файл: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(strings.TrimSpace(testData) + "\n")
	if err != nil {
		t.Fatalf("Ошибка записи в файл: %v", err)
	}
	tmpFile.Close()

	file, err := os.Open(tmpFile.Name())
	if err != nil {
		t.Fatalf("Ошибка открытия временного файла: %v", err)
	}
	defer file.Close()

	lines := createlog.ReadLastLines(file, 3)
	expected := []string{"line 2", "line 3", "line 4"}
	if len(lines) != len(expected) {
		t.Fatalf("Ожидалось %d строк, получено %d", len(expected), len(lines))
	}

	for i, line := range lines {
		if strings.TrimSpace(line) != expected[i] {
			t.Errorf("Ожидалась строка \"%s\", получена \"%s\"", expected[i], line)
		}
	}
}
