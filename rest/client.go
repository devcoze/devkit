package rest

import "net/url"

type Client struct {
	// base is the root URL for all invocations of the client
	base *url.URL
	// group stand for the client group, eg: iam.api, iam.authz
	group string
	// versionedAPIPath is a path segment connecting the base URL to the resource root
	versionedAPIPath string
	// content describes how a RESTClient encodes and decodes responses.
}
