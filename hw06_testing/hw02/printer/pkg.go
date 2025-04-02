package printer

import (
	"fmt"

	"github.com/Lushin415/HW-otus-go/06_testing/hw02/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
}
