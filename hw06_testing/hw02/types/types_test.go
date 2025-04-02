package types_test

import (
	"testing"

	"github.com/Lushin415/HW-otus-go/06_testing/hw02/types"
	"github.com/stretchr/testify/assert"
)

func TestEmployee(t *testing.T) {
	employee1 := types.Employee{
		UserID:       10,
		Age:          25,
		Name:         "Rob",
		DepartmentID: 3,
	}
	employee2 := types.Employee{
		UserID:       11,
		Age:          30,
		Name:         "George",
		DepartmentID: 2,
	}
	result1 := employee1.String()
	expected1 := "User ID: 10; Age: 25; Name: Rob; Department ID: 3; "
	result2 := employee2.String()
	expected2 := "User ID: 11; Age: 30; Name: George; Department ID: 2; "
	assert.Equal(t, expected1, result1)
	assert.Equal(t, expected2, result2)
}
