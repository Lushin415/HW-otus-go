package book_test

import (
	"testing"

	book "github.com/Lushin415/HW-otus-go/06_testing/hw04/testing_book"
	"github.com/stretchr/testify/assert"
)

func TestBookMethods(t *testing.T) {
	// Книга
	newBook := book.NewBook(1, "Go Programming", "John Doe", 2020, 350, 4.5)

	// Проверка геттеров
	assert.Equal(t, 1, newBook.ID(), "ID должен быть 1")
	assert.Equal(t, "Go Programming", newBook.Title(), "Заголовок должен быть 'Go Programming'")
	assert.Equal(t, "John Doe", newBook.Author(), "Автор должен быть 'John Doe'")
	assert.Equal(t, 2020, newBook.Year(), "Год должен быть 2020")
	assert.Equal(t, 350, newBook.Size(), "Размер должен быть 350")
	assert.Equal(t, 4.5, newBook.Rate(), "Рейтинг должен быть 4.5")

	// Проверка сеттеров
	newBook.SetID(2)
	newBook.SetTitle("Go Advanced")
	newBook.SetAuthor("Jane Smith")
	newBook.SetYear(2023)
	newBook.SetSize(400)
	newBook.SetRate(4.8)

	// Проверка новых значений
	assert.Equal(t, 2, newBook.ID(), "ID должен обновиться и стать 2")
	assert.Equal(t, "Go Advanced", newBook.Title(), "Title должен обновиться и стать 'Go Advanced'")
	assert.Equal(t, "Jane Smith", newBook.Author(), "Author должен обновиться и стать 'Jane Smith'")
	assert.Equal(t, 2023, newBook.Year(), "Year должен обновиться и стать 2023")
	assert.Equal(t, 400, newBook.Size(), "Size должен обновиться и стать 400")
	assert.Equal(t, 4.8, newBook.Rate(), "Rate должен обновиться и стать 4.8")
}
