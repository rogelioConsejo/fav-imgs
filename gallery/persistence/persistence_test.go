package persistence

import (
	"fav-imgs/gallery"
	"fav-imgs/gallery/image"
	"fmt"
	"testing"
)

func TestPersistence_GetImages(t *testing.T) {
	imageGallery := gallery.GetReader(GetPersistenceReader())
	fmt.Printf("%+v", imageGallery.ImageList())
}

func TestPersistence_AddImage(t *testing.T) {
	imageWriter := gallery.GetImageAdder(GetPersistenceAdder())
	testImage := image.NewImage("test title", "https://picsum.photos/")
	testImageId := imageWriter.Add(testImage)
	imageGallery := gallery.GetReader(GetPersistenceReader())
	fmt.Printf("%+v", imageGallery.ImageList())
	if _, exists := imageGallery.ImageList()[testImageId]; !exists {
		t.Error("image not retrieved correctly")
	}
	fmt.Printf("%+v", imageGallery.ImageList()[testImageId].GetTitle())
	fmt.Printf("%+v", imageGallery.ImageList()[testImageId].GetUrl())

	imageDeleter := gallery.GetImageDeleter(GetPersistenceDeleter())
	imageDeleter.Delete(testImageId)

	imageReader := gallery.GetReader(GetPersistenceReader())
	if len(imageReader.ImageList()) != 0 {
		t.Error("images not deleted correctly")
	}
}
