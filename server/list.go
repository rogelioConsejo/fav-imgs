package server

import (
	"fav-imgs/gallery"
	"fav-imgs/gallery/interfaces"
	"fav-imgs/gallery/persistence"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ListImages() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		imageGallery := gallery.GetReader(persistence.GetPersistenceReader())

		output := formatGallery(imageGallery)
		output = addGlobalHtml(output)
		_, err := fmt.Fprint(w, output)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func addGlobalHtml(output string) string {
	htmlTemplate, err := ioutil.ReadFile("templates/basic-template.html")
	if err != nil {
		fmt.Println("error adding global html: " + err.Error())
	}
	template := string(htmlTemplate)
	templateWithTitle := strings.ReplaceAll(template, "${TITLE}", "Favorite Images List")
	filledTemplate := strings.ReplaceAll(templateWithTitle, "${BODY}", output)
	return filledTemplate
}

func formatGallery(gallery interfaces.GalleryReader) string {
	images := gallery.ImageList()
	formattedImages := ""
	for key, image := range images {
		formattedImages += addFormattedImage(image, key)
	}
	return formattedImages
}

func addFormattedImage(image interfaces.Image, key string) string {
	title := image.GetTitle()
	url := image.GetUrl()
	return "<figure>" + addTitle(title) + addImage(url) + addDeleteButton(key) + "</figure>"
}

func addDeleteButton(key string) string {
	return "<a href=\"http://localhost:8080/delete?id=" + key + "\"> Delete Image </a>"
}

func addTitle(title string) string {
	return "<figcaption>" + title + "</figcaption>"
}

func addImage(url string) string {
	return "<img src=\"" + url + "\">"
}
