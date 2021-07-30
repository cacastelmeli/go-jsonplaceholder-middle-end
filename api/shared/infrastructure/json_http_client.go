package infrastructure

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type JsonHttpClient struct {
	client  *http.Client
	baseUrl string
}

func NewJsonHttpClient(baseUrl string) *JsonHttpClient {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &JsonHttpClient{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (j *JsonHttpClient) GetJson(path string, result interface{}) error {
	r, err := j.client.Get(j.baseUrl + "/" + path)

	if err != nil {
		return err
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&result)
}

func (j *JsonHttpClient) PostJson(
	path string,
	body map[string]interface{},
	result interface{},
) error {
	jsonBody, _ := json.Marshal(body)
	r, err := j.client.Post(
		j.baseUrl+"/"+path,
		"application/json",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return err
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&result)
}
