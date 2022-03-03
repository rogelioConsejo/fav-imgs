package server

import (
	"fav-imgs/gallery"
	"fav-imgs/gallery/persistence"
	"net/http"
)

func DeleteImage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		err = checkForId(r)
		if err != nil {
			return
		}

		id, err := getId(r)
		if err != nil {
			return
		}

		deleter := gallery.GetImageDeleter(persistence.GetPersistenceDeleter())
		deleter.Delete(id)
		http.Redirect(w, r, r.Referer(), 301)
	}
}
