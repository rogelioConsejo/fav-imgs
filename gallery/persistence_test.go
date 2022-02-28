package gallery

import (
	"fmt"
	"testing"
)

func TestPersistence_GetImages(t *testing.T) {
	imageGallery := GetReader(GetPersistenceReader())
	fmt.Printf("%+v", imageGallery.ImageList())
}
