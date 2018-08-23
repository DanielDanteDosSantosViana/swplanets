package repository

import (
	"strings"
	"time"

	models "github.com/DanielDanteDosSantosViana/swplanets/internal/planet"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/db"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/enviroment"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoRepository struct {
	Session db.Session
}

func NewMongoRepository(session db.Session) *mongoRepository {
	return &mongoRepository{session}
}

type DuplicationKeyError struct {
	message string
}

func NewDuplicationKeyError(message string) *DuplicationKeyError {
	return &DuplicationKeyError{
		message: message,
	}
}
func (e *DuplicationKeyError) Error() string {
	return e.message
}

var (
	collection = "planets"
)

func (m *mongoRepository) Store(planet *models.Planet) (*models.Planet, error) {
	session := m.Session.Clone()
	defer session.Close()
	c := m.getCollection(collection)

	index := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}

	err := c.EnsureIndex(index)

	planet.ID = bson.NewObjectId()
	planet.UpdatedAt = time.Now()
	planet.CreatedAt = time.Now()
	err = c.Insert(planet)

	if err != nil {
		if strings.Contains(err.Error(), "E11000") {
			return planet, NewDuplicationKeyError("Duplication key 'name' ")
		}

		log.WithFields(log.Fields{"error": err}).Error(err.Error())
	}

	return planet, err
}

func (m *mongoRepository) getCollection(collectionName string) db.Collection {
	return m.Session.DB(enviroment.Conf.Db.Name).C(collectionName)
}
