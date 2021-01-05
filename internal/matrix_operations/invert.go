package matrix_operations

import (
	"fmt"
	"matrix-operations/internal/matrix_utils"
	"strconv"
	"strings"
	"sync"
)

func invert(records [][]string) (string, error) {
	valid := matrix_utils.ValidateSquareMatrix(records)
	if !valid {
		return "", fmt.Errorf("matrix is not square")
	}

	_, err := strconv.Atoi(records[0][0])
	if err != nil {
		return "", fmt.Errorf("could not convert %v to int", records[0][0])
	}

	recordsLength := len(records)
	errors := make(chan error, 1)
	var wg sync.WaitGroup

	for i := 1; i < recordsLength; i++ {
		wg.Add(1)

		go func(ii int, records [][]string, wg *sync.WaitGroup) {
			for j := 0; j < ii; j++ {
				if len(errors) > 0 {
					break
				}

				_, err := strconv.Atoi(records[j][ii])
				if err != nil {
					errors <- fmt.Errorf("could not convert %v to int", records[j][ii])
					break
				}

				tmp := records[j][ii]
				records[j][ii] = records[ii][j]
				records[ii][j] = tmp
			}

			wg.Done()
		}(i, records, &wg)
	}

	wg.Wait()
	close(errors)

	if len(errors) > 0 {
		return "", <-errors
	}

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response, nil
}
