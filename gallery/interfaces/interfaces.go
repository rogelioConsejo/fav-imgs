package interfaces

type GalleryReader interface {
	ImageList() []Image
}

type GalleryImageAdder interface {
	Add(image Image) (id uint)
}

type Image interface {
	GetTitle() string
	GetUrl() string
}

type Read interface {
	GetImages() []Image
}

type Add interface {
	AddImage(image Image) (id uint)
}

type Persistence interface {
	Read
}
