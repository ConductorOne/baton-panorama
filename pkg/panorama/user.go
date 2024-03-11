package panorama

import (
	"context"
	"net/http"
	"net/url"
)

type User struct {
	Name     string `xml:"name,attr"`
	Disabled bool   `xml:"disabled"`
}

func (c *Client) ListUsers(ctx context.Context) ([]User, *http.Response, error) {
	stringUrl, err := url.JoinPath(c.baseUrl, API_PATH)
	if err != nil {
		return nil, nil, err
	}

	u, err := url.Parse(stringUrl)
	if err != nil {
		return nil, nil, err
	}

	requestType := "config"
	action := "get"
	xpath := "/config/shared/local-user-database/user"

	query := u.Query()
	query.Set("type", requestType)
	query.Set("action", action)
	query.Set("xpath", xpath)
	u.RawQuery = query.Encode()

	req, err := c.NewRequest(ctx, http.MethodPost, u, WithAcceptXMLHeader())
	if err != nil {
		return nil, nil, err
	}

	var users []User
	resp, err := c.Do(req, WithXMLResponse(&users))
	if err != nil {
		return nil, nil, err
	}

	return users, resp, nil
}
