package mangaeden

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) list(url string) (MangaPage, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return MangaPage{}, err
	}
	defer resp.Body.Close()
	var page MangaPage
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&page)
	if err != nil {
		return MangaPage{}, err
	}
	return page, nil
}

func (c *Client) ListAll(lang int) ([]Manga, error) {
	url := fmt.Sprintf(DEFAULT_API_URL+"list/%d/", lang)
	page, err := c.list(url)
	if err != nil {
		return nil, err
	}
	return page.MangaList, nil
}

func (c *Client) ListPage(lang int, page int, pageSize int) (MangaPage, error) {
	if pageSize == 0 {
		pageSize = 500
	}

	if pageSize < 25 || pageSize > 1500 {
		return MangaPage{}, fmt.Errorf("invalid page size: %d. page size should be between 25 and 1500", pageSize)
	}

	url := fmt.Sprintf(DEFAULT_API_URL+"list/%d/?p=%d&l=%d", lang, page, pageSize)
	return c.list(url)
}
