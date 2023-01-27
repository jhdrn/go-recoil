package recoil

import (
	"net/http"
	"testing"
)

func TestHandle1(t *testing.T) {

	fn := Handler(func(r *http.Request) Response {
		arr := []string{"foo"}
		return response.JSON(arr).WithStatus(http.StatusCreated)
	})

}

func TestHandle2(t *testing.T) {

	fn := Handler(func(r *http.Request) Response {
		return response.OK()
	})

}
