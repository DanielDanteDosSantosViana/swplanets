package database

import (
	env "github.com/DanielDanteDosSantosViana/swplanets/internal/platform/enviroment"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewSessionMysqlWriteDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", env.Conf.Db.MysqlWrite)
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	return db, nil
}
