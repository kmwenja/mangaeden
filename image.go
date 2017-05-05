package mangaeden

import (
	"strings"
)

const DEFAULT_IMAGE_URL = "https://cdn.mangaeden.com/mangasimg/"
const EMPTY_IMAGE_URL = "https://cdn.mangaeden.com/images/logo2.png"

type Image struct {
	ImageURLFragment string
}

func (i *Image) ImageUrl() string {
	if i.ImageURLFragment == "" {
		return EMPTY_IMAGE_URL
	}
	return DEFAULT_IMAGE_URL + i.ImageURLFragment
}

func (i *Image) Ext() string {
	u := i.ImageUrl()
	return u[strings.LastIndex(u, "."):]
}
