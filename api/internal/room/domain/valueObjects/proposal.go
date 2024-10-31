package valueObjects

import "errors"

type Proposal struct {
	value string
}

// NewProposal valida la propuesta y devuelve una instancia de Proposal
func NewProposal(value string) (Proposal, error) {
	if len(value) == 0 {
		return Proposal{}, errors.New("la propuesta no puede estar vacía")
	}
	if len(value) > 100 {
		return Proposal{}, errors.New("la propuesta debe tener menos de 100 caracteres")
	}
	return Proposal{value: value}, nil
}
