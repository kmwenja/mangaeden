package mangaeden

import (
	"encoding/json"
	"strings"
	"time"
)

type MangaInfo struct {
	Image
	Artist          string    `json:"artist"`
	Author          string    `json:"author"`
	CategoriesList  []string  `json:"categories"`
	Chapters        []Chapter `json:"chapters"`
	ChaptersLen     int       `json:"chapters_len"`
	Created         time.Time `json:"created"`
	RawDescription  string    `json:"description"`
	Hits            int       `json:"hits"`
	LanguageCode    int       `json:"language"`
	LastChapterDate time.Time `json:"last_chapter_image"`
	StatusCode      int       `json:"status"`
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

func (mi MangaInfo) IsCompleted() bool {
	return mi.StatusCode == STATUS_COMPLETED
}

func (mi MangaInfo) Status() string {
	return completeString(mi.StatusCode)
}

func (mi MangaInfo) Language() string {
	switch mi.LanguageCode {
	case LANG_ENG:
		return "English"
	case LANG_ITA:
		return "Italian"
	default:
		return "Unknown Language"
	}
}

func (mi MangaInfo) Categories() string {
	return strings.Join(mi.CategoriesList, ",")
}

func (mi MangaInfo) Description() string {
	return strings.Trim(mi.RawDescription, "\n")
}
