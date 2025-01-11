package chess

func Chessboard(rows, columns int) string {
	var boardResult string

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if (i+j)%2 == 0 {
				boardResult += "#"
			} else {
				boardResult += " "
			}
		}
		boardResult += "\n"
	}

	return boardResult
}
