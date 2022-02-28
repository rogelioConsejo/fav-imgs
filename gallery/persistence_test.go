package gallery

import (
	"fmt"
	"testing"
)

func TestPersistence_GetImages(t *testing.T) {
	imageGallery := GetGallery(GetPersistence())
	fmt.Printf("%+v", imageGallery.ImageList())
}
