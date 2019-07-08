package users

import (
  "github.com/jdlubrano/pagerduty-cli/api_client"
  "github.com/jdlubrano/pagerduty-cli/api_client/teams"
)

type UserData struct {
  User User `json:"user"`
}

type User struct {
  Id    string `json:"id"`
  Name  string `json:"name"`
  Teams []teams.Team `json:"teams"`
}

func GetMe(client *api_client.ApiClient) (User, error) {
  resp, err := client.Get("/users/me", nil)

  if err != nil {
    return User{}, err
  }

  var userData UserData
  resp.ParseInto(&userData)
  return userData.User, nil
}
