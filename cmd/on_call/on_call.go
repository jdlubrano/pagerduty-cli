package on_call

import (
  "fmt"
  "os"
  "time"

  "github.com/spf13/cobra"
  "github.com/olekukonko/tablewriter"

  "github.com/jdlubrano/pagerduty-cli/api_client"
  "github.com/jdlubrano/pagerduty-cli/api_client/escalation_policies"
  "github.com/jdlubrano/pagerduty-cli/api_client/users"
)

type OncallsData struct {
  Oncalls []Oncall `json:"oncalls"`
}

type Oncall struct {
  Schedule Schedule `json:"schedule"`
  User User `json:"user"`
  Start time.Time `json:"start"`
  End time.Time `json:"end"`
}

type Schedule struct {
  Id string `json:"id"`
  Summary string `json:"summary"`
}

type User struct {
  Id string `json:"id"`
  Summary string `json:"summary"`
}

func NewOnCallCmd() *cobra.Command {
  onCallCmd := &cobra.Command{
    Use: "on-call",
    Aliases: []string{"oc"},
    Short: "View on-call schedules",
    Long: `Various commands to view your on-call schedule or a team's on-call schedule`,
  }

  client := api_client.NewClient()

  onCallCmd.AddCommand(NewMeCmd(client))
  onCallCmd.AddCommand(NewTeamCmd(client))

  return onCallCmd
}

func NewMeCmd(client *api_client.ApiClient) *cobra.Command {
  return &cobra.Command{
    Use: "me",
    Short: "Show your upcoming on-call schedule",
    Aliases: []string{"mine"},
    Run: func(_ *cobra.Command, _ []string) {
      me, err := users.GetMe(client)

      if err != nil {
        fmt.Println(err)
        return
      }

      params := &map[string]string{
        "user_ids[]": me.Id,
        "until": formatTimeForApi(time.Now().AddDate(0, 3, 0)), // 3 months from now
      }

      oncalls, err := getOnCalls(client, params)

      if err != nil {
        fmt.Println(err)
        return
      }

      buildOncallsTable(oncalls).Render()
    },
  }
}

func NewTeamCmd(client *api_client.ApiClient) *cobra.Command {
  var teamName string

  teamCmd := &cobra.Command{
    Use: "team",
    Short: "Show who is currently on-call for a given team (or area)",
    Run: func(_ *cobra.Command, _ []string) {
      escalationPolicies, err := escalation_policies.ListForQuery(client, teamName)

      if err != nil {
        fmt.Println(err)
        return
      }

      var oncalls []Oncall

      for _, escalationPolicy := range escalationPolicies {
        params := &map[string]string{"escalation_policy_ids[]": escalationPolicy.Id}
        fmt.Printf("%+v\n", escalationPolicy)
        oncallsForPolicy, err := getOnCalls(client, params)
        fmt.Printf("%+v\n", oncallsForPolicy)

        if err != nil {
          fmt.Println(err)
          return
        }

        oncalls = append(oncalls, oncallsForPolicy...)
      }

      oncalls = removeDuplicateSchedules(oncalls)
      buildOncallsTable(oncalls).Render()
    },
  }

  teamCmd.Flags().StringVarP(&teamName, "team", "t", "", "Team name (required)")
  teamCmd.MarkFlagRequired("team")

  return teamCmd
}

func removeDuplicateSchedules(oncalls []Oncall) []Oncall {
  uniqueOncalls := []Oncall{}
  schedules := make(map[string]bool)

  for _, oncall := range oncalls {
    id := oncall.Schedule.Id
    if !schedules[id] {
      uniqueOncalls = append(uniqueOncalls, oncall)
    }

    schedules[id] = true
  }

  return uniqueOncalls
}

func buildOncallsTable(oncalls []Oncall) *tablewriter.Table {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Schedule", "User", "Start", "End"})

  for _, oncall := range oncalls {
    if oncall.Schedule.Id != "" {
      table.Append([]string{
        oncall.Schedule.Summary,
        oncall.User.Summary,
        formatTimeForUser(oncall.Start),
        formatTimeForUser(oncall.End),
      })
    }
  }

  return table
}

func formatTimeForApi(t time.Time) string {
  return t.Format(time.RFC3339)
}

func formatTimeForUser(t time.Time) string {
  tz := time.Now().Location()
  return t.In(tz).Format("02 Jan 15:04 MST")
}

func getOnCalls(client *api_client.ApiClient, params *map[string]string) ([]Oncall, error) {
  resp, err := client.Get("/oncalls", params)

  if err != nil {
    return []Oncall{}, err
  }

  var oncallsData OncallsData
  resp.ParseInto(&oncallsData)
  return oncallsData.Oncalls, nil
}
