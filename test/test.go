package test

import (
	"net/http"

	"github.com/jhdrn/go-recoil"
	"github.com/jhdrn/go-recoil/response"
)

var xmlResponseBuilder = response.NewBuilder(response.WithConfig(
	response.Config{
		Formatter: response.XMLFormatter{},
	},
))

func main() {

	// x := recoil.HandlerFunc(func(r *http.Request) recoil.Response {
	// 	return response.NotImplemented()
	// })

	// y := recoil.Handler(func(r *http.Request) recoil.Response {
	// 	return response.NotImplemented()
	// })

	response.DefaultConfig.Formatter = response.JSONFormatter{}

	http.HandleFunc("/", recoil.HandlerFunc(func(r *http.Request) recoil.Response {
		return xmlResponseBuilder.WithContent("foo").WithStatus(http.StatusAccepted)
	}))

	http.Handle("/", recoil.Handler(func(r *http.Request) recoil.Response {
		var err error
		if err != nil {
			return response.InternalServerError()
		}

		return response.OK().WithContent("foo")
	}))

	mux := http.NewServeMux()
	mux.Handle("/", recoil.Handler(func(r *http.Request) recoil.Response {

		var valid bool
		if !valid {
			return response.Content("").WithStatus(http.StatusBadRequest)
			// return response.BadRequest().WithContent("")
		}

		var err error
		if err != nil {
			return response.InternalServerError().WithContent(err)
			// return response.Content(err).WithStatus(http.StatusInternalServerError)
		}

		obj := []string{}

		response.Status(http.StatusOK).WithContent(obj)
		response.Content(obj).WithStatus(http.StatusOK)
		return response.Created().WithContent(obj)
	}))
}
