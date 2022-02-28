package gallery

import . "fav-imgs/gallery/interfaces"

type galleryReader struct {
	persistence Read
}

func (g galleryReader) ImageList() []Image {
	return g.persistence.GetImages()
}

func GetReader(persistence Read) GalleryReader {
	return galleryReader{persistence: persistence}
}

type galleryImageAdder struct {
	persistence Add
}

func (g galleryImageAdder) Add(image Image) (id uint) {
	return g.persistence.AddImage(image)
}

func GetImageAdder(persistence Add) GalleryImageAdder {
	return &galleryImageAdder{persistence: persistence}
}
