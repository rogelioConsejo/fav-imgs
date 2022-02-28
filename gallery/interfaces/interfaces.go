package interfaces

type Gallery interface {
	ImageList() []Image
}

type Image interface {
	GetTitle() string
	GetUrl() string
}
