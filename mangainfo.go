package mangaeden

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (c *Client) GetInfo(id string) (MangaInfo, error) {
	url := fmt.Sprintf(DEFAULT_API_URL+"manga/%s/", id)
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return MangaInfo{}, err
	}
	defer resp.Body.Close()
	var info MangaInfo
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&info)
	if err != nil {
		return MangaInfo{}, err
	}
	return info, nil
}
