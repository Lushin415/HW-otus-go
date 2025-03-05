package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HW-otus-go/Lushin415/hw12_log_util/analyzlog"
	"github.com/HW-otus-go/Lushin415/hw12_log_util/createlog"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

func main() {
	filePFlag := pflag.StringP("file", "f", "", "Путь к файлу логов")
	levelPFlag := pflag.StringP("level", "l", "", "Уровень логов (INFO, ERROR, WARN, DEBUG)")
	outputPFlag := pflag.StringP("output", "o", "", "Путь к файлу вывода")

	pflag.Parse()

	// Проверяем, что все флаги пустые
	if *filePFlag == "" && *levelPFlag == "" && *outputPFlag == "" {
		// Загружаем .env, но без panic
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("Не удалось загрузить .env: %v", err)
		}

		if os.Getenv("LOG_ANALYZER_FILE") == "" &&
			os.Getenv("LOG_ANALYZER_LEVEL") == "" &&
			os.Getenv("LOG_ANALYZER_OUTPUT") == "" {
			createlog.CreateLog()

			defaultLogPath := "logs.log"

			analyzlog.ScanLog(defaultLogPath, "", "")
			return
		}
	}

	if *filePFlag == "" {
		*filePFlag = os.Getenv("LOG_ANALYZER_FILE")
	}
	if *levelPFlag == "" {
		*levelPFlag = os.Getenv("LOG_ANALYZER_LEVEL")
	}
	if *outputPFlag == "" {
		*outputPFlag = os.Getenv("LOG_ANALYZER_OUTPUT")
	}

	createlog.CreateLog()

	if *filePFlag == "" {
		*filePFlag = "logs.log"
	}

	analyzlog.ScanLog(*filePFlag, *outputPFlag, *levelPFlag)

	fmt.Println("Файл логов:", *filePFlag)
	fmt.Println("Уровень логов:", *levelPFlag)
	fmt.Println("Файл вывода:", *outputPFlag)
}
