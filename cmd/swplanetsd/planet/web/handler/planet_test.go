package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/client"
	"gopkg.in/mgo.v2/bson"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet/repository"

	"github.com/DanielDanteDosSantosViana/swplanets/cmd/swplanetsd/planet/web/handler"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/db/mocks"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/web"

	"github.com/gorilla/mux"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func TestPlanetHandler_Create_Success(t *testing.T) {
	convey.Convey("Given a valid request with payload", t, func() {
		var err error
		var timeNow time.Time
		validPlanetString := `{"name":"Alderaan","climate":"temperate","terrain":"gas giant"}`

		expectedReturn := `{"id":"","name":"Alderaan","climate":"temperate","terrain":"gas giant","appearances":0,"updated_at":"0001-01-01T00:00:00Z","created_at":"0001-01-01T00:00:00Z"}`

		planetReturn := &planet.Planet{Name: "Alderaan", Climate: "temperate", Terrain: "gas giant", UpdatedAt: timeNow, CreatedAt: timeNow}
		body := strings.NewReader(validPlanetString)

		req, _ := http.NewRequest("POST", "/api/v1/planets", body)
		recorder := httptest.NewRecorder()

		client := &mocks.Client{}
		client.On("GetNumberOfAppearancesByPlanetName", mock.Anything).Return(0, nil)

		planetRepository := &mocks.PlanetRepository{}
		planetRepository.On("Store", mock.Anything).Return(planetReturn, nil)

		planetHandler := handler.NewPlanetHandler(planetRepository, client)

		r := mux.NewRouter().StrictSlash(true)
		api := r.PathPrefix("/api/v1").Subrouter()
		api.HandleFunc("/planets", planetHandler.Create).Methods(http.MethodPost)
		api.ServeHTTP(recorder, req)

		convey.Convey("Should not return error", func() {
			convey.Convey("Status code not different 201", func() {
				if status := recorder.Code; status != http.StatusCreated {
					err = errors.New(fmt.Sprintf("handler returned not expected response. response %v expected %v", status, http.StatusCreated))
				}
				convey.So(err, convey.ShouldBeNil)
				convey.So(recorder.Code, convey.ShouldEqual, http.StatusCreated)
			})
			convey.Convey("Should not empty body", func() {
				if err := web.IsRequestValid(recorder.Body.String()); err != nil {
					err = errors.New(fmt.Sprintf("handler not returned body. expected %v", expectedReturn))
				}
				convey.So(err, convey.ShouldBeNil)
			})
			convey.Convey("Body should be equals expected", func() {
				if expectedReturn != recorder.Body.String() {
					err = errors.New(fmt.Sprintf("handler returned not expected body. body %v expected %v", recorder.Body.String(), expectedReturn))
				}

				convey.So(err, convey.ShouldBeNil)
			})
		})
	})
}

func TestPlanetHandler_Create_MongoDuplicateKeyError(t *testing.T) {
	convey.Convey("Given an invalid request with the existing name", t, func() {
		var err error
		validPlanetString := `{"name":"Alderaan","climate":"temperate","terrain":"gas giant"}`
		expectedResponse := `{"error":"Duplication key 'name' "}`

		body := strings.NewReader(validPlanetString)

		req, _ := http.NewRequest("POST", "/api/v1/planets", body)
		recorder := httptest.NewRecorder()

		client := &mocks.Client{}
		client.On("GetNumberOfAppearancesByPlanetName", mock.Anything).Return(0, nil)

		duplicateKeyError := repository.NewDuplicationKeyError("Duplication key 'name' ")
		planetRepository := &mocks.PlanetRepository{}
		planetRepository.On("Store", mock.Anything).Return(&planet.Planet{}, duplicateKeyError)

		planetHandler := handler.NewPlanetHandler(planetRepository, client)

		r := mux.NewRouter().StrictSlash(true)
		api := r.PathPrefix("/api/v1").Subrouter()
		api.HandleFunc("/planets", planetHandler.Create).Methods(http.MethodPost)
		api.ServeHTTP(recorder, req)

		convey.Convey("Should return error", func() {
			convey.Convey("Status code should be equal 409", func() {
				convey.So(recorder.Code, convey.ShouldEqual, http.StatusConflict)
			})
			convey.Convey("Body should be equals expected ", func() {

				if expectedResponse != recorder.Body.String() {
					err = errors.New(fmt.Sprintf("handler returned not expected body. body %v expected %v", recorder.Body.String(), expectedResponse))
				}
				convey.So(err, convey.ShouldBeNil)
			})
		})
	})
}

func TestPlanetHandler_Create_ExternalAPIReturnBadRequest(t *testing.T) {
	convey.Convey("Given an valid request", t, func() {
		var err error
		validPlanetString := `{"name":"Alderaan","climate":"temperate","terrain":"gas giant"}`
		expectedResponse := `{"error":"Bad msg send to API"}`

		body := strings.NewReader(validPlanetString)

		req, _ := http.NewRequest("POST", "/api/v1/planets", body)
		recorder := httptest.NewRecorder()

		clientMock := &mocks.Client{}
		clientMock.On("GetNumberOfAppearancesByPlanetName", mock.Anything).Return(0, client.NewBadMsgError("Bad msg send to API", http.StatusBadRequest))

		planetRepository := &mocks.PlanetRepository{}
		planetRepository.On("Store", mock.Anything).Return(&planet.Planet{}, nil)

		planetHandler := handler.NewPlanetHandler(planetRepository, clientMock)

		r := mux.NewRouter().StrictSlash(true)
		api := r.PathPrefix("/api/v1").Subrouter()
		api.HandleFunc("/planets", planetHandler.Create).Methods(http.MethodPost)
		api.ServeHTTP(recorder, req)

		convey.Convey("Should return error", func() {
			convey.Convey("Status code should be equal 400", func() {
				convey.So(recorder.Code, convey.ShouldEqual, http.StatusBadRequest)
			})
			convey.Convey("Body should be equals expected ", func() {

				if expectedResponse != recorder.Body.String() {
					err = errors.New(fmt.Sprintf("handler returned not expected body. body %v expected %v", recorder.Body.String(), expectedResponse))
				}
				convey.So(err, convey.ShouldBeNil)
			})
		})
	})
}

func TestPlanetHandler_Create_ExternalAPIReturnInternalError(t *testing.T) {
	convey.Convey("Given an valid request", t, func() {
		var err error
		validPlanetString := `{"name":"Alderaan","climate":"temperate","terrain":"gas giant"}`
		expectedResponse := `{"error":"Internal error external api"}`

		body := strings.NewReader(validPlanetString)

		req, _ := http.NewRequest("POST", "/api/v1/planets", body)
		recorder := httptest.NewRecorder()

		clientMock := &mocks.Client{}
		clientMock.On("GetNumberOfAppearancesByPlanetName", mock.Anything).Return(0, client.NewExternalApiInternalError("Internal error external api"))

		planetRepository := &mocks.PlanetRepository{}
		planetRepository.On("Store", mock.Anything).Return(&planet.Planet{}, nil)

		planetHandler := handler.NewPlanetHandler(planetRepository, clientMock)

		r := mux.NewRouter().StrictSlash(true)
		api := r.PathPrefix("/api/v1").Subrouter()
		api.HandleFunc("/planets", planetHandler.Create).Methods(http.MethodPost)
		api.ServeHTTP(recorder, req)

		convey.Convey("Should return error", func() {
			convey.Convey("Status code should be equal 502", func() {
				convey.So(recorder.Code, convey.ShouldEqual, http.StatusBadGateway)
			})
			convey.Convey("Body should be equals expected ", func() {

				if expectedResponse != recorder.Body.String() {
					err = errors.New(fmt.Sprintf("handler returned not expected body. body %v expected %v", recorder.Body.String(), expectedResponse))
				}
				convey.So(err, convey.ShouldBeNil)
			})
		})
	})
}

func TestPlanetHandler_Remove__Success(t *testing.T) {
	convey.Convey("Given an valid request", t, func() {
		body := strings.NewReader("")
		id := bson.NewObjectId()

		req, _ := http.NewRequest("DELETE", "/api/v1/planets/"+id.Hex(), body)
		recorder := httptest.NewRecorder()

		clientMock := &mocks.Client{}
		clientMock.On("GetNumberOfAppearancesByPlanetName", mock.Anything).Return(1, nil)

		planetRepository := &mocks.PlanetRepository{}
		planetRepository.On("Remove", id.Hex()).Return(nil)

		planetHandler := handler.NewPlanetHandler(planetRepository, clientMock)

		r := mux.NewRouter().StrictSlash(true)
		api := r.PathPrefix("/api/v1").Subrouter()
		api.HandleFunc("/planets/{id}", planetHandler.Remove).Methods(http.MethodDelete)
		api.ServeHTTP(recorder, req)

		convey.Convey("Should not return error", func() {
			convey.Convey("Status code should be equal 204", func() {
				convey.So(recorder.Code, convey.ShouldEqual, http.StatusNoContent)
			})
		})
	})
}

func TestPlanetHandler_Remove__InvalidID(t *testing.T) {
	convey.Convey("Given an valid request", t, func() {
		var err error
		expectedReturn := `{"error":"invalid id"}`
		body := strings.NewReader("")
		id := "id"

		req, _ := http.NewRequest("DELETE", "/api/v1/planets/"+id, body)
		recorder := httptest.NewRecorder()

		clientMock := &mocks.Client{}
		clientMock.On("GetNumberOfAppearancesByPlanetName", mock.Anything).Return(1, nil)

		planetRepository := &mocks.PlanetRepository{}
		planetRepository.On("Remove", id).Return(repository.NewInvalidIdError())

		planetHandler := handler.NewPlanetHandler(planetRepository, clientMock)

		r := mux.NewRouter().StrictSlash(true)
		api := r.PathPrefix("/api/v1").Subrouter()
		api.HandleFunc("/planets/{id}", planetHandler.Remove).Methods(http.MethodDelete)
		api.ServeHTTP(recorder, req)

		convey.Convey("Should not return error", func() {
			convey.Convey("Status code should be equal 400", func() {
				convey.So(recorder.Code, convey.ShouldEqual, http.StatusBadRequest)
			})
		})
		convey.Convey("Body should be equals expected", func() {
			if expectedReturn != recorder.Body.String() {
				err = errors.New(fmt.Sprintf("handler returned not expected body. body %v expected %v", recorder.Body.String(), expectedReturn))
			}

			convey.So(err, convey.ShouldBeNil)
		})
	})
}
