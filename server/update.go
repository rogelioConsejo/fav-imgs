package server

import (
	"fav-imgs/gallery"
	"fav-imgs/gallery/image"
	"fav-imgs/gallery/persistence"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func UpdateImage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			performUpdate(w, r)
			break
		case http.MethodGet:
			missing := checkForId(r)
			missing = checkForTitle(r)
			missing = checkForUrl(r)
			if missing != nil {
				http.Redirect(w, r, root, 301)
				return
			}
			showInput(w, r)
			break
		}
	}
}

func performUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("error updating image: invalid parameter(s):%s\n", err.Error())
		return
	}
	id := r.FormValue(idKey)
	title := r.FormValue(titleKey)
	url := r.FormValue(urlKey)

	updater := gallery.GetImageModifier(persistence.GetPersistenceUpdater())
	updater.Update(id, image.NewImage(title, url))

	http.Redirect(w, r, root, 301)
}

func showInput(w http.ResponseWriter, r *http.Request) {

	templateFileContents, err := ioutil.ReadFile("templates/update-image.html")
	err = r.ParseForm()
	if err != nil {
		fmt.Println("error adding global html: " + err.Error())
	}

	template := string(templateFileContents)
	templateWithId := strings.ReplaceAll(template, "${IMAGE_ID}", r.FormValue(idKey))
	templateWithUrl := strings.ReplaceAll(templateWithId, "${IMAGE_URL}", r.FormValue(urlKey))
	filledTemplate := strings.ReplaceAll(templateWithUrl, "${IMAGE_TITLE}", r.FormValue(titleKey))

	_, err = fmt.Fprint(w, filledTemplate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
