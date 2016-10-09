package mangaeden

type MangaPage struct {
	Page      int     `json:"page"`
	Start     int     `json:"start"`
	End       int     `json:"end"`
	Total     int     `json:"total"`
	MangaList []Manga `json:"manga"`
}

func (mp *MangaPage) Pages() int {
	n := len(mp.MangaList)
	if n == 0 {
		return -1
	}
	return mp.Total / n
}
