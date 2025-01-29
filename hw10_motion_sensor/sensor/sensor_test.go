package sensor_test

import (
	"testing"
	"time"

	"github.com/Lushin415/HW-otus-go/hw10_motion_sensor/sensor"
)

func TestRandomGenerator(t *testing.T) {
	for i := 0; i < 100; i++ {
		num := sensor.RandomGenerator()
		if num < 0 || num >= 100000 {
			t.Errorf("Случайное число %d выходит за границы 0-99999", num)
		}
	}
}

// Тест для ReadChan (проверяем, выдаёт ли канал числа).
func TestReadChan(t *testing.T) {
	// Создаём кастомный генератор, который выдаёт последовательность 1, 2, 3...
	mockGen := func() func() int {
		counter := 0
		return func() int {
			counter++
			return counter
		}
	}()

	// Запускаем ReadChan, но читаем только 10 значений
	dataChan := sensor.ReadChan(mockGen)

	// Проверяем первые 10 значений
	for i := 1; i <= 10; i++ {
		select {
		case num, open := <-dataChan:
			if !open {
				t.Fatalf("Канал закрылся раньше времени на шаге %d", i)
			}
			if num != i {
				t.Errorf("Ожидалось %d, но получено %d", i, num)
			}
		case <-time.After(1100 * time.Millisecond): // Уменьшаем таймаут для скорости теста
			t.Fatal("Тест завис, канал ReadChan не выдаёт данные")
		}
	}
}
