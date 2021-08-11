package server

import (
	"github.com/rhperera/marvel-comic-api/config"
	"github.com/rhperera/marvel-comic-api/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	config.InitForTests()
	Init()
	InitAPI()
}

func TestHandler_GetCharacterById_ValidID(t *testing.T) {
	cacheService := &services.RedisCacheService{}
	cacheService.Connect()
	handler := NewHandler(cacheService, &services.MarvelCharacterAPI{})
	req := httptest.NewRequest(http.MethodGet,
		"/api/v1/characters/1011334", nil)
	rec := httptest.NewRecorder()
	c := EchoCon.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1011334")
	if error := handler.GetCharacterById(c); error != nil {
		t.Errorf(error.Error())
	}
	if rec.Code != http.StatusOK {
		t.Errorf("Invalid status code %d", rec.Code)
	}
}

func TestHandler_GetCharacterById_NonExistID(t *testing.T) {
	cacheService := &services.RedisCacheService{}
	cacheService.Connect()
	handler := NewHandler(cacheService, &services.MarvelCharacterAPI{})
	req := httptest.NewRequest(http.MethodGet,
		"/api/v1/characters/23", nil)
	rec := httptest.NewRecorder()
	c := EchoCon.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("23")
	if error := handler.GetCharacterById(c); error != nil {
		t.Errorf(error.Error())
	}
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Invalid status code %d", rec.Code)
	}
}

func TestHandler_GetCharacterById_NotANumberID(t *testing.T) {
	cacheService := &services.RedisCacheService{}
	cacheService.Connect()
	handler := NewHandler(cacheService, &services.MarvelCharacterAPI{})
	req := httptest.NewRequest(http.MethodGet,
		"/api/v1/characters/AAA", nil)
	rec := httptest.NewRecorder()
	c := EchoCon.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("AAA")
	if error := handler.GetCharacterById(c); error != nil {
		t.Errorf(error.Error())
	}
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Invalid status code %d", rec.Code)
	}
}

func TestHandler_GetCharacterById_NoIDParam(t *testing.T) {
	cacheService := &services.RedisCacheService{}
	cacheService.Connect()
	handler := NewHandler(cacheService, &services.MarvelCharacterAPI{})
	req := httptest.NewRequest(http.MethodGet,
		"/api/v1/characters/id", nil)
	rec := httptest.NewRecorder()
	c := EchoCon.NewContext(req, rec)
	if error := handler.GetCharacterById(c); error != nil {
		t.Errorf(error.Error())
	}
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Invalid status code %d", rec.Code)
	}
}

func TestHandler_GetAllCharacters(t *testing.T) {
	cacheService := &services.RedisCacheService{}
	cacheService.Connect()
	handler := NewHandler(cacheService, &services.MarvelCharacterAPI{})
	req := httptest.NewRequest(http.MethodGet,
		"/api/v1/characters", nil)
	rec := httptest.NewRecorder()
	c := EchoCon.NewContext(req, rec)
	if error := handler.GetAllCharacters(c); error != nil {
		t.Errorf(error.Error())
	}
	if rec.Code != http.StatusOK {
		t.Errorf("Invalid status code %d", rec.Code)
	}
}