package api_client

import (
  "fmt"
  "testing"
  "net/http"
  "net/http/httptest"
)

const apiToken = "test-token"
const testResponseContent = "Test response content"

func TestApiClientResponseStatus(t *testing.T) {
  const status = 204

  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(status)
  }))

  defer testServer.Close()

  apiClient := &ApiClient{apiToken: apiToken, baseUrl: testServer.URL}
  resp, _ := apiClient.Get("/test", nil)

  if resp.Status != status {
    t.Errorf("Unexpected Response.Status - expected: %d, got: %d", status, resp.Status)
  }
}

func TestApiClientResponseBody(t *testing.T) {
  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, testResponseContent)
  }))

  defer testServer.Close()

  apiClient := &ApiClient{apiToken: apiToken, baseUrl: testServer.URL}
  resp, _ := apiClient.Get("/test", nil)

  if string(resp.Body) != testResponseContent {
    t.Errorf(`
      Unexpected Response body
      expected: %s
      got: %s`,
      resp.Body,
      testResponseContent,
    )
  }
}

func TestApiClientGet(t *testing.T) {
  var requestMethod string

  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    requestMethod = r.Method
  }))

  defer testServer.Close()

  apiClient := &ApiClient{apiToken: apiToken, baseUrl: testServer.URL}
  apiClient.Get("", nil)

  if requestMethod != "GET" {
    t.Errorf("Unexpected request method - expected: GET, got: %s", requestMethod)
  }
}
