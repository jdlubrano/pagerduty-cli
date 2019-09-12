package schedule_overrides

import (
  "fmt"

  "github.com/spf13/cobra"

  "github.com/jdlubrano/pagerduty-cli/api_client"
  "github.com/jdlubrano/pagerduty-cli/api_client/schedules"
)

func NewOverridesCmd(client *api_client.ApiClient) *cobra.Command {
  overridesCmd := &cobra.Command{
    Use: "overrides",
    Short: "Create schedule overrides",
    Long: `Various commands to adjust a schedule's overrides`,
  }

  overridesCmd.AddCommand(NewCreateCmd(client))

  return overridesCmd
}

func NewCreateCmd(client *api_client.ApiClient) *cobra.Command {
  var scheduleName string

  createCmd := &cobra.Command{
    Use: "create",
    Short: "Create a schedule override for a given schedule",
    Run: func(_ *cobra.Command, _ []string) {
      aSchedules, err := schedules.ListForQuery(client, scheduleName)

      if err != nil {
        fmt.Println(err)
        return
      }

      for _, schedule := range aSchedules {
        fmt.Println(schedule.Name)
      }
    },
  }

  createCmd.Flags().StringVarP(&scheduleName, "schedule", "s", "", "Schedule name (required)")
  createCmd.MarkFlagRequired("schedule")

  return createCmd
}
