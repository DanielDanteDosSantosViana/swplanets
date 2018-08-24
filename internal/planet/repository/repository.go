package repository

import (
	models "github.com/DanielDanteDosSantosViana/swplanets/internal/planet"
)

type PlanetRepository interface {
	Store(planet *models.Planet) (*models.Planet, error)
	Remove(id string) error
	List() ([]models.Planet, error)
	GetById(id string) (*models.Planet, error)
	GetByName(name string) ([]models.Planet, error)
}
