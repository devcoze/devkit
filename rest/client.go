package rest

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"net/url"
)

type Interface interface {
	Verb(verb string) *resty.Response

	Get() *resty.Response

	Post() *resty.Response

	Put() *resty.Response

	Delete() *resty.Response
}

type Client struct {
	// base is the root URL for all invocations of the clientcmd
	base *url.URL
	// group stand for the clientcmd group, eg: iam.api, iam.authz
	group string
	// versionedAPIPath is a path segment connecting the base URL to the resource root
	versionedAPIPath string

	client *resty.Client
	// content describes how a RESTClient encodes and decodes responses.
}

func NewRestClient() *resty.Client {
	client := resty.New().
		SetHeader("Content-Type", "application/json").
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: false,
			Certificates: []tls.Certificate{tls.Certificate{
				Certificate: [][]byte{},
				PrivateKey:  nil,
				Leaf:        nil,
			}},
		})

	rep, err := client.R().Get("/healthz")
	if err != nil {
		panic(err)
	}
	println("Health check response:", rep.String())
	return client
}

type ClientContentConfig struct {
}

func (c *Client) Verb(verb string, url string) (*resty.Response, error) {
	return c.client.R().Execute(verb, url)
}

func (c *Client) Get(url string) (*resty.Response, error) {
	return c.Verb(resty.MethodGet, url)
}

func (c *Client) Post(url string) (*resty.Response, error) {
	return c.Verb(resty.MethodPost, url)
}

func (c *Client) Put() (*resty.Response, error) {
	return c.Verb(resty.MethodPut, "")
}
