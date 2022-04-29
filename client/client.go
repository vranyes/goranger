package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/vranyes/goranger/utils"
)

type RangerClient struct {
	api_base   string
	username   string
	password   string
	httpClient *http.Client
}

func NewClient(ranger_host, username, password string, verify_ssl bool) RangerClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: verify_ssl},
	}

	return RangerClient{
		api_base:   ranger_host + "/service/public/v2/api",
		username:   username,
		password:   password,
		httpClient: &http.Client{Transport: tr},
	}
}

func (r RangerClient) RequestHandler(verb string, path string, body io.Reader) ([]byte, error) {
	full_path := r.api_base + path
	fmt.Println(verb + " " + full_path)
	request, err := http.NewRequest(verb, full_path, body)
	utils.Check(err)
	request.SetBasicAuth(r.username, r.password)
	request.Header.Set("Content-Type", "application/json")
	resp, err := r.httpClient.Do(request)
	utils.Check(err)
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
