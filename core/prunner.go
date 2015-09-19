package core

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"crypto/tls"
)

type postmanRunner struct {
	source string
	requests *[]request
	client, secureClient *http.Client
	
}

func NewPostmanRunner() *postmanRunner {
	r := new(postmanRunner)
	r.requests = new([]request)
	r.client = &http.Client{}
	
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{},
		DisableCompression: true,
	}
	r.secureClient = &http.Client{Transport: tr}
	
	return r
}

func (this postmanRunner) Run(url string) {
	
	this.source = url
	
	if err := this.downloadRequestData(); err != nil {
		fmt.Println(err)
		return
	}
	
	if len(*this.requests) > 0 {
		fmt.Println("Running requests.")		
		for _, r := range *this.requests {
			if err := r.run(); err != nil {
				fmt.Println(err)
				return				
			}
		}
	}
}

func (this postmanRunner) downloadRequestData() error {

	resp, err := this.secureClient.Get(this.source)	
	if err != nil {
		return err
	}
		
	defer resp.Body.Close()
	
	buffer, _ := ioutil.ReadAll(resp.Body)
	
	_, err = resp.Body.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		fmt.Println(err.Error())
		return err
	}
	
	var i interface{}
	if err = json.Unmarshal(buffer, &i); err != nil {
		return err
	}
	
	decodedData := i.(map[string]interface{})
	for _, element := range decodedData["requests"].([]interface{}) {
		requestData := element.(map[string]interface{})
		
		r := request{
			id: requestData["id"].(string),
			name: requestData["name"].(string), 
			url: requestData["url"].(string),
			method: requestData["method"].(string),
			client: this.client,
			secureClient: this.secureClient}
		
		*this.requests = append(*this.requests, r)
	}
	
	fmt.Printf("Downloaded %d request definitions.\r\n", len(*this.requests))
	
	return nil
}