package gallery

import (
	"fav-imgs/gallery/image"
	. "fav-imgs/gallery/interfaces"
	"fav-imgs/gallery/persistence"
	"fmt"
	"testing"
)

func TestGalleryViewer_ImageList(t *testing.T) {
	var fakePersistence = new(stubPersistence)
	var testViewer GalleryReader = GetReader(fakePersistence)
	fmt.Printf("%+v\n", testViewer.ImageList())
	fmt.Printf("%+v\n", testViewer.ImageList()[stubImageId])
	fmt.Printf("%+v\n", testViewer.ImageList()[stubImageId].GetTitle())
	fmt.Printf("%+v\n", testViewer.ImageList()[stubImageId].GetUrl())
}

func TestGalleryImageAdder_Add_Delete(t *testing.T) {
	var fakePersistence = MakeMockPersistence()
	var testImageAdder GalleryImageAdder = GetImageAdder(fakePersistence)
	testImage1 := image.NewImage("test image title 1", "https://picsum.photos/300/300")
	testImage2 := image.NewImage("test image title 2", "https://picsum.photos/200/300")
	imageId1 := testImageAdder.Add(testImage1)
	imageId2 := testImageAdder.Add(testImage2)
	testReader := GetReader(fakePersistence)

	retrievedImage := testReader.ImageList()[imageId1]
	fmt.Printf("%+v\n", retrievedImage.GetTitle())
	fmt.Printf("%+v\n", retrievedImage.GetUrl())

	retrievedImage2 := testReader.ImageList()[imageId2]
	fmt.Printf("%+v\n", retrievedImage2.GetTitle())
	fmt.Printf("%+v\n", retrievedImage2.GetUrl())

	var testImageDeleter GalleryImageDeleter = GetImageDeleter(fakePersistence)
	testImageDeleter.Delete(imageId1)
	testImageDeleter.Delete(imageId2)
	if len(testReader.ImageList()) != 0 {
		t.Error("images not deleted")
	}
}

type stubPersistence struct {
}

const stubImageId = "a"

func (s stubPersistence) GetImages() map[string]Image {
	images := make(map[string]Image)
	images[stubImageId] = image.NewImage("test tile", "localhost")
	return images
}

type mockPersistence struct {
	images map[string]Image
}

func (m *mockPersistence) DeleteImage(id string) {
	delete(m.images, id)
}

func MakeMockPersistence() Persistence {
	mock := mockPersistence{}
	mock.images = make(map[string]Image)
	return &mock
}

func (m mockPersistence) GetImages() map[string]Image {
	return m.images
}

func (m *mockPersistence) AddImage(image Image) (id string) {
	id = persistence.RandStringRunes(10)
	m.images[id] = image
	return id
}
