package mangaeden

import (
	"encoding/json"
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
