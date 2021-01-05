package pkg

import "net/http"

type HttpOperations interface {
	Echo(http.ResponseWriter, *http.Request)
	Flatten(http.ResponseWriter, *http.Request)
	Invert(http.ResponseWriter, *http.Request)
	Sum(http.ResponseWriter, *http.Request)
	Multiply(http.ResponseWriter, *http.Request)
}
