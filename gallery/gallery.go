package gallery

import "fav-imgs/gallery/interfaces"

type gallery struct {
	persistence Persistence
}

func (g gallery) ImageList() []interfaces.Image {
	return g.persistence.GetImages()
}

type Persistence interface {
	GetImages() []interfaces.Image
}

func GetGallery(persistence Persistence) interfaces.Gallery {
	return gallery{persistence: persistence}
}
