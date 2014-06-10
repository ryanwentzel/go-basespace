package basespace

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	sdkVersion = "0.1"
	userAgent  = "go-basespace/" + sdkVersion

	defaultApiURL = "https://api.basespace.illumina.com/"
	apiVersion    = "v1pre3"
)

// A Client handles communication with the BaseSpace API
type Client struct {
	// HTTP client used to communicate with the BaseSpace API
	client *http.Client

	// BaseSpace API URL. ApiURL should always contain a trailing slash.
	ApiURL *url.URL

	// Version of the BaseSpace API to communicate with.
	ApiVersion string

	// User agent used when making requests to the BaseSpace API
	UserAgent string

	Users       *UsersService
	AppSessions *AppSessionsService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	apiURL, _ := url.Parse(defaultApiURL)

	newClient := &Client{client: httpClient, ApiURL: apiURL, ApiVersion: apiVersion, UserAgent: userAgent}
	newClient.Users = &UsersService{client: newClient, mapper: &Mapper{}}
	newClient.AppSessions = &AppSessionsService{client: newClient, mapper: &Mapper{}}

	return newClient
}

func (c *Client) NewRequest(method string, uri string, body interface{}) (*http.Request, error) {
	uri = c.ApiVersion + "/" + uri
	rel, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	u := c.ApiURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request) (*ApiResponse, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	apiResponse, err := c.getResponse(resp)
	if err != nil {
		return nil, err
	}

	return apiResponse, nil
}

func (c *Client) getResponse(r *http.Response) (*ApiResponse, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse ApiResponse
	e := json.Unmarshal(data, &apiResponse)
	if e != nil {
		return nil, e
	}

	return &apiResponse, nil
}
