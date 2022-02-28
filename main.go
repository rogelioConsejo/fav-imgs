package main

import (
	. "fav-imgs/gallery"
	"fmt"
)

func main() {
	imageGallery := GetGallery(GetPersistence())
	fmt.Printf("%+v", imageGallery.ImageList())
}
