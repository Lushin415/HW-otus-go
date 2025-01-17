package binary

func CreateArray(a, b int) []int {
	if a > b {
		return nil
	}
	arr := make([]int, b-a+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = a + i
	}
	return arr
}

func SearchBinary(arr []int, target int) int {
	i := 0
	j := len(arr)
	for i != j {
		m := (i + j) / 2
		if target == arr[m] {
			return m
		}
		if target < arr[m] {
			j = m
		} else {
			i = m + 1
		}
	}
	return -1
}
