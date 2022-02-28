package gallery

import (
	"fmt"
	"testing"
)

func TestPersistence_GetImages(t *testing.T) {
	imageGallery := GetReader(GetPersistence())
	fmt.Printf("%+v", imageGallery.ImageList())
}
