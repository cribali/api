package domain

import (
	sv "suffgo/internal/shared/domain/valueObjects"
	v "suffgo/internal/user/domain/valueObjects"
)

type RoomRepository interface {
	GetByID(id sv.ID) (*Room, error)
	GetAll() ([]Room, error)
	Delete(id sv.ID) error
	GetByEmail(email v.Email) (*Room, error)
	Save(user Room) error
	// Update(user User) error
}
