package mangaeden

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ChapterImage struct {
	Image
	Index  int
	Height int
	Width  int
}

func (ci *ChapterImage) UnmarshalJSON(data []byte) error {
	var aux []interface{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ci.Index = int(aux[0].(float64))
	ci.ImageURLFragment = aux[1].(string)
	ci.Height = int(aux[2].(float64))
	ci.Width = int(aux[3].(float64))
	return nil
}

type ChapterResult struct {
	Images []ChapterImage `json:"images"`
}

func (c *Client) GetChapterImages(id string) ([]ChapterImage, error) {
	url := fmt.Sprintf(DEFAULT_API_URL+"chapter/%s/", id)
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ChapterResult
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Images, nil
}
