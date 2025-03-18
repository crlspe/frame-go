package tlsclient

import (
	"strings"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

type TlsHttpClient struct {
	tlsClient tls_client.HttpClient
}

type RequestMethod string

func NewClient(browserProfile profiles.ClientProfile, timeout int) (TlsHttpClient, error) { //(tls_client.HttpClient, error) {
	var client, err = tls_client.NewHttpClient(
		tls_client.NewNoopLogger(),
		tls_client.WithClientProfile(browserProfile),
		tls_client.WithTimeoutSeconds(timeout),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(tls_client.NewCookieJar()),
	)

	return TlsHttpClient{
		tlsClient: client,
	}, err
}

func (tc *TlsHttpClient) Post(url string, jsonStringData string) (*http.Response, error) {
	var data = strings.NewReader(jsonStringData)
	var request, err = http.NewRequest(http.MethodPost, url, data)
	tc.setHeaders(request)
	if err != nil {
		return nil, err
	}
	return tc.tlsClient.Do(request)
}

func (tc *TlsHttpClient) setHeaders(request *http.Request) {
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Encoding", "Identity")
}
