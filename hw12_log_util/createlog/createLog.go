package createlog

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

func CreateLog() {
	authLog, err := os.Open("/var/log/auth.log")
	if err != nil {
		panic("Ошибка при открытии auth.log: " + err.Error())
	}
	defer authLog.Close()

	logFile, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		panic("Ошибка при открытии файла logs.log: " + err.Error())
	}
	defer logFile.Close()

	logger := slog.New(slog.NewTextHandler(logFile, nil))

	lines := ReadLastLines(authLog, 25)
	for _, line := range lines {
		level := LogLevel(line)
		switch level {
		case "ERROR":
			logger.Error("Auth Log", "message", line)
		case "WARN":
			logger.Warn("Auth Log", "message", line)
		case "DEBUG":
			logger.Debug("Auth Log", "message", line)
		default:
			logger.Info("Auth Log", "message", line)
		}
	}
}

func LogLevel(line string) string {
	lowerLine := strings.ToLower(line)
	if strings.Contains(lowerLine, "error") || strings.Contains(lowerLine, "failed") {
		return "ERROR"
	}
	if strings.Contains(lowerLine, "warn") {
		return "WARN"
	}
	if strings.Contains(lowerLine, "debug") {
		return "DEBUG"
	}
	return "INFO"
}

func ReadLastLines(file *os.File, n int) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) > n {
			lines = lines[1:] // Удаляем первую строку, если больше n
		}
	}
	return lines
}
