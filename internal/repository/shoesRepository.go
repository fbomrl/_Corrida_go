package repository

import (
	"database/sql"

	"github.com/fbomrl/_Corrida_go/internal/model"
)

type ShoesRepository struct {
	DB *sql.DB
}

func (repo *ShoesRepository) CreateShoes(shoes model.Shoes) error {
	_, err := repo.DB.Exec(
		"INSERT INTO Shoes (Name, TotalKm, Bought, Retired, ShoesImage) VALUES(?,?,?,?,?)",
		shoes.Name,
		shoes.TotalKm,
		shoes.Bought,
		shoes.Retired,
		shoes.ShoesImage,
	)
	return err
}
