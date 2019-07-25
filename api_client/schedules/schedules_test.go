package schedules

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/jdlubrano/pagerduty-cli/api_client"
)

const apiToken = "test-token"

func TestListForQueryApiCall(t *testing.T) {
  var path, queryString string

  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    path = r.URL.Path
    queryString = r.URL.Query().Encode()
  }))

  defer testServer.Close()

  apiClient := &api_client.ApiClient{ApiToken: apiToken, BaseUrl: testServer.URL}
  ListForQuery(apiClient, "test query")

  if path != "/schedules" {
    t.Errorf("Unexpected path - expected: /escalation_policies, got %s", path)
  }

  if queryString != "query=test+query" {
    t.Errorf("Unexpected query - expected: query=test+query, got: %s", queryString)
  }
}

func TestListForQueryEscalationPolicies(t *testing.T) {
  const testResponseContent = `
  {
    "schedules": [
      {
        "id": "S1",
        "name": "Schedule 1",
        "description": "A test schedule"
      }
    ]
  }
  `

  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, testResponseContent)
  }))

  defer testServer.Close()

  apiClient := &api_client.ApiClient{ApiToken: apiToken, BaseUrl: testServer.URL}
  schedules, _ := ListForQuery(apiClient, "")
  schedule := schedules[0]

  if schedule.Id != "S1" {
    t.Errorf("Unexpected id - expected: S1, got: %s", schedule.Id)
  }

  if schedule.Name != "Schedule 1" {
    t.Errorf("Unexpected name - expected: Schedule 1, got: %s", schedule.Name)
  }

  if schedule.Description != "A test schedule" {
    t.Errorf("Unexpected description - expected: A test schedule, got: %s", schedule.Description)
  }
}

func TestListForQueryError(t *testing.T) {
  testServer := httptest.NewServer(nil)
  testServer.Close()

  apiClient := &api_client.ApiClient{ApiToken: apiToken, BaseUrl: testServer.URL}
  _, err := ListForQuery(apiClient, "")

  if err == nil {
    t.Errorf("Unexpected response - expected: an error, got: nil")
  }
}
