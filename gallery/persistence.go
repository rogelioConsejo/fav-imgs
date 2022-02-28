package gallery

import . "fav-imgs/gallery/interfaces"
import "fav-imgs/gallery/image"

type persistence struct {
}

func (p persistence) GetImages() []Image {
	return []Image{image.NewImage("1", "https://picsum.photos/200/300"), image.NewImage("2", "https://picsum.photos/300/300")}
}

func GetPersistence() Persistence {
	return persistence{}
}
