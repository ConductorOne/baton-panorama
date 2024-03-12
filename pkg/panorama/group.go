package panorama

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type (
	Group struct {
		Name    string   `xml:"name,attr"`
		Members []string `xml:"user>member"`
	}
	ListGroupsResult struct {
		Groups []Group `xml:"user-group>entry"`
	}
	ListGroupsResponse struct {
		PanoramaResponseBase
		Result ListGroupsResult `xml:"result"`
	}
)

func (c *Client) ListGroups(ctx context.Context) ([]Group, *http.Response, error) {
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
	xpath := "/config/shared/local-user-database/user-group"

	query := u.Query()
	query.Set("type", requestType)
	query.Set("action", action)
	query.Set("xpath", xpath)
	u.RawQuery = query.Encode()

	req, err := c.NewRequest(ctx, http.MethodPost, u, WithAcceptXMLHeader())
	if err != nil {
		return nil, nil, err
	}

	var response ListGroupsResponse
	resp, err := c.Do(req, WithXMLResponse(&response))
	if err != nil {
		return nil, nil, err
	}

	if response.Status != "success" {
		return nil, resp, fmt.Errorf("failed to list user-groups with error code: %s", response.Code)
	}

	return response.Result.Groups, resp, nil
}
