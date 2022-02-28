package gallery

import (
	"fav-imgs/gallery/image"
	"fav-imgs/gallery/interfaces"
	"fmt"
	"testing"
)

func TestGalleryViewer_ImageList(t *testing.T) {
	var fakePersistence = new(stubPersistence)
	var testViewer interfaces.GalleryReader = GetReader(fakePersistence)
	fmt.Printf("%+v\n", testViewer.ImageList())
	fmt.Printf("%+v\n", testViewer.ImageList()[0])
	fmt.Printf("%+v\n", testViewer.ImageList()[0].GetTitle())
	fmt.Printf("%+v\n", testViewer.ImageList()[0].GetUrl())
}

func TestGalleryImageAdder_Add(t *testing.T) {
	var fakePersistence = new(mockPersistence)
	var testImageAdder interfaces.GalleryImageAdder = GetImageAdder(fakePersistence)
	testImage := image.NewImage("test image TITLE", "https://picsum.photos/300/300")
	fmt.Printf("%+v\n", testImageAdder.Add(testImage))
}

type stubPersistence struct {
}

func (s stubPersistence) GetImages() []interfaces.Image {
	return []interfaces.Image{image.NewImage("test tile", "localhost")}
}

type mockPersistence struct {
	images []interfaces.Image
}

func (m mockPersistence) AddImage(image interfaces.Image) (id uint) {
	return 0
}
