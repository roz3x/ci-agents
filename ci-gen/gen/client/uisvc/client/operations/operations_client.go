package operations

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/websocket"

	transport "github.com/erikh/go-transport"
	"github.com/tinyci/ci-agents/clients/jsonbuffer"
	"github.com/tinyci/ci-agents/errors"

	models "github.com/tinyci/ci-agents/ci-gen/gen/client/uisvc/models"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Client is the base client struct.
type Client struct {
	token  string
	client *http.Client
	url    *url.URL
}

// New creates a new *Client. Passing a cert will enable client/server
// certificate authentication; otherwise pass nil for no auth.
func New(baseURL string, token string, cert *transport.Cert) (*Client, *errors.Error) {
	t, err := transport.NewHTTP(cert)
	if err != nil {
		return nil, errors.New(err)
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.New(err)
	}

	client := t.Client(&http.Transport{IdleConnTimeout: 15 * time.Second, MaxIdleConns: 10})
	return &Client{url: u, client: client, token: token}, nil
}

// SetJar overwrites the cookie jar for the underlying HTTP client.
func (c *Client) SetJar(jar *cookiejar.Jar) {
	c.client.Jar = jar
}

// MarshalCookies will marshal any existing cookies to the byte array.
func (c *Client) MarshalCookies() ([]byte, error) {
	return json.Marshal(c.client.Jar.Cookies(c.url))
}

// UnmarshalCookies loads cookies into the given client.
func (c *Client) UnmarshalCookies(content []byte) error {
	cookies := []*http.Cookie{}
	if err := json.Unmarshal(content, &cookies); err != nil {
		return err
	}

	c.client.Jar.SetCookies(c.url, cookies)
	return nil
}

// DeleteCapabilitiesUsernameCapability remove a named capability
func (c *Client) DeleteCapabilitiesUsernameCapability(ctx context.Context, capability string, username string) *errors.Error {
	route := "/capabilities/{username}/{capability}"
	route = strings.Replace(route, "{capability}", url.PathEscape(fmt.Sprintf("%v", capability)), -1)

	route = strings.Replace(route, "{username}", url.PathEscape(fmt.Sprintf("%v", username)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("DELETE", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// DeleteToken remove and reset your tiny c i access token
func (c *Client) DeleteToken(ctx context.Context) *errors.Error {
	route := "/token"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("DELETE", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetErrors retrieve errors
func (c *Client) GetErrors(ctx context.Context) ([]*models.UserError, *errors.Error) {
	route := "/errors"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return make([]*models.UserError, 0, 50), errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return make([]*models.UserError, 0, 50), errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return make([]*models.UserError, 0, 50), errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return make([]*models.UserError, 0, 50), origErr
	}

	defer resp.Body.Close()

	var result []*models.UserError

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return make([]*models.UserError, 0, 50), errors.New(err)
	}

	return result, nil

}

// GetLogAttachID attach to a running log
func (c *Client) GetLogAttachID(ctx context.Context, id int64, w io.WriteCloser) *errors.Error {
	route := "/log/attach/{id}"
	route = strings.Replace(route, "{id}", url.PathEscape(fmt.Sprintf("%v", id)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	if u.Scheme == "https" {
		u.Scheme = "wss"
	} else {
		u.Scheme = "ws"
	}

	conn, err := websocket.Dial(u.String(), "", "http://localhost")
	if err != nil {
		return errors.New(err)
	}
	if _, err := io.Copy(w, jsonbuffer.NewWrapper(conn)); err != nil {
		return errors.New(err)
	}

	if err := conn.Close(); err != nil {
		return errors.New(err)
	}

	return nil

}

// GetLoggedin check logged in state
func (c *Client) GetLoggedin(ctx context.Context) (string, *errors.Error) {
	route := "/loggedin"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return "", errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return "", errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return "", errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return "", origErr
	}

	defer resp.Body.Close()

	var result string

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", errors.New(err)
	}

	return result, nil

}

// GetLogin log into the system
func (c *Client) GetLogin(ctx context.Context, code string, state string) *errors.Error {
	route := "/login"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["code"] = code

	m["state"] = state

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetLoginUpgrade log into the system with upgraded permissions
func (c *Client) GetLoginUpgrade(ctx context.Context) *errors.Error {
	route := "/login/upgrade"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetLogout log out of the system
func (c *Client) GetLogout(ctx context.Context) *errors.Error {
	route := "/logout"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetRepositoriesCiAddOwnerRepo add a specific repository to c i
func (c *Client) GetRepositoriesCiAddOwnerRepo(ctx context.Context, owner string, repo string) *errors.Error {
	route := "/repositories/ci/add/{owner}/{repo}"
	route = strings.Replace(route, "{owner}", url.PathEscape(fmt.Sprintf("%v", owner)), -1)

	route = strings.Replace(route, "{repo}", url.PathEscape(fmt.Sprintf("%v", repo)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetRepositoriesCiDelOwnerRepo removes a specific repository from c i
func (c *Client) GetRepositoriesCiDelOwnerRepo(ctx context.Context, owner string, repo string) *errors.Error {
	route := "/repositories/ci/del/{owner}/{repo}"
	route = strings.Replace(route, "{owner}", url.PathEscape(fmt.Sprintf("%v", owner)), -1)

	route = strings.Replace(route, "{repo}", url.PathEscape(fmt.Sprintf("%v", repo)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetRepositoriesMy fetch all the writable repositories for the user
func (c *Client) GetRepositoriesMy(ctx context.Context, search string) (models.RepositoryList, *errors.Error) {
	route := "/repositories/my"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["search"] = search

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return make(models.RepositoryList, 0, 50), errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return make(models.RepositoryList, 0, 50), errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return make(models.RepositoryList, 0, 50), errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return make(models.RepositoryList, 0, 50), origErr
	}

	defer resp.Body.Close()

	var result models.RepositoryList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return make(models.RepositoryList, 0, 50), errors.New(err)
	}

	return result, nil

}

// GetRepositoriesScan scan repositories from the remote resource
func (c *Client) GetRepositoriesScan(ctx context.Context) *errors.Error {
	route := "/repositories/scan"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetRepositoriesSubAddOwnerRepo subscribe to a repository running c i
func (c *Client) GetRepositoriesSubAddOwnerRepo(ctx context.Context, owner string, repo string) *errors.Error {
	route := "/repositories/sub/add/{owner}/{repo}"
	route = strings.Replace(route, "{owner}", url.PathEscape(fmt.Sprintf("%v", owner)), -1)

	route = strings.Replace(route, "{repo}", url.PathEscape(fmt.Sprintf("%v", repo)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetRepositoriesSubDelOwnerRepo unsubscribe from a repository
func (c *Client) GetRepositoriesSubDelOwnerRepo(ctx context.Context, owner string, repo string) *errors.Error {
	route := "/repositories/sub/del/{owner}/{repo}"
	route = strings.Replace(route, "{owner}", url.PathEscape(fmt.Sprintf("%v", owner)), -1)

	route = strings.Replace(route, "{repo}", url.PathEscape(fmt.Sprintf("%v", repo)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetRepositoriesSubscribed list all subscribed repositories
func (c *Client) GetRepositoriesSubscribed(ctx context.Context, search string) (models.RepositoryList, *errors.Error) {
	route := "/repositories/subscribed"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["search"] = search

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return make(models.RepositoryList, 0, 50), errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return make(models.RepositoryList, 0, 50), errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return make(models.RepositoryList, 0, 50), errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return make(models.RepositoryList, 0, 50), origErr
	}

	defer resp.Body.Close()

	var result models.RepositoryList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return make(models.RepositoryList, 0, 50), errors.New(err)
	}

	return result, nil

}

// GetRepositoriesVisible fetch all the repositories the user can view
func (c *Client) GetRepositoriesVisible(ctx context.Context, search string) (models.RepositoryList, *errors.Error) {
	route := "/repositories/visible"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["search"] = search

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return make(models.RepositoryList, 0, 50), errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return make(models.RepositoryList, 0, 50), errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return make(models.RepositoryList, 0, 50), errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return make(models.RepositoryList, 0, 50), origErr
	}

	defer resp.Body.Close()

	var result models.RepositoryList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return make(models.RepositoryList, 0, 50), errors.New(err)
	}

	return result, nil

}

// GetRunRunID get a run by ID
func (c *Client) GetRunRunID(ctx context.Context, runID int64) (*models.Run, *errors.Error) {
	route := "/run/{run_id}"
	route = strings.Replace(route, "{run_id}", url.PathEscape(fmt.Sprintf("%v", runID)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return nil, errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return nil, origErr
	}

	defer resp.Body.Close()

	var result models.Run

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.New(err)
	}

	return &result, nil

}

// GetRuns obtain the run list for the user
func (c *Client) GetRuns(ctx context.Context, page int64, perPage int64, repository string, sha string) (models.RunList, *errors.Error) {
	route := "/runs"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["page"] = page

	m["perPage"] = perPage

	m["repository"] = repository

	m["sha"] = sha

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return make(models.RunList, 0, 50), errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return make(models.RunList, 0, 50), errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return make(models.RunList, 0, 50), errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return make(models.RunList, 0, 50), origErr
	}

	defer resp.Body.Close()

	var result models.RunList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return make(models.RunList, 0, 50), errors.New(err)
	}

	return result, nil

}

// GetRunsCount count the runs
func (c *Client) GetRunsCount(ctx context.Context, repository string, sha string) (int64, *errors.Error) {
	route := "/runs/count"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["repository"] = repository

	m["sha"] = sha

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return 0, errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return 0, errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return 0, errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return 0, origErr
	}

	defer resp.Body.Close()

	var result int64

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, errors.New(err)
	}

	return result, nil

}

// GetSubmit perform a manual submission to tiny c i
func (c *Client) GetSubmit(ctx context.Context, all bool, repository string, sha string) *errors.Error {
	route := "/submit"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["all"] = all

	m["repository"] = repository

	m["sha"] = sha

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// GetTasks obtain the task list optionally filtering by repository and sha
func (c *Client) GetTasks(ctx context.Context, page int64, perPage int64, repository string, sha string) (models.TaskList, *errors.Error) {
	route := "/tasks"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["page"] = page

	m["perPage"] = perPage

	m["repository"] = repository

	m["sha"] = sha

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return make(models.TaskList, 0, 50), errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return make(models.TaskList, 0, 50), errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return make(models.TaskList, 0, 50), errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return make(models.TaskList, 0, 50), origErr
	}

	defer resp.Body.Close()

	var result models.TaskList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return make(models.TaskList, 0, 50), errors.New(err)
	}

	return result, nil

}

// GetTasksCount count the tasks
func (c *Client) GetTasksCount(ctx context.Context, repository string, sha string) (int64, *errors.Error) {
	route := "/tasks/count"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["repository"] = repository

	m["sha"] = sha

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return 0, errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return 0, errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return 0, errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return 0, origErr
	}

	defer resp.Body.Close()

	var result int64

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, errors.New(err)
	}

	return result, nil

}

// GetTasksRunsID obtain the run list based on the task ID
func (c *Client) GetTasksRunsID(ctx context.Context, id int64, page int64, perPage int64) (models.RunList, *errors.Error) {
	route := "/tasks/runs/{id}"
	route = strings.Replace(route, "{id}", url.PathEscape(fmt.Sprintf("%v", id)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	m["page"] = page

	m["perPage"] = perPage

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return make(models.RunList, 0, 50), errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return make(models.RunList, 0, 50), errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return make(models.RunList, 0, 50), errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return make(models.RunList, 0, 50), origErr
	}

	defer resp.Body.Close()

	var result models.RunList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return make(models.RunList, 0, 50), errors.New(err)
	}

	return result, nil

}

// GetTasksRunsIDCount count the runs corresponding to the task ID
func (c *Client) GetTasksRunsIDCount(ctx context.Context, id int64) (int64, *errors.Error) {
	route := "/tasks/runs/{id}/count"
	route = strings.Replace(route, "{id}", url.PathEscape(fmt.Sprintf("%v", id)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return 0, errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return 0, errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return 0, errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return 0, origErr
	}

	defer resp.Body.Close()

	var result int64

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, errors.New(err)
	}

	return result, nil

}

// GetTasksSubscribed obtain the list of tasks that belong to repositories you are subscribed to
func (c *Client) GetTasksSubscribed(ctx context.Context, page int64, perPage int64) (models.TaskList, *errors.Error) {
	route := "/tasks/subscribed"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}
	m["page"] = page

	m["perPage"] = perPage

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return make(models.TaskList, 0, 50), errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return make(models.TaskList, 0, 50), errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return make(models.TaskList, 0, 50), errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return make(models.TaskList, 0, 50), origErr
	}

	defer resp.Body.Close()

	var result models.TaskList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return make(models.TaskList, 0, 50), errors.New(err)
	}

	return result, nil

}

// GetToken get a tiny c i access token
func (c *Client) GetToken(ctx context.Context) (string, *errors.Error) {
	route := "/token"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return "", errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return "", errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return "", errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return "", origErr
	}

	defer resp.Body.Close()

	var result string

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", errors.New(err)
	}

	return result, nil

}

// GetUserProperties get information about the current user
func (c *Client) GetUserProperties(ctx context.Context) (interface{}, *errors.Error) {
	route := "/user/properties"

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return nil, errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return nil, origErr
	}

	defer resp.Body.Close()

	var result interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.New(err)
	}

	return result, nil

}

// PostCancelRunID cancel by run ID
func (c *Client) PostCancelRunID(ctx context.Context, runID int64) *errors.Error {
	route := "/cancel/{run_id}"
	route = strings.Replace(route, "{run_id}", url.PathEscape(fmt.Sprintf("%v", runID)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	postForm := url.Values{}

	body = []byte(postForm.Encode())

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// PostCapabilitiesUsernameCapability add a named capability
func (c *Client) PostCapabilitiesUsernameCapability(ctx context.Context, capability string, username string) *errors.Error {
	route := "/capabilities/{username}/{capability}"
	route = strings.Replace(route, "{capability}", url.PathEscape(fmt.Sprintf("%v", capability)), -1)

	route = strings.Replace(route, "{username}", url.PathEscape(fmt.Sprintf("%v", username)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	postForm := url.Values{}

	body = []byte(postForm.Encode())

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}

// PostTasksCancelID cancel by task ID
func (c *Client) PostTasksCancelID(ctx context.Context, id int64) *errors.Error {
	route := "/tasks/cancel/{id}"
	route = strings.Replace(route, "{id}", url.PathEscape(fmt.Sprintf("%v", id)), -1)

	tmp := *c.url
	u := &tmp
	u.Path += route

	m := map[string]interface{}{}

	if len(m) > 0 {
		q := u.Query()

		for key, value := range m {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		u.RawQuery = q.Encode()
	}

	var body []byte

	postForm := url.Values{}

	body = []byte(postForm.Encode())

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Add("Authorization", c.token)

	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.New(err)
	}

	if resp.StatusCode == 500 {
		origErr := &errors.Error{}
		if err := json.NewDecoder(resp.Body).Decode(origErr); err != nil {
			return errors.New(err)
		}
		if origErr == nil {
			panic("Cannot return 500 without error")
		}

		return origErr
	}

	defer resp.Body.Close()

	return nil

}
