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

type GalleryImageModifier interface {
	Update(id string, image Image)
}

type Image interface {
	GetTitle() string
	GetUrl() string
}
