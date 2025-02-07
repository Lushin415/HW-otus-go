package analyzLog

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LogLevelCount struct {
	Info  int
	Error int
	Warn  int
	Debug int
	Other int
}

func AnalyzeLog(s string) string {
	idx := strings.Index(s, "level=")
	if idx == -1 {
		return ""
	}
	start := idx + len("level=")
	end := strings.IndexByte(s[start:], ' ')
	if end == -1 {
		return s[start:]
	}
	return s[start : start+end]
}

func ScanLog(filePath, outputPath, logLevel string) {
	log, err := os.Open(filePath)
	if err != nil {
		panic("Ошибка при открытии logs.log: " + err.Error())
	}
	defer log.Close()

	counter := LogLevelCount{}
	scanner := bufio.NewScanner(log)
	for scanner.Scan() {
		line := scanner.Text()
		level := AnalyzeLog(line)
		if logLevel != "" && level != logLevel {
			continue
		}
		switch level {
		case "INFO":
			counter.Info++
		case "ERROR":
			counter.Error++
		case "WARN":
			counter.Warn++
		case "DEBUG":
			counter.Debug++
		default:
			counter.Other++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	outFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Ошибка при открытии файла out.txt: " + err.Error())
	}

	defer outFile.Close()

	_, err = fmt.Fprintf(outFile, "INFO: %d\nERROR: %d\nWARN: %d\nDEBUG: %d\nOther: %d\n",
		counter.Info, counter.Error, counter.Warn, counter.Debug, counter.Other)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
}
