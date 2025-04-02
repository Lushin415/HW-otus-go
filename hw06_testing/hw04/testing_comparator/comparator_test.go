package comparator_test

import (
	"testing"

	book "github.com/Lushin415/HW-otus-go/06_testing/hw04/testing_book"
	comparator "github.com/Lushin415/HW-otus-go/06_testing/hw04/testing_comparator"
	"github.com/stretchr/testify/assert"
)

func TestComparator(t *testing.T) {
	// Создание двух книг с использованием конструктора NewBook
	bookOne := book.NewBook(1, "Book One", "Author One", 2020, 300, 4.0)
	bookTwo := book.NewBook(2, "Book Two", "Author Two", 2021, 350, 4.5)

	// Сравнение по году
	yearComparator := comparator.NewComparator(comparator.PoYear)
	assert.True(t, yearComparator.Compare(*bookTwo, *bookOne), "BookTwo должна быть больше BookOne по году")

	// Сравнение по размеру
	sizeComparator := comparator.NewComparator(comparator.PoSize)
	assert.True(t, sizeComparator.Compare(*bookTwo, *bookOne), "BookTwo должна быть больше BookOne по размеру")

	// Сравнение по рейтингу
	rateComparator := comparator.NewComparator(comparator.PoRate)
	assert.True(t, rateComparator.Compare(*bookTwo, *bookOne), "BookTwo должна быть больше BookOne по рейтингу")

	// Сравнение с неверным значением (по умолчанию)
	invalidComparator := comparator.NewComparator(100) // Передаем некорректное значение
	assert.False(t, invalidComparator.Compare(*bookOne, *bookTwo),
		"Сравнение с некорректным значением должно вернуть false")
}
