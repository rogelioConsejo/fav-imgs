package gallery

import . "fav-imgs/gallery/interfaces"
import "fav-imgs/gallery/image"

type persistence struct {
}

// GetImages TODO: Fully Implement
func (p persistence) GetImages() map[string]Image {
	images := make(map[string]Image)
	images["0"] = image.NewImage("1", "https://picsum.photos/200/300")
	images["1"] = image.NewImage("2", "https://picsum.photos/200/250")
	images["2"] = image.NewImage("3", "https://picsum.photos/200/320")

	return images
}

func GetPersistenceReader() Read {
	return persistence{}
}
