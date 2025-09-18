package services

import (
	"errors"
	"time"

	"github.com/fbomrl/_Corrida_go/internal/model"
	"github.com/fbomrl/_Corrida_go/internal/repository/interfaces"
)

var (
	errFutureDateRunning = errors.New("data não pode superior a data de hoje")
	errEmptyLocal        = errors.New("local não pode ser vazio")
	errInvalidTime       = errors.New("horário inválido")
	errInvalidMinute     = errors.New("minuto inválido")
	errInvalidSecond     = errors.New("segundo inválido")
	errEventIsNUll       = errors.New("evento não pode ser nulo")
	errShoesInvalid      = errors.New("calçado inválido")
	errRunningsNotFound  = errors.New("corrida não localizada")
)

type RunningService struct {
	RepoRunning interfaces.RunningRepositoryInterface
}

func (s *RunningService) CreateRunningService(running model.Running) error {

	if running.Name == "" || len(running.Name) == 0 {
		return errEmptyLocal
	}

	if running.Local == "" || len(running.Local) == 0 {
		return errEmptyLocal
	}

	if running.Date.After(time.Now()) {
		return errFutureDateRunning
	}

	if running.Distance < 0 {
		return errNegativeKM
	}

	if running.Hour < 0 || running.Hour > 23 {
		return errInvalidTime
	}

	if running.Minute < 0 || running.Minute > 59 {
		return errInvalidMinute
	}

	if running.Second < 0 || running.Second > 59 {
		return errInvalidSecond
	}

	if !running.Event {
		return errEventIsNUll
	}

	if running.ShoesId == 0 {
		return errShoesInvalid
	}

	totalSeconds := (running.Hour * 3600) + (running.Minute * 60) + running.Second
	running.Pace = float64(totalSeconds) / 60.0 / running.Distance
	running.AverageSpeed = running.Distance / (float64(totalSeconds) / 3600.0)

	return s.RepoRunning.CreateRunning(running)
}

func (s *RunningService) FindRunningByIdService(id int) (*model.Running, error) {
	running, err := s.RepoRunning.FindRunningById(id)
	if err != nil || running == nil {
		return nil, errRunningsNotFound
	}
	return running, nil
}

func (s *RunningService) FindAllShoesService(id int, name string, local string, date time.Time, distance float64,
	hour int, minute int, second int, pace float64, event bool, image string, shoesId int, averageSpeed float64) ([]*model.Running, error) {
	return s.RepoRunning.FindAllRunnings()
}
