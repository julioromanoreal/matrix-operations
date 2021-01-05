package matrix_operations

import (
	"matrix-operations/pkg"
	"strconv"
)

type matrixOperationsImpl struct {
}

func GetMatrixOperations() pkg.MatrixOperations {
	ops := &matrixOperationsImpl{}
	return ops
}

func (m *matrixOperationsImpl) Echo(i [][]string) (string, error) {
	return echo(i)
}

func (m *matrixOperationsImpl) Flatten(i [][]string) (string, error) {
	return flatten(i)
}

func (m *matrixOperationsImpl) Invert(i [][]string) (string, error) {
	return invert(i)
}

func (m *matrixOperationsImpl) Sum(i [][]string) (string, error) {
	value, err := sum(i)
	return strconv.Itoa(value), err
}

func (m *matrixOperationsImpl) Multiply(i [][]string) (string, error) {
	value, err := multiply(i)
	return strconv.Itoa(value), err
}
