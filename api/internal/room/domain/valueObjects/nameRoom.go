package valueObjects

import "errors"

type NameRoom struct {
	value string
}

func NewNameRoom(value string) (NameRoom, error) {
	if len(value) == 0 {
		return NameRoom{}, errors.New("el nombre de la sala no puede estar vacio")
	}
	if len(value) > 15 {
		return NameRoom{}, errors.New("el nombre de la sala debe tener menos de 15 caracteres")
	}
	return NameRoom{value: value}, nil
}
