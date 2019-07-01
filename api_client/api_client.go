package api_client

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"

  "github.com/jdlubrano/pagerduty-cli/config"
)

type Response struct {
  Status int
  Body []byte
}

func (response *Response) ParseInto(dataStruct interface{}) error {
  return json.Unmarshal(response.Body, dataStruct)
}

const baseUrl = "https://api.pagerduty.com"

type ApiClient struct {
  baseUrl string
}

func NewClient() *ApiClient {
  return &ApiClient{baseUrl: baseUrl}
}

func (api *ApiClient) Get(path string, queryParams *map[string]string) (*Response, error) {
  queryValues := url.Values{}

  if queryParams != nil {
    for k, v := range *queryParams {
      queryValues.Add(k, v)
    }
  }

  req, _ := http.NewRequest("GET", baseUrl+path+queryValues.Encode(), nil)
  return performRequest(req)
}

func (api *ApiClient) Post(path string, params *map[string]interface{}) (*Response, error) {
  body, err := json.Marshal(params)

  if err != nil {
    return nil, err
  }

  req, _ := http.NewRequest("POST", baseUrl+path, bytes.NewBuffer(body))
  return performRequest(req)
}

func performRequest(request *http.Request) (*Response, error) {
  client := &http.Client{}
  req := addHeaders(request)
  resp, err := client.Do(req)
  defer resp.Body.Close()

  if err != nil {
    return nil, err
  }

  return buildResponse(resp)
}

func addHeaders(request *http.Request) *http.Request {
  token := config.GetApiToken()
  request.Header.Add("Authorization", "Token token="+token)
  request.Header.Add("Accept", "application/vnd.pagerduty+json;version=2")
  request.Header.Add("Content-Type", "application/json")
  return request
}

func buildResponse(resp *http.Response) (*Response, error) {
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    return nil, err
  }

  return &Response{Status: resp.StatusCode, Body: body}, nil
}
