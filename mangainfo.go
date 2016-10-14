package mangaeden

import (
	"encoding/json"
	"time"
)

type MangaInfo struct {
	Image
	Artist          string    `json:"artist"`
	Author          string    `json:"author"`
	Categories      []string  `json:"categories"`
	Chapters        []Chapter `json:"chapters"`
	ChaptersLen     int       `json:"chapters_len"`
	Created         time.Time `json:"created"`
	Description     string    `json:"description"`
	Hits            int       `json:"hits"`
	Language        int       `json:"language"`
	LastChapterDate time.Time `json:"last_chapter_image"`
	Status          int       `json:"status"`
	Title           string    `json:"title"`
}

func (mi *MangaInfo) UnmarshalJSON(data []byte) error {
	type Alias MangaInfo
	aux := &struct {
		Created          float64 `json:"created"`
		ImageURLFragment string  `json:"image"`
		LastChapterDate  float64 `json:"last_chapter_image"`
		*Alias
	}{
		Alias: (*Alias)(mi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	mi.Created = time.Unix(int64(aux.Created), 0)
	mi.ImageURLFragment = aux.ImageURLFragment
	mi.LastChapterDate = time.Unix(int64(aux.LastChapterDate), 0)
	return nil
}
