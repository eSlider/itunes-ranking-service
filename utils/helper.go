package utils

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ReadFile and return content
func ReadFile(path string) (*[]byte, error) {
	r, err := ioutil.ReadFile(path)
	return &r, err
}

// LoadOverHttp and return content
func LoadOverHttp(url string) (*[]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Status code error: %d %s", r.StatusCode, r.Status))
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = r.Body.Close()
	if err != nil {
		return nil, err
	}

	return &body, err
}
