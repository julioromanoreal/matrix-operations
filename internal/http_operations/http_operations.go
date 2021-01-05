package http_operations

import (
	"log"
	"matrix-operations/pkg"
	"net/http"
)

type httpOperationsImpl struct {
}

func GetHttpOperations() pkg.HttpOperations {
	ops := &httpOperationsImpl{}
	return ops
}

func (h httpOperationsImpl) Echo(writer http.ResponseWriter, request *http.Request) {
	log.Print("Executing Echo")
	echo(writer, request)
}

func (h httpOperationsImpl) Flatten(writer http.ResponseWriter, request *http.Request) {
	log.Print("Executing Flatten")
	flatten(writer, request)
}

func (h httpOperationsImpl) Invert(writer http.ResponseWriter, request *http.Request) {
	log.Print("Executing Invert")
	invert(writer, request)
}

func (h httpOperationsImpl) Sum(writer http.ResponseWriter, request *http.Request) {
	log.Print("Executing Sum")
	sum(writer, request)
}

func (h httpOperationsImpl) Multiply(writer http.ResponseWriter, request *http.Request) {
	log.Print("Executing Multiply")
	multiply(writer, request)
}
