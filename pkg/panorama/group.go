package panorama

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
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
	GetGroupResponse struct {
		PanoramaResponseBase
		Result GetGroupResult `xml:"result"`
	}
	GetGroupResult struct {
		Group Group `xml:"entry"`
	}
)

func (c *Client) ListGroups(ctx context.Context) ([]Group, *http.Response, error) {
	stringUrl, err := url.JoinPath(c.baseUrl, ApiPath)
	if err != nil {
		return nil, nil, err
	}

	u, err := url.Parse(stringUrl)
	if err != nil {
		return nil, nil, err
	}

	xpath := "/config/shared/local-user-database/user-group"

	query := u.Query()
	query.Set("type", RequstType)
	query.Set("action", RequestAction)
	query.Set("xpath", xpath)
	u.RawQuery = query.Encode()

	req, err := c.NewRequest(ctx, http.MethodPost, u, uhttp.WithAcceptXMLHeader())
	if err != nil {
		return nil, nil, err
	}

	var response ListGroupsResponse
	resp, err := c.Do(req, uhttp.WithXMLResponse(&response))
	if err != nil {
		return nil, nil, err
	}

	if response.Status != SuccessStatus {
		return nil, resp, fmt.Errorf("failed to list user-groups with error code: %s", response.Code)
	}

	return response.Result.Groups, resp, nil
}

func (c *Client) GetGroup(ctx context.Context, name string) (*Group, *http.Response, error) {
	stringUrl, err := url.JoinPath(c.baseUrl, ApiPath)
	if err != nil {
		return nil, nil, err
	}

	u, err := url.Parse(stringUrl)
	if err != nil {
		return nil, nil, err
	}

	xpath := fmt.Sprintf("/config/shared/local-user-database/user-group/entry[@name='%s']", name)

	query := u.Query()
	query.Set("type", RequstType)
	query.Set("action", RequestAction)
	query.Set("xpath", xpath)
	u.RawQuery = query.Encode()

	req, err := c.NewRequest(ctx, http.MethodPost, u, uhttp.WithAcceptXMLHeader())
	if err != nil {
		return nil, nil, err
	}

	var response GetGroupResponse
	resp, err := c.Do(req, uhttp.WithXMLResponse(&response))
	if err != nil {
		return nil, nil, err
	}

	if response.Status != SuccessStatus {
		return nil, resp, fmt.Errorf("failed to get user-group with error code: %s", response.Code)
	}

	return &response.Result.Group, resp, nil
}
