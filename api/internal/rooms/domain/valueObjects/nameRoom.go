package valueObjects

import "errors"

type NameRoom struct {
	name string
}

func NewNameRoom(nameRoom string) (*NameRoom, error) {
	if nameRoom == "" {
		return nil, errors.New("no se puede crear una sala sin nombre")
	}
	return &NameRoom{name: nameRoom}, nil
}
