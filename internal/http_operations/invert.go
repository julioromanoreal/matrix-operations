package http_operations

import (
	"fmt"
	"matrix-operations/internal/matrix_operations"
	"matrix-operations/internal/matrix_utils"
	"net/http"
)

func invert(w http.ResponseWriter, r *http.Request) {
	records, ok := matrix_utils.ParseMatrix(w, r)
	if !ok {
		fmt.Fprintf(w, "error: not able to read the file")
		return
	}

	ops := matrix_operations.GetMatrixOperations()
	val, err := ops.Invert(records)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err.Error())
		return
	}

	fmt.Fprintf(w, "%v", val)
}
