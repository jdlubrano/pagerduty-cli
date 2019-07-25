package escalation_policies

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

  if path != "/escalation_policies" {
    t.Errorf("Unexpected path - expected: /escalation_policies, got %s", path)
  }

  if queryString != "query=test+query" {
    t.Errorf("Unexpected query - expected: query=test+query, got: %s", queryString)
  }
}

func TestListForQueryEscalationPolicies(t *testing.T) {
  const testResponseContent = `
  {
    "escalation_policies": [
      {
        "id": "EP1",
        "name": "Escalation policy 1"
      }
    ]
  }
  `

  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, testResponseContent)
  }))

  defer testServer.Close()

  apiClient := &api_client.ApiClient{ApiToken: apiToken, BaseUrl: testServer.URL}
  escalationPolicies, _ := ListForQuery(apiClient, "")
  escalationPolicy := escalationPolicies[0]

  if escalationPolicy.Id != "EP1" {
    t.Errorf("Unexpected id - expected: EP1, got: %s", escalationPolicy.Id)
  }

  if escalationPolicy.Name != "Escalation policy 1" {
    t.Errorf("Unexpected id - expected: Escalation policy 1, got: %s", escalationPolicy.Name)
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
