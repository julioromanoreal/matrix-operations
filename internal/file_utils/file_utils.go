package file_utils

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func ReadFile(r *http.Request, w http.ResponseWriter) ([][]string, bool) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, true
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, true
	}
	return records, false
}
