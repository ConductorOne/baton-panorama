package panorama

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type (
	User struct {
		Name     string
		Disabled bool
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
	GetUserResult struct {
		User UserRaw `xml:"entry"`
	}
	GetUserResponse struct {
		PanoramaResponseBase
		Result GetUserResult `xml:"result"`
	}
)

func (c *Client) ListUsers(ctx context.Context) ([]User, *http.Response, error) {
	stringUrl, err := url.JoinPath(c.baseUrl, ApiPath)
	if err != nil {
		return nil, nil, err
	}

	u, err := url.Parse(stringUrl)
	if err != nil {
		return nil, nil, err
	}

	xpath := "/config/shared/local-user-database/user"

	query := u.Query()
	query.Set("type", RequstType)
	query.Set("action", RequestAction)
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

	if response.Status != SuccessStatus {
		return nil, nil, fmt.Errorf("failed to list users with error code: %s", response.Code)
	}

	var users []User
	for _, user := range response.Result.Users {
		users = append(users, user.mapToUser())
	}

	return users, resp, nil
}

func (c *Client) GetUser(ctx context.Context, name string) (*User, *http.Response, error) {
	stringUrl, err := url.JoinPath(c.baseUrl, ApiPath)
	if err != nil {
		return nil, nil, err
	}

	u, err := url.Parse(stringUrl)
	if err != nil {
		return nil, nil, err
	}

	xpath := fmt.Sprintf("/config/shared/local-user-database/user/entry[@name='%s']", name)

	query := u.Query()
	query.Set("type", RequstType)
	query.Set("action", RequestAction)
	query.Set("xpath", xpath)
	u.RawQuery = query.Encode()

	req, err := c.NewRequest(ctx, http.MethodPost, u, WithAcceptXMLHeader())
	if err != nil {
		return nil, nil, err
	}

	var response GetUserResponse
	resp, err := c.Do(req, WithXMLResponse(&response))
	if err != nil {
		return nil, nil, err
	}

	if response.Status != SuccessStatus {
		return nil, nil, fmt.Errorf("failed to get user with error code: %s", response.Code)
	}

	user := response.Result.User.mapToUser()

	return &user, resp, nil
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
