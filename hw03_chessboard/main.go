package main

import "fmt"

func main() {
	var rows, columns int
	fmt.Print("Введите количество строк доски: ")
	fmt.Scanln(&rows)
	fmt.Print("Введите количество столбцов доски: ")
	fmt.Scanln(&columns)

	board := make([][]rune, rows)
	for i := range board {
		board[i] = make([]rune, columns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if (i+j)%2 == 0 {
				board[i][j] = ' '
			} else {
				board[i][j] = '#'
			}
			fmt.Print(string(board[i][j]))
		}
		fmt.Println()
	}

}
