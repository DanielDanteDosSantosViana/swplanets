package client

import (
	"encoding/json"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/enviroment"
	"github.com/parnurzeal/gorequest"
)

type Client interface {
	GetNumberOfAppearancesByPlanetName(search string) (int, error)
}

type SWApi struct{}

type SearchPlanetResult struct {
	Results []Result `json:"results"`
}

type Result struct {
	Name  string   `json:"name"`
	Films []string `json:"films"`
}

func NewSWApi() *SWApi {
	return &SWApi{}
}

type ExternalApiInternalError struct {
	message string
}

func NewExternalApiInternalError(message string) *ExternalApiInternalError {
	return &ExternalApiInternalError{
		message: message,
	}
}
func (e *ExternalApiInternalError) Error() string {
	return e.message
}

type BadMsgError struct {
	message    string
	statusCode int `json:"_"`
}

func NewBadMsgError(message string, statusCode int) *BadMsgError {
	return &BadMsgError{message: message, statusCode: statusCode}
}
func (e *BadMsgError) Error() string {
	return e.message
}
func (e *BadMsgError) StatusCode() int {
	return e.statusCode
}

func (s *SWApi) GetNumberOfAppearancesByPlanetName(search string) (int, error) {

	total := 0
	resp, body, errs := gorequest.New().Get(enviroment.Conf.ExternalAPI.Url + search).End()

	if len(errs) > 0 {
		return total, errs[0]
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return total, NewBadMsgError("Bad msg send to API", resp.StatusCode)
	} else if resp.StatusCode >= 500 {
		return total, NewExternalApiInternalError("Internal error external api")
	}

	searchPlanetResult := &SearchPlanetResult{}

	json.Unmarshal([]byte(body), searchPlanetResult)

	if len(searchPlanetResult.Results) > 0 {
		for _, planetDetails := range searchPlanetResult.Results {
			if planetDetails.Name == search {
				total = len(planetDetails.Films)
			}
		}
	}
	return total, nil
}
