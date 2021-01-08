package matrix_operations

func multiply(records [][]string) (int, error) {
	return arithmetic(records, 1, func(i int, i2 int) int {
		return i * i2
	})
}
