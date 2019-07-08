package users

import (
  "fmt"
  "os"
  "strings"

  "github.com/spf13/cobra"
  "github.com/olekukonko/tablewriter"

  "github.com/jdlubrano/pagerduty-cli/api_client"
  "github.com/jdlubrano/pagerduty-cli/api_client/teams"
  "github.com/jdlubrano/pagerduty-cli/api_client/users"
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
  return &cobra.Command{
    Use: "me",
    Short: "Show your user information",
    Run: func(_ *cobra.Command, _ []string) {
      user, err := users.GetMe(client)

      if err != nil {
        fmt.Println(err)
        return
      }

      buildUsersTable([]users.User{user}).Render()
    },
  }
}

func buildUsersTable(users []users.User) *tablewriter.Table {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"ID", "Name", "Teams"})

  for _, user := range users {
    table.Append([]string{user.Id, user.Name, extractTeamNames(user.Teams)})
  }

  return table
}

func extractTeamNames(teams []teams.Team) string {
  teamNames := make([]string, len(teams))

  for i, team := range teams {
    teamNames[i] = team.Name
  }

  return strings.Join(teamNames, ", ")
}
