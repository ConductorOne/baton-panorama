package panorama

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
)

const API_PATH = "/api"

type Client struct {
	uhttp.BaseHttpClient

	baseUrl string
}

func New(baseUrl, username, password string, httpClient *http.Client) (*Client, error) {
	return &Client{
		BaseHttpClient: *uhttp.NewBaseHttpClient(httpClient),
		baseUrl:        baseUrl,
	}, nil
}

// TODO: move following to SDK

func WithAcceptXMLHeader() uhttp.RequestOption {
	return func() (io.ReadWriter, map[string]string, error) {
		return nil, map[string]string{
			"Accept": "application/xml",
		}, nil
	}
}

func WithXMLResponse(response interface{}) uhttp.DoOption {
	return func(resp *uhttp.WrapperResponse) error {
		return xml.Unmarshal(resp.Body, response)
	}
}
