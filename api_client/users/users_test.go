package users

import(
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/jdlubrano/pagerduty-cli/api_client"
)

const apiToken = "test-token"

func TestGetMeApiCall(t *testing.T) {
  var path string

  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    path = r.URL.Path
  }))

  defer testServer.Close()

  apiClient := &api_client.ApiClient{ApiToken: apiToken, BaseUrl: testServer.URL}
  GetMe(apiClient)

  if path != "/users/me" {
    t.Errorf("Unexpected path - expected: /escalation_policies, got %s", path)
  }
}

func TestGetMeUser(t *testing.T) {
  const testResponseContent = `
  {
    "user": {
      "id": "TestUserId",
      "name": "Test User",
      "teams": [
        {
          "id": "TestTeamId",
          "summary": "Test Team"
        }
      ]
    }
  }
  `

  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, testResponseContent)
  }))

  apiClient := &api_client.ApiClient{ApiToken: apiToken, BaseUrl: testServer.URL}
  user, _ := GetMe(apiClient)

  if user.Id != "TestUserId" {
    t.Errorf("Unexpected user ID - expected: TestUserId, got: %s", user.Id)
  }

  if user.Name != "Test User" {
    t.Errorf("Unexpected user Name - expected: Test User, got: %s", user.Name)
  }

  team := user.Teams[0]

  if team.Id != "TestTeamId" {
    t.Errorf("Unexpected team ID - expected: TestTeamId, got: %s", team.Id)
  }

  if team.Name != "Test Team" {
    t.Errorf("Unexpected team Name - expected: Test Team, got: %s", team.Name)
  }
}

func TestGetMeError(t *testing.T) {
  testServer := httptest.NewServer(nil)
  testServer.Close()

  apiClient := &api_client.ApiClient{ApiToken: apiToken, BaseUrl: testServer.URL}
  _, err := GetMe(apiClient)

  if err == nil {
    t.Errorf("Unexpected response - expected: an error, got: nil")
  }
}
