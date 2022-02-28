package gallery

import . "fav-imgs/gallery/interfaces"
import "fav-imgs/gallery/image"

type persistence struct {
}

func (p persistence) GetImages() []Image {
	return []Image{image.NewImage("1", "url1"), image.NewImage("2", "url2")}
}

func GetPersistence() Persistence {
	return persistence{}
}
