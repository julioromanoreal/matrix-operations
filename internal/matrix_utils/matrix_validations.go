package matrix_utils

import (
	"matrix-operations/internal/file_utils"
	"net/http"
)

func ParseMatrix(w http.ResponseWriter, r *http.Request) ([][]string, bool) {
	records, err := file_utils.ReadFile(r, w)
	if err {
		return nil, false
	}

	return records, true
}

func ValidateSquareMatrix(matrix [][]string) bool {
	nrRows := len(matrix)
	for _, row := range matrix {
		if nrRows != len(row) {
			return false
		}
	}
	return true
}
