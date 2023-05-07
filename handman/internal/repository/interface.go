package repository

import "Handmans/internal/entity"

type Repository interface {
	AddWords(in entity.Dictionary) error
	GetWord(in string) (string, error)
	GetThems() []string
}
