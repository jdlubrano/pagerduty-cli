package schedules

import (
  "github.com/jdlubrano/pagerduty-cli/api_client"
)

type SchedulesData struct {
  Schedules []Schedule `json:"schedules"`
}

type Schedule struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
}

func ListForQuery(client *api_client.ApiClient, scheduleQuery string) ([]Schedule, error) {
  params := &map[string]string{"query": scheduleQuery}

  resp, err := client.Get("/schedules", params)

  if err != nil {
    return []Schedule{}, err
  }

  var data SchedulesData
  resp.ParseInto(&data)
  return data.Schedules, nil
}
