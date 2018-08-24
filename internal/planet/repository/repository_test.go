package repository_test

import (
	"testing"
	"time"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet/repository"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/db/mocks"
)

func TestMongoRepository_CreatePlanet_Successfully(t *testing.T) {
	convey.Convey("Give a valid configuration", t, func() {
		planet := &planet.Planet{Name: "Test", Climate: "Test", Appearances: 0, Terrain: "Teste", CreatedAt: time.Now(), UpdatedAt: time.Now()}

		mockSession := &mocks.Session{}
		mockDL := &mocks.DataLayer{}
		mockCollection := &mocks.Collection{}

		repository := repository.NewMongoRepository(mockSession)

		mockSession.On("Clone").Return(mockSession)
		mockSession.On("Close").Return()
		mockSession.On("DB", mock.AnythingOfType("string")).Return(mockDL)
		mockDL.On("C", mock.AnythingOfType("string")).Return(mockCollection)
		mockCollection.On("EnsureIndex", mgo.Index{Key: []string{"name"}, Unique: true}).Return(nil)
		mockCollection.On("Insert", planet).Return(nil)

		convey.Convey("When try create in the database", func() {
			_, err := repository.Store(planet)

			convey.Convey("Should not return error on repository", func() {
				convey.So(err, convey.ShouldBeNil)
			})
		})
	})
}

func TestMongoRepository_RemovePlanet_Successfully(t *testing.T) {
	convey.Convey("Give a valid configuration", t, func() {

		id := bson.NewObjectId()

		mockSession := &mocks.Session{}
		mockDL := &mocks.DataLayer{}
		mockCollection := &mocks.Collection{}

		repository := repository.NewMongoRepository(mockSession)

		mockSession.On("Clone").Return(mockSession)
		mockSession.On("Close").Return()
		mockSession.On("DB", mock.AnythingOfType("string")).Return(mockDL)
		mockDL.On("C", mock.AnythingOfType("string")).Return(mockCollection)
		mockCollection.On("Remove", bson.M{"_id": id}).Return(nil)

		convey.Convey("When try remove in the database", func() {
			err := repository.Remove(id.Hex())

			convey.Convey("Should not return error on repository", func() {
				convey.So(err, convey.ShouldBeNil)
			})
		})
	})
}
