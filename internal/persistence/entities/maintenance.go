package entities

import "time"

type Maintenance struct {
	ID    int
	CarID int
	Name  string
	Date  time.Time
	KM    int
	Cost  int
}
