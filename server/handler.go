package server

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rhperera/marvel-comic-api/domain"
	"github.com/rhperera/marvel-comic-api/services"
	"net/http"
	"strconv"
)

type Handler struct {
	cacheService services.ICacheService
	apiService services.IComicCharacterAPI
}

// GetCharacterById godoc
// @Tags Characters
// @Summary Get a comic character by Id
// @ID GetCharacterById
// @Param id path int true "Id of marvel character"
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.ComicCharacter
// @Router /characters/{id} [get]
func (handler *Handler) GetCharacterById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, domain.HttpResponse{
			Data:  nil,
		})
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.HttpResponse{
			Data:  nil,
			Error: domain.Error{
				Code:    4,
				Message: "Error getting Id",
			},
		})
	}
	data, dataErr := handler.apiService.GetByID(idInt)
	if dataErr != nil {
		return c.JSON(http.StatusBadRequest, domain.HttpResponse{
			Data:  nil,
			Error: *dataErr,
		})
	}
	return c.JSON(http.StatusOK, domain.HttpResponse{
		Data:  data,
	})
}

// GetAllCharacters godoc
// @Tags Characters
// @Summary Grt Ids of all characters
// @ID GetAllCharacters
// @Accept  json
// @Produce  json
// @Success 200 {array} int
// @Router /characters [get]
func (handler *Handler) GetAllCharacters(c echo.Context) error {
	data, err := handler.cacheService.GetIds()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.HttpResponse{})
	}
	if data != "" {
		count := handler.apiService.GetCharacterCount()
		var idsArray []int
		if err := json.Unmarshal([]byte(data), &idsArray); err != nil {
			return c.JSON(http.StatusInternalServerError, domain.HttpResponse{})
			log.Error(err)
		}

		if len(idsArray) == count {
			// cache count and count from API is same
			return c.JSON(http.StatusOK, domain.HttpResponse{Data: idsArray})
		}
		// Counts in API server and cache mismatch. Get the new ones from API and update cache
		if count > len(idsArray) {
			newIds, newIdsErr := handler.apiService.GetAllIDs(len(idsArray))
			if newIdsErr != nil {
				return c.JSON(http.StatusBadRequest, domain.HttpResponse{
					Data:  nil, Error: *newIdsErr})
			}
			idsArray = append(idsArray, newIds...)
			// Also update the cache
			handler.cacheService.AddIds(&services.IdsHolder{Ids: idsArray})
		}
		return c.JSON(http.StatusOK, domain.HttpResponse{Data:  idsArray})
	}

	// We are loading all data for the first time
	dataFromApi, errFromApi := handler.apiService.GetAllIDs(0)
	if errFromApi != nil {
		return c.JSON(http.StatusBadRequest, domain.HttpResponse{
			Data:  nil,Error: *errFromApi})
	}
	handler.cacheService.AddIds(&services.IdsHolder{Ids: dataFromApi})

	return c.JSON(http.StatusOK, domain.HttpResponse{Data:  dataFromApi})
}

func NewHandler(cS services.ICacheService, apiS services.IComicCharacterAPI) *Handler {
	return &Handler{
		cacheService: cS,
		apiService:   apiS,
	}
}
