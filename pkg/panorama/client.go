package panorama

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
)

const (
	RequstType    = "config"
	RequestAction = "get"
	SuccessStatus = "success"
)

type (
	Client struct {
		uhttp.BaseHttpClient

		apiUrl *url.URL
	}
	PanoramaResponseBase struct {
		XMLName xml.Name `xml:"response"`
		Status  string   `xml:"status,attr"`
		Code    string   `xml:"code,attr"`
	}
)

func New(baseUrl string, httpClient *http.Client) (*Client, error) {
	stringUrl, err := url.JoinPath(baseUrl, "/api")
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(stringUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseHttpClient: *uhttp.NewBaseHttpClient(httpClient),
		apiUrl:         u,
	}, nil
}

func (c *Client) GetUrl() *url.URL {
	return &url.URL{
		Scheme:     c.apiUrl.Scheme,
		Opaque:     c.apiUrl.Opaque,
		User:       c.apiUrl.User,
		Host:       c.apiUrl.Host,
		Path:       c.apiUrl.Path,
		RawPath:    c.apiUrl.RawPath,
		ForceQuery: c.apiUrl.ForceQuery,
		RawQuery:   c.apiUrl.RawQuery,
		Fragment:   c.apiUrl.Fragment,
	}
}
