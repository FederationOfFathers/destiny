package destiny

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var Debug = false

type Client struct {
	*http.Client
	key string
}

// Get issues a GET to the specified URL.
func (c *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Head issues a HEAD to the specified URL.
func (c *Client) Head(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Post issues a POST to the specified URL.
func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	return c.Do(req)
}

// PostForm issues a POST to the specified URL,
// with data's keys and values URL-encoded as the request body.
func (c *Client) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

// Do sends an HTTP request and returns an HTTP response, following
// policy (such as redirects, cookies, auth) as configured on the
// client.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-Key", c.key)
	req.Header.Set("User-Agent", "github.com/FederationOfFathers/lib-destiny")
	rsp, err := c.Client.Do(req)
	if Debug {
		buf, _ := httputil.DumpRequest(req, true)
		os.Stderr.Write(buf)
		buf2, _ := httputil.DumpResponse(rsp, true)
		os.Stderr.Write(buf2)
	}
	return rsp, err
}

// New returns a new Client which is an http.Client which embeds the
// API Key and User Agent into all its requests automatically
func New(APIKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		Client: httpClient,
		key:    APIKey,
	}
}

func (c *Client) getAndUnwrap(url string, into interface{}) (bool, error) {
	rsp, err := c.Get(url)
	defer rsp.Body.Close()

	if err != nil {
		return false, err
	}
	var e *envelope
	err = json.NewDecoder(rsp.Body).Decode(&e)
	if err != nil {
		return false, err
	}
	if !e.success() {
		return false, e
	}
	return true, e.into(&into)
}

func (c *Client) getAndUnwrapData(url string, into interface{}) (bool, error) {
	rsp, err := c.Get(url)
	defer rsp.Body.Close()
	if err != nil {
		return false, err
	}
	var e *envelope
	err = json.NewDecoder(rsp.Body).Decode(&e)
	if err != nil {
		return false, err
	}
	data, err := e.data()
	if err != nil {
		return false, err
	}
	return true, data.into(&into)
}
