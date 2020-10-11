package main

import (
	"go-study/errorhandling/filelistingserve/filelisting"
	"go-study/modules/log"
	"net/http"
	"os"
)

type appHandle func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handle appHandle) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handle(w, r)
		if err != nil {
			log.Warn("Error Handling Request: %s\n", err.Error())
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

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
