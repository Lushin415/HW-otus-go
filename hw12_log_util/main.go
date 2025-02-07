package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HW-otus-go/Lushin415/hw12_log_util/analyzLog"
	"github.com/HW-otus-go/Lushin415/hw12_log_util/createlog"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

func main() {
	filePFlag := pflag.StringP("file", "f", "", "Путь к файлу логов")
	levelPFlag := pflag.StringP("level", "l", "", "Уровень логов (INFO, ERROR, WARN, DEBUG)")
	outputPFlag := pflag.StringP("output", "o", "", "Путь к файлу вывода")

	pflag.Parse()

	if *filePFlag == "" || *levelPFlag == "" || *outputPFlag == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("Не удалось загрузить .env: %v", err)
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

	if *filePFlag == "" || *outputPFlag == "" {
		log.Fatal("Необходимо указать --file и --output либо задать их в .env")
	}
	createlog.CreateLog()
	file, errOpen := os.Open(*filePFlag)
	if errOpen != nil {
		fmt.Printf("Ошибка открытия файла %s: %v\n", *filePFlag, errOpen)
	}

	file.Close()

	fmt.Println("Файл логов:", *filePFlag)
	fmt.Println("Уровень логов:", *levelPFlag)
	fmt.Println("Файл вывода:", *outputPFlag)

	analyzLog.ScanLog(*filePFlag, *outputPFlag, *levelPFlag)
}
