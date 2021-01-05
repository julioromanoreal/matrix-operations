package pkg

type MatrixOperations interface {
	Echo([][]string) (string, error)
	Flatten([][]string) (string, error)
	Invert([][]string) (string, error)
	Sum([][]string) (string, error)
	Multiply([][]string) (string, error)
}
