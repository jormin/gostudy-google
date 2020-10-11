package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(w http.ResponseWriter, r *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(w http.ResponseWriter, r *http.Request) error {
	return testingUserError("this is user error")
}

func errNotFound(w http.ResponseWriter, r *http.Request) error {
	return os.ErrNotExist
}

func errPermission(w http.ResponseWriter, r *http.Request) error {
	return os.ErrPermission
}

func errUnknow(w http.ResponseWriter, r *http.Request) error {
	return errors.New("unknow error")
}

func noErr(w http.ResponseWriter, r *http.Request) error {
	_, _ = fmt.Fprintf(w, "no error")
	return nil
}

var tests = []struct {
	h       appHandle
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "this is user error"},
	{errNotFound, 404, "Not Found"},
	{errPermission, 403, "Forbidden"},
	{errUnknow, 500, "Internal Server Error"},
	{noErr, 200, "no error"},
}

func Test_errWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func Test_errWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(response *http.Response, code int, message string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != code || body != message {
		t.Errorf("expect (%d, %s); got (%d, %s)", code, message, response.StatusCode, body)
	}
}
