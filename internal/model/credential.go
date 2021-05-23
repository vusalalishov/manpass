package model

import "time"

type Credential struct {
	Id        int64
	Title     string
	Login     string
	Password  string
	UpdatedAt time.Time
}
