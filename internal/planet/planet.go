package planet

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Planet struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name        string        `json:"name" validate:"required"`
	Climate     string        `json:"climate" validate:"required"`
	Terrain     string        `json:"terrain" validate:"required"`
	Appearances int           `json:"appearances"`
	UpdatedAt   time.Time     `json:"updated_at"`
	CreatedAt   time.Time     `json:"created_at"`
}
