package schedule_overrides

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
  "time"

  "github.com/jdlubrano/pagerduty-cli/api_client"
)

const apiToken = "test-token"
var startTime, _ = time.Parse(time.RFC3339, "2019-09-23T12:00:00Z")
var endTime, _ = time.Parse(time.RFC3339, "2019-09-23T13:00:00Z")

type TestOverride struct {
  scheduleId string
  userId string
  startTime time.Time
  endTime time.Time
}

func (o *TestOverride) ScheduleId() string {
  return o.scheduleId
}

func (o *TestOverride) UserId() string {
  return o.userId
}

func (o *TestOverride) StartTime() time.Time {
  return o.startTime
}

func (o *TestOverride) EndTime() time.Time {
  return o.endTime
}

func TestCreateOverrideApiCall(t *testing.T) {
  var path, requestBody string

  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    path = r.URL.Path
    bodyBytes, _ := ioutil.ReadAll(r.Body)
    requestBody = string(bodyBytes)
  }))

  defer testServer.Close()

  apiClient := &api_client.ApiClient{ApiToken: apiToken, BaseUrl: testServer.URL}

  override := &TestOverride{
    "test_schedule_id",
    "test_user_id",
    startTime,
    endTime,
  }

  expectedBody := `{"override":{"start":"2019-09-23T12:00:00Z","end":"2019-09-23T13:00:00Z","user":{"id":"test_user_id","type":"user_reference"}}}`

  CreateOverride(apiClient, override)

  if path != "/schedules/test_schedule_id/overrides" {
    t.Errorf("Unexpected path - expected: /schedules/test_schedule_id/overrides, got %s", path)
  }

  if requestBody != expectedBody {
    t.Errorf("Unexpected path - expected: %s, got %s", expectedBody, requestBody)
  }
}

func TestCreateOverride(t *testing.T) {
  const testResponseContent = `
  {
    "override": {
      "id": "test_override_id",
      "start": "2019-09-23T12:00:00Z",
      "end": "2019-09-23T13:00:00Z",
      "user": {
        "id": "test_user_id"
      }
    }
  }
  `

  testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, testResponseContent)
  }))

  defer testServer.Close()

  apiClient := &api_client.ApiClient{ApiToken: apiToken, BaseUrl: testServer.URL}

  override := &TestOverride{
    "test_schedule_id",
    "test_user_id",
    startTime,
    endTime,
  }

  o, _ := CreateOverride(apiClient, override)

  if o.Id != "test_override_id" {
    t.Errorf("Unexpected override ID - expected: %s, got: %s", "test_schedule_id", o.Id)
  }

  if o.StartTime != startTime {
    t.Errorf("Unexpected start time - expected: %s, got: %s", startTime.Format(time.RFC3339), o.StartTime.Format(time.RFC3339))
  }

  if o.EndTime != endTime {
    t.Errorf("Unexpected start time - expected: %s, got: %s", endTime.Format(time.RFC3339), o.EndTime.Format(time.RFC3339))
  }

  if o.User.Id != "test_user_id" {
    t.Errorf("Unexpected user id - expected: test_user_id, got: %s", o.User.Id)
  }
}
