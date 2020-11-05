package core

import "github.com/louisevanderlith/husk/validation"

type Property struct {
	Bedrooms  int
	Bathrooms int
	Garages   int
	Pool      bool
	Address   string
	Size      string
	Type      string
}

func (p Property) Valid() error {
	return validation.Struct(p)
}
