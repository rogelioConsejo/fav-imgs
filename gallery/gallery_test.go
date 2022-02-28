package gallery

import (
	"fmt"
	"testing"
)

func TestGalleryInterface(t *testing.T) {
	var testGallery = newTestGallery()
	var mockImageList = testGallery.ImageList()
	fmt.Printf("%+v\n", mockImageList[1].GetTitle())
	fmt.Printf("%+v\n", mockImageList[1].GetImageUrl())
}

type testGallery struct {
}

func (t testGallery) ImageList() []Image {
	var images []Image = newMockImages()
	return images
}

type testImage struct {
}

func (i testImage) GetImageUrl() string {
	return "localhost/testImage"
}

func (i testImage) GetTitle() string {
	return "test image title"
}

func newMockImages() []Image {
	return []Image{testImage{}, testImage{}}
}

func newTestGallery() Gallery {
	return new(testGallery)
}
