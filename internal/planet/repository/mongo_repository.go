package repository

import (
	models "github.com/DanielDanteDosSantosViana/swplanets/internal/planet"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/db"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/enviroment"
	"gopkg.in/mgo.v2/bson"
)

type mongoRepository struct {
	Session db.Session
}

func NewMongoRepository(session db.Session) *mongoRepository {
	return &mongoRepository{session}
}

var collection = "planets"

func (m *mongoRepository) Store(planet *models.Planet) (*models.Planet, error) {
	session := m.Session.Clone()
	defer session.Close()
	collection := m.getCollection(collection)
	planet.ID = bson.NewObjectId()
	err := collection.Insert(planet)
	return planet, err
}

func (m *mongoRepository) getCollection(collectionName string) db.Collection {
	return m.Session.DB(enviroment.Conf.Db.Name).C(collectionName)
}
