package server

import (
	"fav-imgs/gallery"
	"fav-imgs/gallery/image"
	"fav-imgs/gallery/persistence"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AddImage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			htmlTemplate, _ := ioutil.ReadFile("templates/add-image.html")
			fmt.Fprint(w, string(htmlTemplate))
			break
		case http.MethodPost:
			adder := gallery.GetImageAdder(persistence.GetPersistenceAdder())
			err := r.ParseForm()
			println(r.FormValue("title"))
			if err != nil {
				return
			}
			adder.Add(image.NewImage(r.FormValue("title"), r.FormValue("url")))
			http.Redirect(w, r, r.Referer(), 301)
			break
		}
	}
}
