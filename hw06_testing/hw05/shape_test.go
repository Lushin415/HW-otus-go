package calculateshape_test

import (
	"fmt"
	"math"
	"testing"

	calculateshape "HW-otus/hw06_testing/hw05"
)

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		shape    calculateshape.Shape
		expected float64
	}{
		{shape: calculateshape.Circle{R: 5}, expected: 78.5},
		{shape: calculateshape.Triangle{A: 10, H: 5}, expected: 25},
		{shape: calculateshape.Rectangle{Q: 4, B: 5}, expected: 20},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Testing %T", tc.shape),
			func(t *testing.T) {
				result, err := calculateshape.CalculateArea(tc.shape)
				if err != nil {
					t.Errorf("Неожиданная ошибка: %v", err)
				}
				if math.Abs(result-tc.expected) > 1e-9 {
					t.Errorf("Ожидание: %2f, получено: %2f", tc.expected, result)
				}
			})
	}
}

func TestInvalidShape(t *testing.T) {
	_, err := calculateshape.CalculateArea("Неверная фигура")
	if err == nil {
		t.Errorf("Ошибка ожидаемой фигуры, nil")
	}
}
