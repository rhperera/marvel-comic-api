package services

import (
	"github.com/rhperera/marvel-comic-api/config"
	"testing"
)

func TestGetByID_ValidID(t *testing.T) {
	config.InitForTests()
	apiService := &MarvelCharacterAPI{}
	resp, err := apiService.GetByID(1011334)
	if err != nil {
		t.Error(err.Message)
	}
	if resp == nil {
		t.Error("Fetching character failed")
	}
	if resp.Id != 1011334 {
		t.Error("Fetching character failed")
	}
}

func TestGetByID_IdAsZero(t *testing.T) {
	config.InitForTests()
	apiService := &MarvelCharacterAPI{}
	_, err := apiService.GetByID(0)
	if err == nil {
		t.Errorf("Error must be returned")
	}
}

func TestGetByID_IdNotFound(t *testing.T) {
	config.InitForTests()
	apiService := &MarvelCharacterAPI{}
	_, err := apiService.GetByID(1111111111888888)
	if err == nil {
		t.Errorf("Error must be returned")
	}
}

func TestMarvelCharacterAPI_GetAllIDs(t *testing.T) {
	config.InitForTests()
	apiService := &MarvelCharacterAPI{}
	resp, err := apiService.GetAllIDs(0)
	if err != nil {
		t.Errorf("Error must be returned")
	}
	if len(resp) < 1 {
		t.Error("IDs array is empty")
	}
}

func TestMarvelCharacterAPI_getAPIResponse_emptyURL(t *testing.T) {
	config.InitForTests()
	apiService := &MarvelCharacterAPI{}
	res, err := apiService.getAPIResponse("");
	if err == nil || res != nil  {
		t.Error("No error returned")
	}
}

func TestMarvelCharacterAPI_GetCharacterCount(t *testing.T) {
	config.InitForTests()
	apiService := &MarvelCharacterAPI{}
	count := apiService.GetCharacterCount()
	if count < 1 {
		t.Errorf("Error count is less than 1. count-%d", count)
	}
}