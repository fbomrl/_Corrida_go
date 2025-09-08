package repository

import (
	"database/sql"

	"github.com/fbomrl/_Corrida_go/internal/model"
)

type RunningRepository struct {
	DB *sql.DB
}

func (repo *RunningRepository) CreateRunning(running model.Running) error {
	_, err := repo.DB.Exec(
		"INSERT INTO Running (Name, Local, Date, Distance, Hour, Minute, Second, Pace, Event, Image, ShoesId) VALUES (?,?,?,?,?,?,?,?,?,?,?)",
		running.Name,
		running.Local,
		running.Date,
		running.Distance,
		running.Hour,
		running.Minute,
		running.Second,
		running.Pace,
		running.Event,
		running.Image,
		running.ShoesId,
	)
	return err
}
