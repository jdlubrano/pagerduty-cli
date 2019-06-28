package users

import (
  "fmt"
  "io/ioutil"
  "github.com/spf13/cobra"
  "github.com/jdlubrano/pagerduty-cli/api_client"
)

func NewUsersCmd() *cobra.Command {
  usersCmd := &cobra.Command{
    Use: "users",
    Short: "View team member info",
  }

  client := api_client.NewClient()

  usersCmd.AddCommand(NewMeCmd(client))

  return usersCmd
}

func NewMeCmd(client *api_client.ApiClient) *cobra.Command {
  meCmd := &cobra.Command{
    Use: "me",
    Short: "Show your user information",
    Run: func(_ *cobra.Command, _ []string) {
      resp, err := client.Get("/users/me", nil)
      defer resp.Body.Close()

      if err != nil {
        fmt.Println(err)
        return
      }

      json, err := ioutil.ReadAll(resp.Body)

      if err != nil {
        fmt.Println(err)
      }

      fmt.Println(string(json))
    },
  }

  return meCmd
}
