package main

import (
	"Handmans/internal/entity"
	"Handmans/internal/repository/map_repository"
	"Handmans/internal/use_case/hangman"
)

func main() {
	r := map_repository.New()
	r.AddWords(entity.Dictionary{Word: "zebra", Theme: "animal"})
	r.AddWords(entity.Dictionary{Word: "footbal", Theme: "sport"})
	r.AddWords(entity.Dictionary{Word: "black", Theme: "color"})
	r.AddWords(entity.Dictionary{Word: "cat", Theme: "animal"})
	r.AddWords(entity.Dictionary{Word: "dog", Theme: "animal"})
	r.AddWords(entity.Dictionary{Word: "yellow", Theme: "color"})
	r.AddWords(entity.Dictionary{Word: "lion", Theme: "animal"})
	r.AddWords(entity.Dictionary{Word: "beysball", Theme: "sport"})
	r.AddWords(entity.Dictionary{Word: "white", Theme: "color"})
	r.AddWords(entity.Dictionary{Word: "hockey", Theme: "sport"})
	r.AddWords(entity.Dictionary{Word: "hourse", Theme: "animal"})
	r.AddWords(entity.Dictionary{Word: "red", Theme: "color"})
	r.AddWords(entity.Dictionary{Word: "volleyball", Theme: "sport"})

	h := hangman.New(r)
	h.Run()

}
