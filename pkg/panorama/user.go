package panorama

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type (
	User struct {
		Name     string `xml:"name,attr"`
		Disabled bool   `xml:"disabled"`
	}
	UserRaw struct {
		Name     string `xml:"name,attr"`
		Disabled string `xml:"disabled"`
	}
	ListUsersResult struct {
		Users []UserRaw `xml:"user>entry"`
	}
	ListUsersResponse struct {
		PanoramaResponseBase
		Result ListUsersResult `xml:"result"`
	}
)

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

	var response ListUsersResponse
	resp, err := c.Do(req, WithXMLResponse(&response))
	if err != nil {
		return nil, nil, err
	}

	if response.Status != "success" {
		return nil, nil, fmt.Errorf("failed to list users with error code: %s", response.Code)
	}

	var users []User
	for _, user := range response.Result.Users {
		users = append(users, user.mapToUser())
	}

	return users, resp, nil
}

func (u *UserRaw) mapToUser() User {
	disabled := false
	if u.Disabled == "yes" {
		disabled = true
	}
	return User{
		Name:     u.Name,
		Disabled: disabled,
	}
}
