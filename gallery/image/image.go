package image

import (
	"fav-imgs/gallery/interfaces"
)

type image struct {
	title string
	url   string
}

func (i image) GetTitle() string {
	return i.title
}

func (i image) GetUrl() string {
	return i.url
}

func NewImage(title string, url string) interfaces.Image {
	return image{title: title, url: url}
}
