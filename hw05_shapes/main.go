package main

import (
	"errors"
	"fmt"
)

type Shape interface {
	CalculateArea() float64
}

type Circle struct {
	R float64
}
type Triangle struct {
	A float64
	H float64
}

type Rectangle struct {
	Q float64
	B float64
}

func (c Circle) CalculateArea() float64 {
	return 3.14 * c.R * c.R
}

func (t Triangle) CalculateArea() float64 {
	return 0.5 * t.A * t.H
}

func (re Rectangle) CalculateArea() float64 {
	return re.Q * re.B
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
		_, scanErr := fmt.Scan(&r)
		if scanErr != nil {
			fmt.Println("Недопустимое значение", scanErr)
			return
		}
		shape = Circle{R: r}
	case 2:
		var a, h float64
		fmt.Println("Введите основание и высоту треугольника:")
		_, scanErr := fmt.Scan(&a, &h)
		if scanErr != nil {
			fmt.Println("Ошибка ввода:", scanErr)
			return
		}
		shape = Triangle{A: a, H: h}
	case 3:
		var q, b float64
		fmt.Println("Введите длину и ширину прямоугольника:")
		_, scanErr := fmt.Scan(&q, &b)
		if scanErr != nil {
			fmt.Println("Недопустимое значение:", scanErr)
			return
		}
		shape = Rectangle{Q: q, B: b}
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
