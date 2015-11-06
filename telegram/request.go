package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// HTTPClient Small interface for http
type HTTPClient interface {
	Get(url string) (*http.Response, error)
	Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error)
}

// Request basic api requester
type Request interface {
	Call(method string, in interface{}) (*Response, error)
}

type requestImpl struct {
	client  HTTPClient
	baseURL string
}

// NewRequest Create a new request
func NewRequest(client HTTPClient, baseURL string) Request {
	return &requestImpl{client, baseURL}
}

// Call sends a message to telegram.
func (req *requestImpl) Call(method string, in interface{}) (*Response, error) {
	url := fmt.Sprintf("%s%s", req.baseURL, method)
	var response *http.Response
	var err error

	if in == nil {
		if response, err = req.client.Get(url); err != nil {
			return nil, err
		}
	} else {
		outData, err := json.Marshal(in)
		if err != nil {
			return nil, err
		}
		log.Debug("Request: " + string(outData))

		postData := bytes.NewBuffer(outData)
		mime := "application/json"
		if response, err = req.client.Post(url, mime, postData); err != nil {
			return nil, err
		}
	}

	if response == nil {
		log.Warning("Res: %v Err: %v Url: %v Data: %v", response, err, url, in)

		return nil, err
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	log.Debug("Response: " + string(bytes))

	var out Response
	err = json.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
