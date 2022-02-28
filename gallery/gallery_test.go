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
	fmt.Printf("%+v", testGallery.ImageList())
}
