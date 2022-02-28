package interfaces

type GalleryReader interface {
	ImageList() map[string]Image
}

type GalleryImageAdder interface {
	Add(image Image) (id string)
}

type Image interface {
	GetTitle() string
	GetUrl() string
}

type Read interface {
	GetImages() map[string]Image
}

type Add interface {
	AddImage(image Image) (id string)
}

type Persistence interface {
	Read
	Add
}
