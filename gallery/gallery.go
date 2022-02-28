package gallery

type Gallery interface {
	ImageList() []Image
}

type Image interface {
	GetTitle() string
	GetImageUrl() string
}
