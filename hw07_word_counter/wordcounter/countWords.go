package wordcounter

import (
	"strings"
	"unicode"
)

// CountWords принимает строку текста и возвращает мапу,
// где ключ — это слово, а значение — количество его появлений в тексте.
func CountWords(text string) map[string]int {
	// Создаем мапу для хранения слов и их частоты
	wordCount := make(map[string]int)

	// Разделяем текст на слова
	words := strings.Fields(text)

	// Создаем слайс для хранения очищенных слов
	var cleanedWords []string

	for _, word := range words {
		// Приводим слово к нижнему регистру
		word = strings.ToLower(word)

		// Убираем пунктуацию из слова
		var cleanedWord strings.Builder
		for _, char := range word {
			if unicode.IsLetter(char) || unicode.IsNumber(char) {
				cleanedWord.WriteRune(char)
			}
		}

		// Получаем итоговое очищенное слово и добавляем в слайс
		finalWord := cleanedWord.String()
		if finalWord != "" {
			cleanedWords = append(cleanedWords, finalWord)
		}
	}

	// Подсчитываем частоту каждого слова в слайсе
	for _, word := range cleanedWords {
		wordCount[word]++
	}

	return wordCount
}
