package analyzLog_test

import (
	"os"
	"strings"
	"testing"

	"github.com/HW-otus-go/Lushin415/hw12_log_util/analyzLog"
)

func TestAnalyzeLog(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"time=2025-02-06T12:38:00.136+08:00 level=INFO msg=\"Test\"", "INFO"},
		{"time=2025-02-06T12:38:00.136+08:00 level=ERROR msg=\"Test\"", "ERROR"},
		{"time=2025-02-06T12:38:00.136+08:00 level=WARN msg=\"Test\"", "WARN"},
		{"time=2025-02-06T12:38:00.136+08:00 level=DEBUG msg=\"Test\"", "DEBUG"},
		{"time=2025-02-06T12:38:00.136+08:00 level=UNKNOWN msg=\"Test\"", "UNKNOWN"},
		{"time=2025-02-06T12:38:00.136+08:00 message=\"No level\"", ""},
	}
	for _, tt := range tests {
		result := analyzLog.AnalyzeLog(tt.input)
		if result != tt.expected {
			t.Errorf("AnalyzeLog(\"%s\") = \"%s\"; expected \"%s\"", tt.input, result, tt.expected)
		}
	}
}

func TestScanLog(t *testing.T) {
	testData := `
    time=2025-02-06T12:38:00.136+08:00 level=INFO msg="Test Info"
    time=2025-02-06T12:38:01.136+08:00 level=ERROR msg="Test Error"
    time=2025-02-06T12:38:02.136+08:00 level=WARN msg="Test Warn"
    time=2025-02-06T12:38:03.136+08:00 level=DEBUG msg="Test Debug"
    time=2025-02-06T12:38:04.136+08:00 level=UNKNOWN msg="Test Unknown"
    `
	tmpFile, err := os.CreateTemp("", "logs.log")
	if err != nil {
		t.Fatalf("Не удалось создать временный файл: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(strings.TrimSpace(testData) + "\n")
	if err != nil {
		t.Fatalf("Ошибка записи во временный файл: %v", err)
	}
	tmpFile.Close()

	oldLogPath := "logs.log"
	os.Rename(tmpFile.Name(), oldLogPath)
	defer os.Rename(oldLogPath, tmpFile.Name())

	inputPath := oldLogPath
	outputPath := "out.txt"
	logLevel := ""

	analyzLog.ScanLog(inputPath, outputPath, logLevel)

	outData, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Ошибка чтения out.txt: %v", err)
	}

	expectedOut := "INFO: 1\nERROR: 1\nWARN: 1\nDEBUG: 1\nOther: 1\n"
	if string(outData) != expectedOut {
		t.Errorf("Ожидался результат:\n%s,\nно получено:\n%s", expectedOut, string(outData))
	}
}
