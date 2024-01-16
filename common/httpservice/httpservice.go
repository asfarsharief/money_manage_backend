package httpservice

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"

	log "github.com/asfarsharief/money_management_backend/common/logingservice"
)

type HttpService struct{}

func NewHttpService() *HttpService {
	return &HttpService{}
}

// Post - Post function for a Rest API call
func (h *HttpService) Post(url string, body io.Reader) (*http.Response, error) {
	client := createHTTPClient()

	clientRequest, err := createHTTPRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return client.Do(clientRequest)

}

// func post with headers
func (h *HttpService) PostWithHeaders(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	client := createHTTPClient()

	clientRequest, err := createHTTPRequestWithHeaders(http.MethodPost, url, body, headers)
	if err != nil {
		return nil, err
	}

	return client.Do(clientRequest)

}

// Get - Get function for the Rest API calls
func (h *HttpService) Get(url string) (*http.Response, error) {
	client := createHTTPClient()

	clientRequest, err := createHTTPRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(clientRequest)
}

// GetWithAuth - GetWithAuth function for the Rest API calls with Authorization header
func (h *HttpService) GetWithAuth(url string, authHeader string) (*http.Response, error) {
	client := createHTTPClient()

	clientRequest, err := createHTTPRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	clientRequest.Header.Set("Authorization", authHeader)

	return client.Do(clientRequest)
}

// Delete - Delete function for the Rest API calls
func (h *HttpService) Delete(url string) (*http.Response, error) {
	client := createHTTPClient()

	clientRequest, err := createHTTPRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(clientRequest)
}

// Put - Put function for a Rest API call
func (h *HttpService) Put(url string, body io.Reader) (*http.Response, error) {
	client := createHTTPClient()

	clientRequest, err := createHTTPRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	return client.Do(clientRequest)

}

// createHTTPClient : create http client with timeout and auth details
func createHTTPClient() http.Client {
	timeout := time.Duration(15) * time.Second
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
}

// createHTTPRequest : returns http request object for the Api calls,
// TODO Add methods for customizing header and Auth
func createHTTPRequest(requestType string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(requestType, url, body)
	if err != nil {
		log.Errorf("%s", "error while creating request object")
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func createHTTPRequestWithHeaders(requestType string, url string, body io.Reader, headers map[string]string) (*http.Request, error) {
	request, err := http.NewRequest(requestType, url, body)
	if err != nil {
		log.Errorf("%s", "error while creating request object")
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	for key, val := range headers {
		request.Header.Set(key, val)
	}

	return request, nil
}
