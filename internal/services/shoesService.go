package services

import (
	"errors"
	"strings"
	"time"

	"github.com/fbomrl/_Corrida_go/internal/model"
	"github.com/fbomrl/_Corrida_go/internal/repository/interfaces"
)

var (
	errEmptyName  = errors.New("nome do tênis é obrigatório")
	errNegativeKM = errors.New("não pode ter quilometragem negativa")
	errFutureDate = errors.New("tênis não pode ser adquirido em data futura")
)

type ShoesService struct {
	RepoShoes interfaces.ShoesRepositoryInterface
}

func (s *ShoesService) CreateShoesService(shoes model.Shoes) error {
	if strings.TrimSpace(shoes.Name) == "" {
		return errEmptyName
	}
	if shoes.TotalKm < 0 {
		return errNegativeKM
	}
	if shoes.Bought.After(time.Now()) {
		return errFutureDate
	}
	return s.RepoShoes.CreateShoes(shoes)
}
