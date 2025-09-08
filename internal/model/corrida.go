package model

import "time"

type Corrida struct {
	Id       int       `json: id`
	Name     string    `json: name`
	City     string    `json: city`
	Date     time.Time `json: date`
	Distance float64   `json: distance`
	hour     int       `json: hour`
	minute   int       `json: minute`
	second   int       `json: second`
	event    bool      `json: event`
}
