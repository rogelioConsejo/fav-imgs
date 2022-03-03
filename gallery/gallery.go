package gallery

import . "fav-imgs/gallery/interfaces"

type galleryReader struct {
	persistence Read
}

func (g galleryReader) ImageList() map[string]Image {
	return g.persistence.GetImages()
}

func GetReader(persistence Read) GalleryReader {
	return galleryReader{persistence: persistence}
}

type galleryImageAdder struct {
	persistence Add
}

func (g galleryImageAdder) Add(image Image) (id string) {
	return g.persistence.AddImage(image)
}

func GetImageAdder(persistence Add) GalleryImageAdder {
	return &galleryImageAdder{persistence: persistence}
}

type galleryImageDeleter struct {
	persistence Delete
}

func (g galleryImageDeleter) Delete(id string) {
	g.persistence.DeleteImage(id)
}

func GetImageDeleter(persistence Delete) GalleryImageDeleter {
	return &galleryImageDeleter{persistence: persistence}
}
