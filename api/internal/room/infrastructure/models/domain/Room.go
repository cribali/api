package domain

import (
	v "suffgo/internal/room/infrastructure/models/domain/valueObjects"
)

type (
	Room struct {
		ID         int
		Name       v.nombreSala
		Admin      v.admin
		Proposal   []v.Proposal
		Linkinvite bool
	}

	RoomDTO struct {
		ID         int      `json:"id"`
		Name       string   `json:"name"`
		Admin      string   `json:"admin"`
		Proposals  []string `json:"proposals"`
		Linkinvite bool     `json:"link_invite"`
	}

	RoomCreateRequest struct {
		Name       string   `json:"name" binding:"required"`
		AdminID    int      `json:"admin_id" binding:"required"` // ID del administrador que creará la sala
		Proposals  []string `json:"proposals,omitempty"`         // Opcionalmente se pueden crear propuestas al inicio
		LinkInvite bool     `json:"link_invite"`
	}
)
