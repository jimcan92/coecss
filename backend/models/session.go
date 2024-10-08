package models

import "time"

type Session struct {
	ID         string    `json:"id"`
	UserID     string    `json:"userId"`
	Active     bool      `json:"active"`
	Created    time.Time `json:"created"`
	Terminated time.Time `json:"terminated"`
}

func (s Session) GetID() string {
	return s.ID
}
