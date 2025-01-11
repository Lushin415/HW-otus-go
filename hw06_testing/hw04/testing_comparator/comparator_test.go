package comparator_test

import (
	book "HW-otus/hw06_testing/hw04/testing_book"
	comparator "HW-otus/hw06_testing/hw04/testing_comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparator(t *testing.T) {
	// Создание двух книг с использованием конструктора NewBook
	bookOne := book.NewBook(1, "Book One", "Author One", 2020, 300, 4.0)
	bookTwo := book.NewBook(2, "Book Two", "Author Two", 2021, 350, 4.5)

	// Сравнение по году
	yearComparator := comparator.NewComparator(comparator.PoYear)
	assert.True(t, yearComparator.Compare(*bookTwo, *bookOne), "BookTwo should be after BookOne by Year")

	// Сравнение по размеру
	sizeComparator := comparator.NewComparator(comparator.PoSize)
	assert.True(t, sizeComparator.Compare(*bookTwo, *bookOne), "BookTwo should be after BookOne by Size")

	// Сравнение по рейтингу
	rateComparator := comparator.NewComparator(comparator.PoRate)
	assert.True(t, rateComparator.Compare(*bookTwo, *bookOne), "BookTwo should be after BookOne by Rate")

	// Сравнение с неверным значением (по умолчанию)
	invalidComparator := comparator.NewComparator(100) // Передаем некорректное значение
	assert.False(t, invalidComparator.Compare(*bookOne, *bookTwo), "Comparison with invalid field should return false")
}
