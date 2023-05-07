package hangman

import (
	"Handmans/internal/repository"
	"fmt"
	"strings"
)

type Hangman interface {
	Run()
}

type char interface {
}

type hangman struct {
	repo repository.Repository // нов тип данных
	in   map[string]struct{}
}

func New(repos repository.Repository) *hangman {
	return &hangman{
		repo: repos,
		in:   make(map[string]struct{}),
	}
}

func (h *hangman) Run() {
	fmt.Println("RUN")
	word := h.getWord()
	star := getStar(word)
	fmt.Println(star)
	var l string
	for {
		fmt.Println("Введите букву")
		fmt.Scanln(&l)
		_, ok := h.in[l]
		if ok == true {
			fmt.Println("такая буква уже есть")
			continue
		}
		h.in[l] = struct{}{}

		val := strings.Index(word, l)
		if val == -1 {
			fmt.Println("такой буквы в слове нет")
			continue
		}

		fmt.Println("Такая буква есть")
		star = getSucsses(word, star, l)
		fmt.Println(star)
		if star == word {
			fmt.Println("The winner")
			break
		}
	}
}

func (h *hangman) getWord() string {
	fmt.Println("Выбери тему: ")
	tems := h.repo.GetThems()
	fmt.Println(tems)
	value := ""
	for {
		fmt.Scanln(&value)
		word, err := h.repo.GetWord(value)
		if err != nil {
			fmt.Println("Такой темы нет, попробуй еще")
			continue
		}
		return word
	}
	return ""
}

func getStar(in string) string {
	star := ""
	for i := 0; i < len(in); i++ {
		star += "*"
	}
	return star
}

func getSucsses(in, s, l string) string {
	star := ""
	for j, i := range in {
		if l == string(i) {
			star += string(i)
		} else {
			star += string(s[j])
		}
	}
	return star
}
