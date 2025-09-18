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
		"INSERT INTO Running (Name, Local, Date, Distance, Hour, Minute, Second, Pace, Event, Image, ShoesId, AverageSpeed)"+
			"VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12)",
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
		running.AverageSpeed,
	)
	return err
}

func (repo *RunningRepository) FindRunningById(id int) (*model.Running, error) {
	var running model.Running

	err := repo.DB.QueryRow("SELECT * FROM Running WHERE Id = ?", id).Scan(
		&running.Name,
		&running.Local,
		&running.Date,
		&running.Distance,
		&running.Hour,
		&running.Minute,
		&running.Second,
		&running.Pace,
		&running.Event,
		&running.Image,
		&running.ShoesId,
		&running.AverageSpeed,
	)
	if err != nil {
		return nil, err
	}
	return &running, nil
}

func (repo *RunningRepository) FindAllRunnings() ([]*model.Running, error) {
	rows, err := repo.DB.Query("SELECT Name, Local, Date, Distance, Hour, Minute, Second, Pace, Event, Image, ShoesId, AverageSpeed FROM Running")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allRunnings []*model.Running

	for rows.Next() {
		var running model.Running
		err := rows.Scan(
			&running.Name,
			&running.Local,
			&running.Date,
			&running.Distance,
			&running.Hour,
			&running.Minute,
			&running.Second,
			&running.Pace,
			&running.Event,
			&running.Image,
			&running.ShoesId,
			&running.AverageSpeed,
		)
		if err != nil {
			return nil, err
		}
		allRunnings = append(allRunnings, &running)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return allRunnings, nil
}
