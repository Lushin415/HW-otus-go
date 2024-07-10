package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Lushin415/HW-otus-go/hw07_word_counter/wordcounter"
)

func main() {
	fmt.Println("Введите текст:")

	// Читаем, что написал пользователь
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
		return
	}

	// Вызываем функцию CountWords и выводим результат
	wordCounts := wordcounter.CountWords(text)
	fmt.Println("Мапа с количеством упоминаний слов:", wordCounts)
}
