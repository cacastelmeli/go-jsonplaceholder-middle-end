package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const JSON_PLACEHOLDER_BASE_URL = "https://jsonplaceholder.typicode.com"

type JsonPlaceholderClient struct {
	client  *http.Client
	baseUrl string
}

type JsonPlaceholderPost struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func NewJsonPlaceholderClient() *JsonPlaceholderClient {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &JsonPlaceholderClient{
		client:  client,
		baseUrl: JSON_PLACEHOLDER_BASE_URL,
	}
}

func (j *JsonPlaceholderClient) GetJson(path string, result interface{}) error {
	r, err := j.client.Get(j.baseUrl + "/" + path)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&result)
}

func (j *JsonPlaceholderClient) PostJson(
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

func (j *JsonPlaceholderClient) GetAllPosts() ([]*JsonPlaceholderPost, error) {
	posts := []*JsonPlaceholderPost{}
	err := j.GetJson("posts", &posts)

	return posts, err
}
