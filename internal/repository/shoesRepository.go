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

func (repo *ShoesRepository) FindShoesById(id int) (*model.Shoes, error) {
	var shoes model.Shoes

	err := repo.DB.QueryRow("SELECT Id, Name, TotalKm, Bought, Retired, ShoesImage FROM Shoes WHERE Id = ?", id).Scan(
		&shoes.Id,
		&shoes.Name,
		&shoes.TotalKm,
		&shoes.Bought,
		&shoes.Retired,
		&shoes.ShoesImage,
	)

	if err != nil {
		return nil, err
	}

	return &shoes, nil
}

func (repo *ShoesRepository) FindAllShoes() ([]*model.Shoes, error) {
	rows, err := repo.DB.Query("SELECT Id, Name, TotalKm, Bought, Retired, ShoesImage FROM Shoes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allShoes []*model.Shoes

	for rows.Next() {
		var shoes model.Shoes
		err := rows.Scan(
			&shoes.Id,
			&shoes.Name,
			&shoes.TotalKm,
			&shoes.Bought,
			&shoes.Retired,
			&shoes.ShoesImage,
		)

		if err != nil {
			return nil, err
		}

		allShoes = append(allShoes, &shoes)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return allShoes, nil
}
