package mangaeden

import (
	"encoding/json"
	"strings"
	"time"
)

type Manga struct {
	Image
	Alias           string    `json:"a"`
	CategoriesList  []string  `json:"c"`
	Hits            int       `json:"h"`
	ID              string    `json:"i"`
	LastChapterDate time.Time `json:"ld"`
	StatusCode      int       `json:"s"`
	Title           string    `json:"t"`
}

func (m *Manga) UnmarshalJSON(data []byte) error {
	type Alias Manga
	aux := &struct {
		ImageURLFragment string  `json:"im"`
		LastChapterDate  float64 `json:"ld"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	m.ImageURLFragment = aux.ImageURLFragment
	m.LastChapterDate = time.Unix(int64(aux.LastChapterDate), 0)
	return nil
}

func (m Manga) IsCompleted() bool {
	return m.StatusCode == STATUS_COMPLETED
}

func (m Manga) Status() string {
	return completeString(m.StatusCode)
}

func (m Manga) Categories() string {
	return strings.Join(m.CategoriesList, ",")
}
