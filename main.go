package main

import (
	. "fav-imgs/gallery"
	. "fav-imgs/gallery/interfaces"
	"fmt"
)

func main() {
	gallery := GetGallery(GetPersistence())
	output := formatGallery(gallery)
	fmt.Println(output)
}

func formatGallery(gallery Gallery) string {
	images := gallery.ImageList()
	formattedImages := ""
	for _, image := range images {
		formattedImages += addFormattedImage(image)
	}
	return formattedImages
}

func addFormattedImage(image Image) string {
	title := image.GetTitle()
	url := image.GetUrl()
	return "<div>" + addTitle(title) + addImage(url) + "</div>"
}

func addTitle(title string) string {
	return "<p>" + title + "</p>"
}

func addImage(url string) string {
	return "<img src=\"" + url + "\">"
}
