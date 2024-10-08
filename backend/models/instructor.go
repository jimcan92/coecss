package models

import "github.com/google/uuid"

type Instructor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (i Instructor) GetID() string {
	return i.ID
}

func NewInstructor(name string) *Instructor {
	return &Instructor{
		ID:   uuid.NewString(),
		Name: name,
	}
}
