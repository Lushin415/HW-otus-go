package reader

import (
	"errors"
	"testing"

	"HW-otus/hw06_testing/hw02/types"
	"github.com/stretchr/testify/assert"
)

func TestReadJSON(t *testing.T) {
	testCases := []struct {
		name        string
		filePath    string
		expected    []types.Employee
		expectError error
	}{
		{
			name:     "Valid JSON",
			filePath: "data.json",
			expected: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
				{
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
		}, {
			name:        "Invalid JSON",
			filePath:    "invalid.json",
			expected:    nil,
			expectError: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ReadJSON(tc.filePath)
			if errors.Is(tc.expectError, assert.AnError) {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, result, "")
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.Equal(t, tc.expected, result, "Expected values to be equal")
			}
		})
	}
}
