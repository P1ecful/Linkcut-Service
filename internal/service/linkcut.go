package service

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type LinkCutService interface {
	Cut(link string) string
}

type LinkcutService struct{}

func (LinkcutService) Cut(link string) string {
	shortedUrl := "/cutted/" + strconv.Itoa(rand.Intn(99)) + "/"

	router := mux.NewRouter()
	router.HandleFunc(shortedUrl, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, link, http.StatusSeeOther)
	})

	return "localhost:1200" + shortedUrl
}
