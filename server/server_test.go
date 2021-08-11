package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultHandler(t *testing.T) {
	Init()
	InitAPI()
	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	rec := httptest.NewRecorder()
	c := EchoCon.NewContext(req, rec)

	if error := defaultGetOk(c); error != nil {
		t.Errorf(error.Error())
	}
	if rec.Code != http.StatusOK {
		t.Errorf("Invalid status code %d", rec.Code)
	}
}
