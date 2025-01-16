package edgex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	client              *http.Client
	CoreMetadataService *CoreMetadataService
}

type service struct {
	client     *Client
	BaseURL    *url.URL
	ApiVersion string
}

func NewClient() *Client {
	c := &Client{}
	c.client = &http.Client{}
	c.initialize()

	return c
}

func (c *Client) initialize() {
	host, ok := os.LookupEnv("EDGEX_HOST")
	if !ok {
		host = "http://localhost"
	}

	apiVersion, ok := os.LookupEnv("EDGEX_API_VERSION")
	if !ok {
		apiVersion = "v2"
	}

	coreMetadataPort, ok := os.LookupEnv("EDGEX_CORE_METADATA_PORT")
	if !ok {
		coreMetadataPort = "59881"
	}

	coreMetadataUrl, _ := url.Parse(fmt.Sprintf("%s:%s/api/%s", host, coreMetadataPort, apiVersion))
	c.CoreMetadataService = &CoreMetadataService{client: c, BaseURL: coreMetadataUrl, ApiVersion: apiVersion}
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if body != nil {
		buf = &bytes.Buffer{}
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// dump the whole request pretty
	// dump, err := httputil.DumpRequestOut(req, true)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("request: %s\n", string(dump))

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// dump the whole response pretty
	// dump, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("response: %s\n", string(dump))

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil && err != io.EOF {
			return nil, err
		}
	}

	return resp, nil
}
