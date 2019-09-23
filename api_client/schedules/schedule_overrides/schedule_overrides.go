package schedule_overrides

import (
  "fmt"
  "time"

  "github.com/jdlubrano/pagerduty-cli/api_client"
  "github.com/jdlubrano/pagerduty-cli/api_client/users"
)

type Createable interface {
  ScheduleId() string
  StartTime() time.Time
  EndTime() time.Time
  UserId() string
}

type OverrideData struct {
  Override Override `json:"override"`
}

type Override struct {
  Id string `json:"id"`
  StartTime time.Time `json:"start"`
  EndTime time.Time `json:"end"`
  User users.User `json:"user"`
}

type overrideRequest struct {
  Start string `json:"start"`
  End string `json:"end"`
  User overrideRequestUser `json:"user"`
}

type overrideRequestUser struct {
  Id string `json:"id"`
  Type string `json:"type"`
}

func CreateOverride(client *api_client.ApiClient, override Createable) (*Override, error) {
  overrideParams := overrideRequest{
    override.StartTime().Format(time.RFC3339),
    override.EndTime().Format(time.RFC3339),
    overrideRequestUser{
      Id: override.UserId(),
      Type: "user_reference",
    },
  }

  params := make(map[string]interface{})
  params["override"] = overrideParams
  requestPath := fmt.Sprintf("/schedules/%s/overrides", override.ScheduleId())
  resp, err := client.Post(requestPath, &params)

  if err != nil {
    return nil, err
  }

  var data OverrideData
  resp.ParseInto(&data)
  return &data.Override, nil
}
