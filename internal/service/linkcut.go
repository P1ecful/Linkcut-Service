package service

import "net/http"

type LinkCutService interface {
	Cut(link string) string
}

type LinkcutService struct{}

func (LinkcutService) Cut(link string) string {

	http.HandleFunc("/cutted", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, link, http.StatusSeeOther)
	})

	return "localhost:1200/cutted"
}
