package printer

import (
	"fmt"

	"HW-otus/hw06_testing/hw02/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
}
