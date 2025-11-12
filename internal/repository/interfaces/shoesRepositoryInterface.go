package interfaces

import "github.com/fbomrl/_Corrida_go/internal/model"

type ShoesRepositoryInterface interface {
	CreateShoes(shoes model.Shoes) error
	FindShoesById(id int) (*model.Shoes, error)
	FindAllShoes() ([]*model.Shoes, error)
	UpdateShoes(shoes model.Shoes) error
	UpdateShoesKm(id int, totalKm float64) error
}
