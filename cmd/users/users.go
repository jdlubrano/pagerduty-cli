package users

import (
  "fmt"
  "os"
  "strings"

  "github.com/spf13/cobra"
  "github.com/olekukonko/tablewriter"

  "github.com/jdlubrano/pagerduty-cli/api_client"
)

type UserData struct {
  User User `json:"user"`
}

type User struct {
  Id    string `json:"id"`
  Name  string `json:"name"`
  Teams []Team `json:"teams"`
}

type Team struct {
  Id   string `json:"id"`
  Name string `json:"summary"`
}

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

      if err != nil {
        fmt.Println(err)
        return
      }

      var userData UserData
      resp.ParseInto(&userData)
      user := userData.User

      buildUsersTable([]User{user}).Render()
    },
  }

  return meCmd
}

func buildUsersTable(users []User) *tablewriter.Table {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"ID", "Name", "Teams"})

  for _, user := range users {
    table.Append([]string{user.Id, user.Name, extractTeamNames(user.Teams)})
  }

  return table
}

func extractTeamNames(teams []Team) string {
  teamNames := make([]string, len(teams))

  for i, team := range teams {
    teamNames[i] = team.Name
  }

  return strings.Join(teamNames, ", ")
}
