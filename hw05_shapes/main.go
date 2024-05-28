package main

import (
	"errors"
	"fmt"
)

type Shape interface {
	CalculateArea() float64
}

type Circle struct {
	r float64
}
type Triangle struct {
	a float64
	h float64
}

type Rectangle struct {
	q float64
	b float64
}

func (c Circle) CalculateArea() float64 {
	return 3.14 * c.r * c.r
}
func (t Triangle) CalculateArea() float64 {
	return 0.5 * t.a * t.h
}
func (re Rectangle) CalculateArea() float64 {
	return re.q * re.b
}
func CalculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("объект не реализует интерфейс Shape")
	}
	return shape.CalculateArea(), nil
}
func main() {
	var choice int
	fmt.Println("Выберите фигуру: 1. Круг 2. Треугольник 3. Прямоугольник:")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("ошибка выбора фигуры", err)
		return
	}

	var shape Shape
	switch choice {
	case 1:
		var r float64
		fmt.Println("Введите диаметр круга")
		_, err := fmt.Scan(&r)
		if err != nil {
			fmt.Println("Недопустимое значение", err)
			return
		}
		shape = Circle{r: r}
	case 2:
		var a, h float64
		fmt.Println("Введите основание и высоту треугольника:")
		_, err := fmt.Scan(&a, &h)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		shape = Triangle{a: a, h: h}
	case 3:
		var q, b float64
		fmt.Println("Введите длину и ширину прямоугольника:")
		_, err := fmt.Scan(&q, &b)
		if err != nil {
			fmt.Println("Недопустимое значение:", err)
			return
		}
		shape = Rectangle{q: q, b: b}
	default:
		fmt.Println("Недопустимое значение")
		return
	}
	area, err := CalculateArea(shape)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Площадь выбранной фигуры: %.2f\n", area)
	}
}
