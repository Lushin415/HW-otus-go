package binary_test

import (
	"reflect"
	"testing"

	"github.com/Lushin415/HW-otus-go/hw08_binary_search/binary"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Границы", args: args{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 3}, want: 3},
		{name: "Элемент отсутствует", args: args{arr: []int{1, 2, 3, 4, 5}, target: 10}, want: -1},
		{name: "Поиск первого элемента", args: args{arr: []int{1, 2, 3, 4, 5}, target: 1}, want: 0},
		{name: "Поиск последнего элемента", args: args{arr: []int{1, 2, 3, 4, 5}, target: 5}, want: 4},
		{name: "Пустой массив", args: args{arr: []int{}, target: 1}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binary.SearchBinary(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateArray(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Границы", args: args{a: 0, b: 10}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{name: "Отрицательные", args: args{a: -5, b: 0}, want: []int{-5, -4, -3, -2, -1, 0}},
		{name: "Один элемент", args: args{a: 3, b: 3}, want: []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binary.CreateArray(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
