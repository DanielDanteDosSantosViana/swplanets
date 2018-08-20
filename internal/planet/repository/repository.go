package repository

import (
	models "github.com/DanielDanteDosSantosViana/swplanets/internal/planet"
)

type PlanetRepository interface {
	Store(planet *models.Planet) (*models.Planet, error)
}
