package comparator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookMethods(t *testing.T) {
	// Создание объекта книги
	book := &Book{
		id:     1,
		title:  "Go Programming",
		author: "John Doe",
		year:   2020,
		size:   350,
		rate:   4.5,
	}

	// Проверка геттеров
	assert.Equal(t, 1, book.ID(), "ID should be 1")
	assert.Equal(t, "Go Programming", book.Title(), "Title should be 'Go Programming'")
	assert.Equal(t, "John Doe", book.Author(), "Author should be 'John Doe'")
	assert.Equal(t, 2020, book.Year(), "Year should be 2020")
	assert.Equal(t, 350, book.Size(), "Size should be 350")
	assert.Equal(t, 4.5, book.Rate(), "Rate should be 4.5")

	// Проверка сеттеров
	book.SetID(2)
	book.SetTitle("Go Advanced")
	book.SetAuthor("Jane Smith")
	book.SetYear(2023)
	book.SetSize(400)
	book.SetRate(4.8)

	// Проверка обновленных значений
	assert.Equal(t, 2, book.ID(), "ID should be updated to 2")
	assert.Equal(t, "Go Advanced", book.Title(), "Title should be updated to 'Go Advanced'")
	assert.Equal(t, "Jane Smith", book.Author(), "Author should be updated to 'Jane Smith'")
	assert.Equal(t, 2023, book.Year(), "Year should be updated to 2023")
	assert.Equal(t, 400, book.Size(), "Size should be updated to 400")
	assert.Equal(t, 4.8, book.Rate(), "Rate should be updated to 4.8")
}

func TestComparator(t *testing.T) {
	// Создание двух книг
	bookOne := &Book{
		id:     1,
		title:  "Book One",
		author: "Author One",
		year:   2020,
		size:   300,
		rate:   4.0,
	}

	bookTwo := &Book{
		id:     2,
		title:  "Book Two",
		author: "Author Two",
		year:   2021,
		size:   350,
		rate:   4.5,
	}

	// Сравнение по году
	comparator := NewComparator(PoYear)
	assert.True(t, comparator.Compare(bookTwo, bookOne), "BookTwo should be after BookOne by Year")

	// Сравнение по размеру
	comparator = NewComparator(PoSize)
	assert.True(t, comparator.Compare(bookTwo, bookOne), "BookTwo should be after BookOne by Size")

	// Сравнение по рейтингу
	comparator = NewComparator(PoRate)
	assert.True(t, comparator.Compare(bookTwo, bookOne), "BookTwo should be after BookOne by Rate")

	// Сравнение по полю по умолчанию
	comparator = NewComparator(0) // Подаем неверное значение
	assert.False(t, comparator.Compare(bookOne, bookTwo), "Comparison with invalid field should return false")
}
