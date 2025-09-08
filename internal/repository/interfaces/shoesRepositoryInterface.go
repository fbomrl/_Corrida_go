package interfaces

import "github.com/fbomrl/_Corrida_go/internal/model"

type ShoesRepositoryInterface interface {
	CreateShoes(shoes model.Shoes) error
}
