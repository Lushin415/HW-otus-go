package chess_test

import (
	"testing"
)

func TestBoardSize(t *testing.T) {
	tests := []struct {
		rows    int
		columns int
		want    string
	}{
		{rows: 1, columns: 1, want: "#\n"},
		{rows: 3, columns: 3, want: "# #\n # \n# #\n"},
		{rows: 4, columns: 4, want: "# # \n # #\n# # \n # #\n"},
	}

	for _, size := range tests {
		t.Run("", func(t *testing.T) {
			result := logic.Chessboard(size.rows, size.columns)
			if result != size.want {
				t.Errorf("size %d: expected %q, got %q", size.rows, size.columns, size.want, result)
			}
		})
	}
}
