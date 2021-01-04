package migmig

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
)

var (
	// errors
	errMethodMissing = errors.New("Please provide method")
	errInvalidMethod = errors.New("Provided method is invalid")

	// supported HTTP methods
	methods = map[string]struct{}{
		http.MethodGet:     {},
		http.MethodPost:    {},
		http.MethodPut:     {},
		http.MethodDelete:  {},
		http.MethodPatch:   {},
		http.MethodHead:    {},
		http.MethodOptions: {},
	}
)

// MigMig is an HTTP Request builder and sender.
// MigMig struct will hold all default configs.
type MigMig struct {
	Defaults Config
	Client   *http.Client
}

// Config struct represent a config
type Config struct {
	BaseURL     string
	URL         string
	Method      string
	TargetType  string
	Headers     map[string]string
	QueryParams map[string]string
	Body        map[string]interface{}
}

// New method creates a new MigMig client.
func New() *MigMig {
	return &MigMig{
		Defaults: Config{},
		Client:   &http.Client{},
	}
}

// Create method creates a new MigMig client with default configs.
func Create(c Config) *MigMig {
	return &MigMig{
		Defaults: c,
		Client:   &http.Client{},
	}
}

// getMethod returns request method
func (m *MigMig) getMethod(method string) (string, error) {
	// if method is not provided in request level, set method to the default one.
	if method == "" {
		method = m.Defaults.Method
	}

	// if no method is not provided at all either in request level or default config return an error
	if method == "" {
		return "", errMethodMissing
	}

	// method must be valid.
	if _, ok := methods[method]; !ok {
		return "", errInvalidMethod
	}

	return method, nil
}

// getRequestURL returns request url
func (m *MigMig) getRequestURL(requestURL string, baseURL string) (string, error) {
	// if requestURL is not provided in request level, set requestURL to the default one.
	if requestURL == "" {
		requestURL = m.Defaults.URL
	}

	// if baseURL is not provided in request level, set baseURL to the default one.
	if baseURL == "" {
		baseURL = m.Defaults.BaseURL
	}

	parsedURL, err := url.Parse(requestURL)
	if err != nil {
		return "", err
	}

	// if provided requestURL is absolute we ignore the baseURL.
	if parsedURL.IsAbs() {
		return parsedURL.String(), nil
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// if provided requestURL is not absolute prepend the baseURL to the requestURL.
	parsedBaseURL.Path = path.Join(parsedBaseURL.Path, requestURL)

	return parsedBaseURL.String(), nil
}

func doesMethodAcceptBody(method string) bool {
	return !(method == http.MethodHead || method == http.MethodOptions || method == http.MethodGet)
}

// getRequestBody returns request body
func (m *MigMig) getRequestBody(body map[string]interface{}, method string) (*bytes.Reader, error) {
	// if body is not provided in request level, set body to the default one.
	// the provided body in request level will overright the default one.
	if len(body) == 0 {
		body = m.Defaults.Body
	}

	// if no body provided at all or request method does not suppoort body return nil
	if len(body) == 0 || !doesMethodAcceptBody(method) {
		return nil, nil
	}

	content, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(content), nil
}

// setQueryParams attach provided query parameters to request
// this will merge provided query parameters with default query parameters.
func (m *MigMig) setQueryParams(req *http.Request, queryParams map[string]string) {
	qp := make(map[string]string)

	for k, v := range m.Defaults.QueryParams {
		qp[k] = v
	}

	for k, v := range queryParams {
		qp[k] = v
	}

	if qp != nil {
		q := req.URL.Query()
		for param, value := range qp {
			q.Add(param, value)
		}
		req.URL.RawQuery = q.Encode()
	}
}

// setHeaders attach provided headers to request
// this will merge provided header with default headers.
func (m *MigMig) setHeaders(req *http.Request, headers map[string]string) {
	h := make(map[string]string)

	for k, v := range m.Defaults.Headers {
		h[k] = v
	}

	for k, v := range headers {
		h[k] = v
	}

	if h != nil {
		for param, value := range h {
			req.Header.Set(param, value)
		}
	}

	// default content type is application/json
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
}

// Request returns creates new http.Request and send the request.
func (m *MigMig) Request(c *Config) (*http.Response, error) {
	method, err := m.getMethod(c.Method)
	if err != nil {
		return nil, err
	}

	requestURL, err := m.getRequestURL(c.URL, c.BaseURL)
	if err != nil {
		return nil, err
	}

	contentReader, err := m.getRequestBody(c.Body, method)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	if contentReader != nil {
		req, err = http.NewRequest(method, requestURL, contentReader)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, requestURL, nil)
		if err != nil {
			return nil, err
		}
	}

	m.setQueryParams(req, c.QueryParams)
	m.setHeaders(req, c.Headers)

	resp, err := m.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get method does GET HTTP request.
func (m *MigMig) Get(url string, c *Config) (*http.Response, error) {
	if c == nil {
		c = &Config{}
	}

	c.URL = url
	c.Method = http.MethodGet
	return m.Request(c)
}

// Post method does Post HTTP request.
func (m *MigMig) Post(url string, c *Config) (*http.Response, error) {
	if c == nil {
		c = &Config{}
	}

	c.URL = url
	c.Method = http.MethodPost
	return m.Request(c)
}

// Delete method does Delete HTTP request.
func (m *MigMig) Delete(url string, c *Config) (*http.Response, error) {
	if c == nil {
		c = &Config{}
	}

	c.URL = url
	c.Method = http.MethodDelete
	return m.Request(c)
}

// Patch method does Patch HTTP request.
func (m *MigMig) Patch(url string, c *Config) (*http.Response, error) {
	if c == nil {
		c = &Config{}
	}
	c.URL = url
	c.Method = http.MethodPatch
	return m.Request(c)
}

// Put method does Put HTTP request.
func (m *MigMig) Put(url string, c *Config) (*http.Response, error) {
	if c == nil {
		c = &Config{}
	}
	c.URL = url
	c.Method = http.MethodPut
	return m.Request(c)
}

// Options method does Options HTTP request.
func (m *MigMig) Options(url string, c *Config) (*http.Response, error) {
	if c == nil {
		c = &Config{}
	}
	c.URL = url
	c.Method = http.MethodOptions
	return m.Request(c)
}

// Head method does Head HTTP request.
func (m *MigMig) Head(url string, c *Config) (*http.Response, error) {
	if c == nil {
		c = &Config{}
	}
	c.URL = url
	c.Method = http.MethodHead
	return m.Request(c)
}
