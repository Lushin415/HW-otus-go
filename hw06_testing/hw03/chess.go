package chess

func Chessboard(rows, columns int) string {

	var boardResult string
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
			boardResult = string(board[i][j])
		}

	}
	return boardResult
}
