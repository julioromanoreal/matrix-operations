package http_operations

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_httpOperationsImpl_Echo(t *testing.T) {
	requestBody, multipartWriter, err := getFileContent(t)
	request, err := http.NewRequest("POST", "localhost:8080/echo", &requestBody)
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ops := GetHttpOperations()
	handler := http.HandlerFunc(ops.Echo)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got \n%v\n want \n%v\n",
			rr.Body.String(), expected)
	}
}

func Test_httpOperationsImpl_Flatten(t *testing.T) {
	requestBody, multipartWriter, err := getFileContent(t)
	request, err := http.NewRequest("POST", "localhost:8080/echo", &requestBody)
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ops := GetHttpOperations()
	handler := http.HandlerFunc(ops.Flatten)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "1,2,3,4,5,6,7,8,9"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got \n%v\n want \n%v\n",
			rr.Body.String(), expected)
	}
}

func Test_httpOperationsImpl_Invert(t *testing.T) {
	requestBody, multipartWriter, err := getFileContent(t)
	request, err := http.NewRequest("POST", "localhost:8080/echo", &requestBody)
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ops := GetHttpOperations()
	handler := http.HandlerFunc(ops.Invert)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "1,4,7\n2,5,8\n3,6,9\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got \n%v\n want \n%v\n",
			rr.Body.String(), expected)
	}
}

func Test_httpOperationsImpl_Multiply(t *testing.T) {
	requestBody, multipartWriter, err := getFileContent(t)
	request, err := http.NewRequest("POST", "localhost:8080/echo", &requestBody)
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ops := GetHttpOperations()
	handler := http.HandlerFunc(ops.Multiply)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "362880"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got \n%v\n want \n%v\n",
			rr.Body.String(), expected)
	}
}

func Test_httpOperationsImpl_Sum(t *testing.T) {
	requestBody, multipartWriter, err := getFileContent(t)
	request, err := http.NewRequest("POST", "localhost:8080/echo", &requestBody)
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ops := GetHttpOperations()
	handler := http.HandlerFunc(ops.Sum)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "45"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got \n%v\n want \n%v\n",
			rr.Body.String(), expected)
	}
}

func getFileContent(t *testing.T) (bytes.Buffer, *multipart.Writer, error) {
	file, err := os.Open("../../resources/matrix.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)
	fileWriter, err := multipartWriter.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		t.Fatal(err)
	}
	multipartWriter.Close()
	return requestBody, multipartWriter, err
}
