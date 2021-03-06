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
  ApiToken string
  BaseUrl string
}

func NewClient() *ApiClient {
  return &ApiClient{ApiToken: config.GetApiToken(), BaseUrl: baseUrl}
}

func (api *ApiClient) Get(path string, queryParams *map[string]string) (*Response, error) {
  queryValues := url.Values{}

  if queryParams != nil {
    for k, v := range *queryParams {
      queryValues.Add(k, v)
    }
  }

  url := api.BaseUrl+path+"?"+queryValues.Encode()
  req, _ := http.NewRequest("GET", url, nil)
  return api.performRequest(req)
}

func (api *ApiClient) Post(path string, params *map[string]interface{}) (*Response, error) {
  body, err := json.Marshal(params)

  if err != nil {
    return nil, err
  }

  req, _ := http.NewRequest("POST", api.BaseUrl+path, bytes.NewBuffer(body))
  return api.performRequest(req)
}

func (api *ApiClient) performRequest(request *http.Request) (*Response, error) {
  client := &http.Client{}
  api.addHeaders(request)
  resp, err := client.Do(request)

  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

  return buildResponse(resp)
}

func (api *ApiClient) addHeaders(request *http.Request) {
  request.Header.Add("Authorization", "Token token="+api.ApiToken)
  request.Header.Add("Accept", "application/vnd.pagerduty+json;version=2")
  request.Header.Add("Content-Type", "application/json")
}

func buildResponse(resp *http.Response) (*Response, error) {
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    return nil, err
  }

  return &Response{Status: resp.StatusCode, Body: body}, nil
}
