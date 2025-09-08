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
		"INSERT INTO Running (Name) VALUES (?)",
		running.Name,
	)
	return nil
}
