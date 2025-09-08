package model

import "time"

type Shoes struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	TotalKm    float64    `json:"totalKm"`
	Bought     time.Time  `json:"bought"`
	Retired    *time.Time `json:"retired,omitempty"`
	ShoesImage string     `json:"image"`
}
