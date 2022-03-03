package server

import (
	"errors"
	"fav-imgs/gallery"
	"fav-imgs/gallery/persistence"
	"net/http"
	"strings"
)

const idKey = "id"

func DeleteImage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		if !r.URL.Query().Has(idKey) {
			err = errors.New("id not found")
		}
		if err != nil {
			return
		}

		var id = r.URL.Query().Get(idKey)
		if strings.Trim(id, " ") == "" {
			err = errors.New("id parameter incorrect")
		}
		if err != nil {
			return
		}

		deleter := gallery.GetImageDeleter(persistence.GetPersistenceDeleter())
		deleter.Delete(id)
		http.Redirect(w, r, r.Referer(), 301)
	}
}
