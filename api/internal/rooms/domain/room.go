package domain

import (
	v "suffgo/internal/room/domain/valueObjects"
	s "suffgo/internal/shared/domain/valueObjects"
	
)

type (
	Room struct {
		ID         s.ID
		Name       v.NameRoom
		Admin      s.ID
		Linkinvite v.LinkInvite //crear
	}

	RoomDTO struct {
		ID         uint      `json:"id"`
		Name       string   `json:"name"`
		Admin      string   `json:"admin_id"`
		Linkinvite bool     `json:"link_invite"`
	}

	RoomCreateRequest struct {
		Name       string   `json:"name"`
		AdminID    int      `json:"admin_id"` 
		LinkInvite bool     `json:"link_invite"`
	}
)
