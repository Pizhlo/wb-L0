package models

import "time"

type User struct {
	ID       int
	Login    string
	Password string
	Created  time.Time
}
