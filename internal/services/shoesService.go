package services

import (
	"errors"
	"strings"
	"time"

	"github.com/fbomrl/_Corrida_go/internal/model"
	"github.com/fbomrl/_Corrida_go/internal/repository/interfaces"
)

var (
	ErrShoesNotFound = errors.New("Tênis não encontrado")
	errEmptyName     = errors.New("Nome do tênis é obrigatório")
	errNegativeKM    = errors.New("Não pode ter quilometragem negativa")
	errFutureDate    = errors.New("Tênis não pode ser adquirido em data futura")
)

type ShoesService struct {
	repoShoes interfaces.ShoesRepositoryInterface
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
	return s.repoShoes.CreateShoes(shoes)
}
