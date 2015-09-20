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
	concise bool
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

func (this postmanRunner) ouputMessage(message string) {
	if !this.concise {
		fmt.Println(message)
	}
}

func (this postmanRunner) Export(source string) bool {
	this.ouputMessage("Not implemented feature.")
	this.source = source
	return true
}

func (this postmanRunner) Run(source string, concise bool) bool {
	
	this.source = source
	this.concise = concise
	
	if err := this.downloadRequestData(); err != nil {
		fmt.Println(err)
		return false
	}
	
	failedCount := 0
	if len(*this.requests) > 0 {
		this.ouputMessage("Running requests.")		
		for _, r := range *this.requests {
			result, err := r.run()
			if err != nil {
				fmt.Println(err)
				return false		
			}
			
			if !result {
				failedCount++
			}
		}
	}
	
	fmt.Printf("%d requests returned non success status.\r\n", failedCount)	
	
	return failedCount == 0
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
	
	this.ouputMessage(fmt.Sprintf("Downloaded %d request definitions.", len(*this.requests)))
	return nil
}