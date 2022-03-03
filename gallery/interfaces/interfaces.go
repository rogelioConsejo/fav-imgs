package interfaces

type GalleryReader interface {
	ImageList() map[string]Image
}

type GalleryImageAdder interface {
	Add(image Image) (id string)
}

type GalleryImageDeleter interface {
	Delete(id string)
}

type Image interface {
	GetTitle() string
	GetUrl() string
}
