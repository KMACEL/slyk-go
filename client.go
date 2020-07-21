package slyk

type Client struct {
	apiKey string
}

func New(setApiKey string) *Client {
	return &Client{apiKey: setApiKey}
}
