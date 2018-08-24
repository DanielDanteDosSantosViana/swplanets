package repository

import (
	"strings"
	"time"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet"

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

type InvalidIdError struct {
	message string
}

func NewInvalidIdError() *InvalidIdError {
	return &InvalidIdError{
		message: "invalid id",
	}
}
func (e *InvalidIdError) Error() string {
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
			log.WithFields(log.Fields{"error": err}).Error(err.Error())
			return planet, NewDuplicationKeyError("Duplication key 'name' ")
		}

		log.WithFields(log.Fields{"error": err}).Error(err.Error())
	}

	return planet, err
}

func (m *mongoRepository) Remove(id string) error {
	session := m.Session.Clone()
	defer session.Close()
	c := m.getCollection(collection)

	if !bson.IsObjectIdHex(id) {
		err := NewInvalidIdError()
		log.WithFields(log.Fields{"error": err}).Error(err.Error())
		return err
	}

	if err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		log.WithFields(log.Fields{"error": err}).Error(err.Error())
		return err
	}
	return nil
}

func (m *mongoRepository) List() ([]planet.Planet, error) {
	session := m.Session.Clone()
	defer session.Close()
	c := m.getCollection(collection)
	planets := []planet.Planet{}
	if err := c.Find(nil).All(&planets); err != nil {
		log.WithFields(log.Fields{"error": err}).Error(err.Error())
		return planets, err
	}
	return planets, nil
}

func (m *mongoRepository) getCollection(collectionName string) db.Collection {
	return m.Session.DB(enviroment.Conf.Db.Name).C(collectionName)
}
