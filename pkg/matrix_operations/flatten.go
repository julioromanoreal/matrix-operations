package matrix_operations

import (
	"fmt"
	"matrix-operations/pkg/matrix_utils"
	"strconv"
	"strings"
	"sync"
)

func flatten(records [][]string) (string, error) {
	valid := matrix_utils.ValidateSquareMatrix(records)
	if !valid {
		return "", fmt.Errorf("matrix is not square")
	}

	errors := make(chan error, 1)
	var wg sync.WaitGroup

	for _, row := range records {
		wg.Add(1)

		go func(r []string, wg *sync.WaitGroup) {
			for _, val := range r {
				if len(errors) > 0 {
					break
				}

				value, err := strconv.Atoi(val)
				if err != nil {
					errors <- fmt.Errorf("could not convert %v to int", value)
					break
				}
			}

			wg.Done()
		}(row, &wg)
	}

	wg.Wait()
	close(errors)

	if len(errors) > 0 {
		return "", <-errors
	}

	var str strings.Builder
	for _, row := range records {
		if str.Len() > 0 {
			str.WriteString(",")
		}

		str.WriteString(strings.Join(row, ","))
	}
	return str.String(), nil
}
