package map_repository

import (
	"Handmans/internal/entity"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type repository struct {
	storage map[string][]string
}

func New() *repository {
	return &repository{
		storage: make(map[string][]string),
	}
}

func (r *repository) AddWords(in entity.Dictionary) error {
	value := r.storage[in.Theme]
	value = append(value, in.Word)
	r.storage[in.Theme] = value
	return nil
}

func (r *repository) GetWord(in string) (string, error) {
	value, ok := r.storage[in]
	if ok != true {
		return "", fmt.Errorf("don't exist them: %s", in)
	}
	index := rand.Intn(len(value))
	return value[index], nil
}

func (r *repository) GetThems() []string {
	temsMix := []string{}
	for key := range r.storage {
		temsMix = append(temsMix, key)
	}
	return temsMix
}
