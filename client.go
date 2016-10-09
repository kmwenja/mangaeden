package mangaeden

const (
	DEFAULT_URL      = "https://www.mangaeden.com/"
	DEFAULT_API_URL  = DEFAULT_URL + "api/"
	LANG_ENG         = 0
	LANG_ITA         = 1
	STATUS_COMPLETED = 2
	STATUS_ONGOING   = 1
)

type Client struct{}

func New() *Client {
	c := Client{}
	return &c
}
