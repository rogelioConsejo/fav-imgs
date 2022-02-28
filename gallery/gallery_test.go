package gallery

import (
	"fav-imgs/gallery/image"
	"fav-imgs/gallery/interfaces"
	"fmt"
	"testing"
)

type stubPersistence struct {
}

func (s stubPersistence) GetImages() []interfaces.Image {
	return []interfaces.Image{image.NewImage("test tile", "localhost")}
}

func TestGalleryInterface(t *testing.T) {
	var fakePersistence = new(stubPersistence)
	var testGallery interfaces.Gallery = GetGallery(fakePersistence)
	fmt.Printf("%+v\n", testGallery.ImageList())
	fmt.Printf("%+v\n", testGallery.ImageList()[0])
	fmt.Printf("%+v\n", testGallery.ImageList()[0].GetTitle())
	fmt.Printf("%+v\n", testGallery.ImageList()[0].GetUrl())
}
