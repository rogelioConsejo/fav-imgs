package main

import (
	"fav-imgs/gallery"
	"fav-imgs/gallery/interfaces"
	"fmt"
	"net/http"
)

func listImages() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		imageGallery := gallery.GetGallery(gallery.GetPersistence())
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
	return "<html>" + output + "</html>"
}

func formatGallery(gallery interfaces.Gallery) string {
	images := gallery.ImageList()
	formattedImages := ""
	for _, image := range images {
		formattedImages += addFormattedImage(image)
	}
	return formattedImages
}

func addFormattedImage(image interfaces.Image) string {
	title := image.GetTitle()
	url := image.GetUrl()
	return "<figure>" + addTitle(title) + addImage(url) + "</figure>"
}

func addTitle(title string) string {
	return "<figcaption>" + title + "</figcaption>"
}

func addImage(url string) string {
	return "<img src=\"" + url + "\">"
}
