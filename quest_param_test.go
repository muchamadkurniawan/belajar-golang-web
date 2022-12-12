package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "hello")
	} else {
		fmt.Fprintf(w, "hi %s", name)
	}
}

func MultipleParam(writer http.ResponseWriter, request *http.Request) {
	fist_name := request.URL.Query().Get("first")
	last_name := request.URL.Query().Get("last")
	fmt.Fprintf(writer, "first name %s, last name %s", fist_name, last_name)
}
func TestQuerryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Eko", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestMultipleQuerryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,
		"http://localhost:8080/hello?first=Eko&last=kurniawan",
		nil)
	recorder := httptest.NewRecorder()

	MultipleParam(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleParamArray(writer http.ResponseWriter, request *http.Request) {
	querry := request.URL.Query()
	names := querry["param"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultiplePamasValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,
		"http://localhost:8080/hello?param=Eko&param=kurniawan&param=testing",
		nil)
	recorder := httptest.NewRecorder()

	MultipleParamArray(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
