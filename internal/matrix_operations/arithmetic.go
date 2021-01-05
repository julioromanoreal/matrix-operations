package matrix_operations

import (
	"fmt"
	"matrix-operations/internal/matrix_utils"
	"strconv"
	"sync"
)

func arithmetic(records [][]string, initialValue int, operation func(int, int) int) (int, error) {
	valid := matrix_utils.ValidateSquareMatrix(records)
	if !valid {
		return 0, fmt.Errorf("matrix is not square")
	}

	results := make(chan int, len(records))
	errors := make(chan error, 1)
	var wg sync.WaitGroup

	for _, row := range records {
		wg.Add(1)

		go func(r []string, wg *sync.WaitGroup) {
			tmpResult := initialValue
			for _, val := range r {
				if len(errors) > 0 {
					break
				}

				value, err := strconv.Atoi(val)
				if err != nil {
					errors <- fmt.Errorf("could not convert %v to int", value)
					break
				}
				tmpResult = operation(tmpResult, value)
			}

			results <- tmpResult

			wg.Done()
		}(row, &wg)
	}

	wg.Wait()
	close(results)
	close(errors)

	if len(errors) > 0 {
		return 0, <-errors
	}

	finalResult := initialValue
	for rowResult := range results {
		finalResult = operation(rowResult, finalResult)
	}

	return finalResult, nil
}
