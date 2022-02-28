package gallery

import (
	"fav-imgs/gallery/image"
	. "fav-imgs/gallery/interfaces"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGalleryViewer_ImageList(t *testing.T) {
	var fakePersistence = new(stubPersistence)
	var testViewer GalleryReader = GetReader(fakePersistence)
	fmt.Printf("%+v\n", testViewer.ImageList())
	fmt.Printf("%+v\n", testViewer.ImageList()[stubImageId])
	fmt.Printf("%+v\n", testViewer.ImageList()[stubImageId].GetTitle())
	fmt.Printf("%+v\n", testViewer.ImageList()[stubImageId].GetUrl())
}

func TestGalleryImageAdder_Add(t *testing.T) {
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

func MakeMockPersistence() Persistence {
	mock := mockPersistence{}
	mock.images = make(map[string]Image)
	return &mock
}

func (m mockPersistence) GetImages() map[string]Image {
	return m.images
}

func (m *mockPersistence) AddImage(image Image) (id string) {
	id = RandStringRunes(10)
	m.images[id] = image
	return id
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890#$%&-_*çÇ¡¿?!")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
