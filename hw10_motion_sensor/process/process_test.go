package process_test

import (
	"testing"

	"github.com/Lushin415/HW-otus-go/hw10_motion_sensor/process"
)

func TestProcess(t *testing.T) {
	// Создаём каналы для ввода и вывода
	input := make(chan int, 20)
	output := make(chan float64, 2)

	go func() {
		for i := 1; i <= 20; i++ {
			input <- i
		}
		close(input)
	}()

	go process.Process(input, output)

	expected := []float64{5.5, 15.5}

	for i, exp := range expected {
		got, ok := <-output
		if !ok {
			t.Fatalf("Канал output закрылся раньше времени на шаге %d", i+1)
		}
		if got != exp {
			t.Errorf("Ожидалось %.2f, но получено %.2f на шаге %d", exp, got, i+1)
		}
	}

	_, open := <-output
	if open {
		t.Error("Канал output должен быть закрыт")
	}
}
