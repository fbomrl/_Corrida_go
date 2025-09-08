package model

import "time"

type Running struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Local    string    `json:"local"`
	Date     time.Time `json:"date"`
	Distance float64   `json:"distanceKm"`
	Hour     int       `json:"hour"`
	Minute   int       `json:"minute"`
	Second   int       `json:"second"`
	Pace     float64   `json:"pace"`
	Event    bool      `json:"event"`
	Image    string    `json:"image"`
	ShoesId  int       `json:"shoesId"`
}
