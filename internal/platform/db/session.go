package db

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/enviroment"
)

func NewSession() (Session, error) {
	db, err := mgo.Dial(enviroment.Conf.Db.Mongo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("An error occurred while trying to open connection with database . %v", err))
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("An error occurred while trying to verify database connection. %v", err))
	}

	return MongoSession{db}, nil
}
