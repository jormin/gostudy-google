package main

import (
	"fmt"
	"go-study/errorhandling/filelistingserve/filelisting"
	"go-study/modules/log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandle func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handle appHandle) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Info(fmt.Sprintf("Panic: %v", r))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handle(w, r)
		if err != nil {
			log.Warn("Error Handling Request: %s\n", err.Error())
			if userErr, ok := err.(UserError); ok {
				http.Error(w, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

type UserError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
