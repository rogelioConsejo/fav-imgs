package interfaces

type Read interface {
	GetImages() map[string]Image
}

type Add interface {
	AddImage(image Image) (id string)
}

type Delete interface {
	DeleteImage(id string)
}

type Persistence interface {
	Read
	Add
	Delete
}
