package gallery

import . "fav-imgs/gallery/interfaces"

type gallery struct {
	persistence Persistence
}

func (g gallery) ImageList() []Image {
	return g.persistence.GetImages()
}

func GetGallery(persistence Persistence) Gallery {
	return gallery{persistence: persistence}
}
