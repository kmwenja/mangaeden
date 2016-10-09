package mangaeden

const DEFAULT_IMAGE_URL = "https://cdn.mangaeden.com/mangasimg/"

type Image struct {
	ImageURLFragment string
}

func (i *Image) ImageUrl() string {
	return DEFAULT_IMAGE_URL + i.ImageURLFragment
}
