package mangaeden

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	DEFAULT_URL      = "https://www.mangaeden.com/"
	DEFAULT_API_URL  = DEFAULT_URL + "api/"
	LANG_ENG         = 0
	LANG_ITA         = 1
	STATUS_COMPLETED = 2
	STATUS_ONGOING   = 1
)

func completeString(status int) string {
	switch status {
	case STATUS_COMPLETED:
		return "Completed"
	case STATUS_ONGOING:
		return "Ongoing"
	default:
		return "Unknown"
	}
}

type Client struct {
	client *http.Client
}

func New(hc *http.Client) *Client {
	c := Client{client: hc}
	if hc == nil {
		// TODO add timeout
		c.client = &http.Client{}
	}
	return &c
}

func (c *Client) list(url string) (MangaPage, error) {
	resp, err := c.client.Get(url)
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

func (c *Client) Manga(id string) (MangaInfo, error) {
	url := fmt.Sprintf(DEFAULT_API_URL+"manga/%s/", id)
	resp, err := c.client.Get(url)
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

func (c *Client) Chapter(id string) ([]ChapterImage, error) {
	url := fmt.Sprintf(DEFAULT_API_URL+"chapter/%s/", id)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Images []ChapterImage `json:"images"`
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Images, nil
}

func (c *Client) DownloadImage(i Image) (io.ReadCloser, error) {
	u := i.ImageUrl()
	resp, err := c.client.Get(u)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
