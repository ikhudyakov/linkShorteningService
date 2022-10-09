package repo

import (
	"math/rand"
	"time"
)

type Link struct {
	FullLink     string `json:"link"`
	ShortLink    string `json:"shortlink"`
	Domain       int    `json:"domain"`
	lenShortLink int    `json:"-"`
}

func GetLink() *Link {
	return &Link{
		lenShortLink: 10,
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func (l *Link) Generate() string {
	b := make([]rune, l.lenShortLink)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
