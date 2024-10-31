package valueObjects

import "errors"

type Admin struct {
	ID       int
	Username string
}

// NewAdmin valida y crea una nueva instancia de Admin
func NewAdmin(id int, username string) (Admin, error) {
	if id <= 0 {
		return Admin{}, errors.New("el ID del administrador debe ser un número positivo")
	}
	if len(username) == 0 {
		return Admin{}, errors.New("el nombre de usuario del administrador no puede estar vacío")
	}
	if len(username) > 20 {
		return Admin{}, errors.New("el nombre de usuario debe tener menos de 20 caracteres")
	}
	return Admin{ID: id, Username: username}, nil
}
