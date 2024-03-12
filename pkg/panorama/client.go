package panorama

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
)

const (
	ApiPath       = "/api"
	RequstType    = "config"
	RequestAction = "get"
	SuccessStatus = "success"
)

type (
	Client struct {
		uhttp.BaseHttpClient

		baseUrl string
	}
	PanoramaResponseBase struct {
		XMLName xml.Name `xml:"response"`
		Status  string   `xml:"status,attr"`
		Code    string   `xml:"code,attr"`
	}
)

func New(baseUrl string, httpClient *http.Client) (*Client, error) {
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
