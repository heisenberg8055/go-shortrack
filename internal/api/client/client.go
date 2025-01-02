package client

import (
	"net/http"
	"time"
)

func newClient() *http.Client {
	c := http.Client{
		Timeout: time.Second * 3,
	}
	return &c
}

func RequestURL(url string) {
	c := newClient()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	_, err = c.Do(req)
	if err != nil {
		panic(err)
	}

}
