package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (gb *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gb.l.Println("Goodbye")
}
