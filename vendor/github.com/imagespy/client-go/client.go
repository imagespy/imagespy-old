package imagespy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/die-net/lrucache"
	"github.com/gregjones/httpcache"
)

const (
	// DefaultAPIEndpoint is the default endpoint of the ImageSpy API.
	DefaultAPIEndpoint = "https://imagespy.hydrantosaurus.com"
	version            = "0.1.0"
	userAgent          = "imagespy-go/" + version
)

type requester struct {
	baseURL *url.URL
	client  *http.Client
}

func (r *requester) readAsJSON(path string) (*http.Response, error) {
	url, err := r.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	return r.send(req)
}

func (r *requester) writeAsJSON(path string, v interface{}) (*http.Response, error) {
	url, err := r.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	return r.send(req)
}

func (r *requester) parseJSON(re io.ReadCloser, v interface{}) error {
	defer re.Close()
	b, err := ioutil.ReadAll(re)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return nil
}

func (r *requester) send(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", userAgent)
	Log.Debug(fmt.Sprintf("Sending \"%s\" request to \"%s\"", req.Method, req.URL.String()))
	resp, err := r.client.Do(req)
	if resp.Header.Get(httpcache.XFromCache) == "1" {
		Log.Debug(fmt.Sprintf("\"%s\" served from cache", req.URL.String()))
	}
	if resp != nil {
		Log.Debug(fmt.Sprintf("Received status code \"%d\" for request to \"%s\"", resp.StatusCode, req.URL.String()))
	}

	return resp, err
}

// ClientV1 is the client for the Image Spy API version 1.
type ClientV1 struct {
	ImageSpy *ImageSpyService
	r        *requester
}

// WithBaseURL sets the base URL of the ImageSpy API.
func (c *ClientV1) WithBaseURL(baseURL string) *ClientV1 {
	u, _ := url.Parse(baseURL)
	c.r.baseURL = u
	return c
}

// WithHTTPClient sets the http.Client.
func (c *ClientV1) WithHTTPClient(hc *http.Client) *ClientV1 {
	c.r.client = hc
	return c
}

// WithTimeout sets the timeout of the underlying http.Client.
func (c *ClientV1) WithTimeout(s string) *ClientV1 {
	d, _ := time.ParseDuration(s)
	c.r.client.Timeout = d
	return c
}

// NewClientV1 returns a new client for the V1 HTTP API.
func NewClientV1() *ClientV1 {
	httpClient := &http.Client{}
	requester := &requester{}
	c := &ClientV1{
		ImageSpy: &ImageSpyService{requester: requester},
		r:        requester,
	}
	cache := lrucache.New(12582912, 0)
	t := httpcache.NewTransport(cache)
	t.Transport = httpClient.Transport
	httpClient.Transport = t
	c.WithHTTPClient(httpClient).WithBaseURL(DefaultAPIEndpoint).WithTimeout("5s")
	return c
}
