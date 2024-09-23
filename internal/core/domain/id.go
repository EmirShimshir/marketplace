package domain

import "github.com/google/uuid"

type ID string

func (id ID) String() string {
	return string(id)
}

func NewID() ID {
	return ID(uuid.New().String())
}
