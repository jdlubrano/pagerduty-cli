package schedules

import (
  "github.com/spf13/cobra"

  "github.com/jdlubrano/pagerduty-cli/api_client"
  "github.com/jdlubrano/pagerduty-cli/cmd/schedules/schedule_overrides"
)

func NewSchedulesCmd() *cobra.Command {
  schedulesCmd := &cobra.Command {
    Use: "schedules",
    Short: "Adjust schedules",
    Long: `Various commands to adjust schedules`,
  }

  client := api_client.NewClient()

  schedulesCmd.AddCommand(schedule_overrides.NewOverridesCmd(client))

  return schedulesCmd
}
