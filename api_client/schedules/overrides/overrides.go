package schedule_overrides

import (
  "time"

  "github.com/jdlubrano/pagerduty-cli/api_client"
)

type OverrideData struct {
  Override Override `json:"override"`
}

type Createable interface {
  ScheduleId() string
  StartTime() time.Time
  EndTime() time.Time
  UserId() string
}

type Override struct {
  Id string `json:"id"`
  StartTime time.Time `json:"start"`
  EndTime time.Time `json:"end"`
  User users.User `json:"user"`
}

func CreateOverride(client *api_client.ApiClient, override *Createable) (Override, error) {
}
