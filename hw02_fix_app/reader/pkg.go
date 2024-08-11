package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/HW-otus-go/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)

	return data, err
}
