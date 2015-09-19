package core

import (
	"fmt"
	"net/http"
)

type request struct {
	id, name, url, method string
	client, secureClient *http.Client
}

func (this request) run() error {
	request, _ := http.NewRequest(this.method, this.url, nil)
	resp, err := this.client.Do(request)
	if err != nil {
		return err
	}
	
	display := "OK"
	if resp.StatusCode != 200 {
		display = "KO"
	}
	
	fmt.Printf("%s %d %s\r\n", display, resp.StatusCode, this.url)
	
	return nil
}

func (this request) GoString() string {
	return this.url
}