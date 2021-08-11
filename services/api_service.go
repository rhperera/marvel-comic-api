package services

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/rhperera/marvel-comic-api/config"
	"github.com/rhperera/marvel-comic-api/domain"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const EndPoint string = "http://gateway.marvel.com/v1/public"


type IComicCharacterAPI interface {
	GetByID(id int) (*domain.ComicCharacter, *domain.Error)
	GetAllIDs(offset int) ([]int, *domain.Error)
	GetCharacterCount() int
}

// implements the IComicCharacter interface
type MarvelCharacterAPI struct {
	mu sync.Mutex
}

func (mcAPI *MarvelCharacterAPI) GetCharacterCount() int {
	urlString := EndPoint + "/characters" + "?" + mcAPI.getRequestAuthArgs() +
		"&limit=1"
	charRes, err := mcAPI.getAPIResponse(urlString + "&offset=" + strconv.Itoa(0))
	if err != nil || charRes == nil{
		return 0
	}
	return charRes.Total
}

func (mcAPI *MarvelCharacterAPI) getAPIResponse(url string) (*charactersResult, *domain.Error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return nil, &domain.Error{
			Code:    1,
			Message: "Error fetching data from server",
		}
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Error("API fetch failed " + resp.Status)
		return nil, &domain.Error{
			Code:    1,
			Message: "Error fetching data from server",
		}
	}
	body, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		log.Error(ioErr)
		return nil, &domain.Error{
			Code:    1,
			Message: "Error fetching data from server",
		}
	}
	apiResp := apiResponse{}
	jMErr := json.Unmarshal(body, &apiResp)
	if jMErr != nil {
		log.Error(jMErr)
		return nil, &domain.Error{
			Code:    1,
			Message: "Error fetching data from server",
		}
	}
	charRes := apiResp.Data
	if charRes.Count == 0 || len(charRes.Results) == 0 {
		log.Warnf("No character found for request %s", url)
		return nil, &domain.Error{
			Code:    2,
			Message: "No data found for given Id",
		}
	}
	return &charRes, nil
}

func (mcAPI *MarvelCharacterAPI) GetByID(id int) (*domain.ComicCharacter, *domain.Error) {
	urlString := EndPoint + "/characters/" + strconv.Itoa(id) + "?" + mcAPI.getRequestAuthArgs()

	charRes, err := mcAPI.getAPIResponse(urlString)
	if err != nil {
		return nil, err
	}

	comicCharacter := &domain.ComicCharacter{
		Name: charRes.Results[0].Name,
		Id: charRes.Results[0].Id,
		Description: charRes.Results[0].Description,
	}
	return comicCharacter, nil
}

func (mcAPI *MarvelCharacterAPI) GetAllIDs(offset int) ([]int, *domain.Error) {
	// We first fetch the first 100 and get the full count from it
	limit := 100
	urlString := EndPoint + "/characters" + "?" + mcAPI.getRequestAuthArgs() +
		"&limit=" + strconv.Itoa(limit)
	charRes, err := mcAPI.getAPIResponse(urlString + "&offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	var idsArray []int
	mcAPI.addToArray(&idsArray, charRes)

	countsTobeFetched := charRes.Total - offset

	itersNeeded := countsTobeFetched/limit
	if (itersNeeded > 0) && (countsTobeFetched % limit != 0) {
		itersNeeded += 1
	}
	// We now need to invoke (fullCount-100)/100 requests to get all IDs
	waitGroup := sync.WaitGroup{}

	for i := 1; i < itersNeeded; i++ {
		waitGroup.Add(1)
		// call api requests concurrently
		go func(iteration int) {
			defer waitGroup.Done()
			charRes, err := mcAPI.getAPIResponse(urlString + "&offset=" + strconv.Itoa((limit*iteration) + offset))
			if err != nil {
				return
			}
			mcAPI.mu.Lock()
			mcAPI.addToArray(&idsArray, charRes)
			mcAPI.mu.Unlock()
		}(i)

	}
	waitGroup.Wait()
	return idsArray, nil
}

func (mcAPI *MarvelCharacterAPI) addToArray(idsArray *[]int, charRes *charactersResult) {
	limit := len(charRes.Results)
	for i := 0; i < limit; i ++ {
		*idsArray = append(*idsArray, charRes.Results[i].Id)
	}
}

func (mcAPI *MarvelCharacterAPI) getRequestAuthArgs() string {
	ts := time.RFC3339
	privateKey := config.Get("PRIVATE_KEY")
	publicKey := config.Get("PUBLIC_KEY")
	hashBytes := md5.Sum([]byte(ts+privateKey+publicKey))
	return "ts=" + ts + "&apikey=" + publicKey + "&hash=" + hex.EncodeToString(hashBytes[:])
}

type singleCharacter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type charactersResult struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
	Count  int `json:"count"`
	Results []singleCharacter `json:"results"`
}

type apiResponse struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Data   charactersResult `json:"data"`
}
