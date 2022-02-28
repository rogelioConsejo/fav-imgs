package gallery

import (
	"fmt"
	"testing"
)

func TestGallery(t *testing.T) {
	var testGallery = newTestGallery()
	fmt.Printf("%+v\n", testGallery.ImageList())
}

type testGallery struct {
}

func (t testGallery) ImageList() []Image {
	var images []Image = newMockImages()
	return images
}

func newMockImages() []Image {
	return []Image{}
}

func newTestGallery() Gallery {
	return new(testGallery)
}
