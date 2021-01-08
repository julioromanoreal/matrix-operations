package http_operations

import (
	"fmt"
	"matrix-operations/pkg/matrix_operations"
	"matrix-operations/pkg/matrix_utils"
	"net/http"
)

func sum(w http.ResponseWriter, r *http.Request) {
	records, ok := matrix_utils.ParseMatrix(w, r)
	if !ok {
		fmt.Fprintf(w, "error: not able to read the file")
		return
	}

	ops := matrix_operations.GetMatrixOperations()
	val, err := ops.Sum(records)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err.Error())
		return
	}

	fmt.Fprintf(w, "%v", val)
}
