package mangaeden

import (
	"strings"
)

const DEFAULT_IMAGE_URL = "https://cdn.mangaeden.com/mangasimg/"

type Image struct {
	ImageURLFragment string
}

func (i *Image) ImageUrl() string {
	return DEFAULT_IMAGE_URL + i.ImageURLFragment
}

func (i *Image) Ext() string {
	u := i.ImageURLFragment
	return u[strings.LastIndex(u, "."):]
}
