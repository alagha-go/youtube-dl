package youtube

import "net/http"


type Header struct {
	Key			string
	Value		string
}

func SetHeaders(req *http.Request, client *Client) {
	for _, header := range client.Headers {
		req.Header.Set(header.Key, header.Value)
	}
}