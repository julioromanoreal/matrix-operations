package matrix_operations

func sum(records [][]string) (int, error) {
	return arithmetic(records, 0, func(i int, i2 int) int {
		return i + i2
	})
}
