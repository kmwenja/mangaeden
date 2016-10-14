package mangaeden

import (
	"encoding/json"
	"time"
)

type Chapter struct {
	Index int
	Date  time.Time
	Title string
	ID    string
}

func (c *Chapter) UnmarshalJSON(data []byte) error {
	var aux []interface{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.Index = int(aux[0].(float64))
	c.Date = time.Unix(int64(aux[1].(float64)), 0)
	if aux[2] == nil {
		c.Title = ""
	} else {
		c.Title = aux[2].(string)
	}
	c.ID = aux[3].(string)
	return nil
}
